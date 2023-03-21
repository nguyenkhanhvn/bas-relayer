package service

import (
	"context"
	"fmt"
	"log"
	"math/big"

	relayer "relayer"
	"relayer/pkg/abi"
	"relayer/pkg/utils"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type TransferAssetWorker struct {
	Worker

	LastestReadedBlock map[*TransferAssetWorker]int64
	Bridge             *abi.CrossChainBridge
	stopSignal         chan struct{}
	stopping           bool
}

func NewTransferAssetWorker(chainName string, config *relayer.Config) *TransferAssetWorker {
	if config == nil {
		return nil
	}
	worker := &TransferAssetWorker{
		Worker: Worker{
			ChainName: chainName,
			Config:    config,
		},
		LastestReadedBlock: make(map[*TransferAssetWorker]int64),
		stopSignal:         make(chan struct{}),
		stopping:           false,
	}
	worker.ConnectToRpc()
	return worker
}

func (w *TransferAssetWorker) ConnectToRpc() error {
	log.Printf("TransferAssetWorker ConnectToRpc")
	err := w.Worker.ConnectToRpc()
	if err != nil {
		return err
	}
	w.Bridge, err = abi.NewCrossChainBridge(common.HexToAddress(w.Config.ChainConfig.CrossChainBridgeAddress), w.Client)
	if err != nil {
		return fmt.Errorf("ConnectToRpc: %v", err)
	}
	return nil
}

func (w *TransferAssetWorker) Stop() {
	if w.running && !w.stopping {
		w.stopSignal <- struct{}{}
		w.stopping = true
	}
}

func (w *TransferAssetWorker) stopped() {
	w.running = false
	w.stopping = false
}

// listen to deposit requests, get Proof and withdraw at destination chain
func (w *TransferAssetWorker) ListenToRequestsFromChain(fromChain *TransferAssetWorker, id uint64) {
	if w.isRunning() {
		return
	}

	w.start()
	defer w.stopped()

	log.Printf("id %v: start ListenToRequests from %v to %v\n", id, fromChain.ChainName, w.ChainName)
	defer log.Printf("id %v: stop ListenToRequests from %v to %v\n", id, fromChain.ChainName, w.ChainName)

	var blockNumberChan *BlockNumberChannel
	defer func() {
		if blockNumberChan != nil {
			blockNumberChan.Stop()
		}
	}()
	filterLogsMaximum := fromChain.Config.ChainConfig.FilterLogsMaximumBlockRange

MAIN_LOOP_ASSET_TRANSFER:
	for {
		if blockNumberChan == nil {
			blockNumberChan = fromChain.GetBlockNumber()
			if blockNumberChan == nil {
				err := w.ConnectToRpc()
				if err != nil {
					log.Println(err.Error())
					continue MAIN_LOOP_ASSET_TRANSFER
				}
			}
		}

		if blockNumberChan != nil {
			blockNumberChan.Start()

			select {
			case <-w.stopSignal:
				return

			case <-blockNumberChan.ErrorChannel:
				blockNumberChan.Stop()
				blockNumberChan = nil
				w.ConnectToRpc()

			case toBlock := <-blockNumberChan.DataTranferChannel:
				fromBlock := w.LastestReadedBlock[fromChain] + 1
				if fromBlock > toBlock {
					continue
				}

				if filterLogsMaximum > 0 && toBlock-fromBlock > filterLogsMaximum {
					toBlock = fromBlock + filterLogsMaximum
				}

				log.Printf("id %v: listening to transfer assets request from %v to %v, fromBlock: %v, toBlock: %v\n", id, fromChain.ChainName, w.ChainName, fromBlock, toBlock)
				query := ethereum.FilterQuery{
					FromBlock: big.NewInt(int64(fromBlock)),
					ToBlock:   big.NewInt(int64(toBlock)),
					Addresses: []common.Address{common.HexToAddress(fromChain.Config.ChainConfig.CrossChainBridgeAddress)},
					Topics:    [][]common.Hash{{crypto.Keccak256Hash([]byte(fromChain.Config.ChainConfig.EventSignature))}},
				}

				logs, err := fromChain.Client.FilterLogs(context.TODO(), query)
				if err != nil {
					log.Printf("id %v: failed to FilterLogs from %v: %v\n", id, fromChain.ChainName, err)

					blockNumberChan.Stop()
					blockNumberChan = nil
					w.ConnectToRpc()
					continue
				}

				for _, logDetail := range logs {
					if len(logDetail.Topics) >= relayer.LOG_TOPICS_NUMBER && logDetail.Topics[relayer.LOG_CHAIN_ID_INDEX].Big().Int64() == w.Config.ChainConfig.ChainId {
						go w.TransferAsset(fromChain, logDetail.TxHash, id)
					}
				}

				w.LastestReadedBlock[fromChain] = toBlock
			}
		}
	}
}

func (w *TransferAssetWorker) TransferAsset(fromChain *TransferAssetWorker, txHash common.Hash, id uint64) {
	log.Printf("id %v: start withdraw from %v, txHash: %v, waiting for confirmation...\n", id, fromChain.ChainName, txHash)
	// make proof at source chain
	var lastestNumber int64
	blockNumber, _, err := utils.GetBlockNumberFromTransactionHash(fromChain.Client, txHash)
	if err != nil {
		log.Printf("id %v: cannot getBlockNumberFromTransactionHash, error: %v", id, err.ErrorString())
	}
	blockNumberChan := fromChain.GetBlockNumber()
	blockNumberChan.Start()
	for {
		lastestNumber = <-blockNumberChan.DataTranferChannel
		if blockNumber.Int64()+fromChain.Config.ChainConfig.BlockConfirmations() > lastestNumber {
			continue
		}
		blockNumberChan.Stop()
		break
	}

	_, proofRaw, _, _, err := utils.GenerateProofWithoutCheck(fromChain.Client, blockNumber.Int64(), blockNumber.Int64()+fromChain.Config.ChainConfig.BlockConfirmations(), txHash)
	if err != nil {
		log.Printf("id %v: failed to generateProof from %v, txHash: %v, error: %v\n", id, fromChain.ChainName, txHash, err.ErrorString())
		return
	}

	// withdraw at destination chain
	tx, err := w.Withdraw(proofRaw)
	if err != nil {
		if tx != nil {
			log.Printf("id %v:failed to withdraw from %v, txHash: %v, to %v: %v, error: %v\n", id, fromChain.ChainName, txHash, w.ChainName, tx.Hash(), err.ErrorString())
		} else {
			log.Printf("id %v:failed to withdraw from %v to %v, txHash: %v, error: %v\n", id, fromChain.ChainName, txHash, w.ChainName, err.ErrorString())
		}
		return
	}
	log.Printf("id %v: finish withdraw from %v: %v, to %v: %v\n", id, fromChain.ChainName, txHash, w.ChainName, tx.Hash())
}

func (w *TransferAssetWorker) FetchLastestReadedBlock(chains []*TransferAssetWorker) {
	for _, chain := range chains {
		if chain.Config.OtherConfig.BlockNumber > 0 {
			w.LastestReadedBlock[chain] = int64(chain.Config.OtherConfig.BlockNumber)
		} else {
			err := w.FetchLastestReadedBlockToNewest([]*TransferAssetWorker{chain})
			if err != nil {
				log.Panicln(err.ErrorString())
			}
		}
	}
}

func (w *TransferAssetWorker) FetchLastestReadedBlockToNewest(chains []*TransferAssetWorker) utils.Error {
	for _, chain := range chains {
		lastestBlock, err := chain.Client.BlockNumber(context.TODO())
		if err != nil {
			return utils.NewError(utils.ClientError, err)
		}

		w.LastestReadedBlock[chain] = int64(lastestBlock)
	}
	return nil
}

func (w *TransferAssetWorker) Withdraw(proof *utils.ProofRaw) (*types.Transaction, utils.Error) {
	opts, err := utils.GetTransactOpts(w.Client, w.Signer, big.NewInt(w.Config.ChainConfig.ChainId))
	if err != nil {
		return nil, err
	}

	// withdraw
	tx, errS := w.Bridge.Withdraw(opts, proof.BlockProofs, proof.RawReceipt, proof.ProofPath, proof.ProofSiblings)
	if errS != nil {
		return tx, utils.NewError(utils.OtherError, fmt.Errorf("failed to Withdraw: %v", errS))
	}
	log.Printf("withdrawing in %v txHash: %v...\n", w.ChainName, tx.Hash())
	return tx, utils.WaitTxToBeMined(context.TODO(), w.Client, tx.Hash())
}
