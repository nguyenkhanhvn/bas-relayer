package service

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	relayer "github.com/nguyenkhanhvn/bas-relayer"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/abi"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/utils"
)

type TransferAssetWorker struct {
	Worker

	LastestReadedBlock uint64
	Bridge             *abi.CrossChainBridge
	stopSignal         chan struct{}
}

func NewTransferAssetWorker(chainName string, client *ethclient.Client, config *relayer.Config, signer *ecdsa.PrivateKey, bridge *abi.CrossChainBridge) *TransferAssetWorker {
	return &TransferAssetWorker{
		Worker: Worker{
			ChainName: chainName,
			Client:    client,
			Config:    config,
			Signer:    signer,
		},
		LastestReadedBlock: 0,
		Bridge:             bridge,
		stopSignal:         make(chan struct{}),
	}
}

func (w *TransferAssetWorker) Stop() {
	w.stopSignal <- struct{}{}
}

// listen to deposit requests, get Proof and withdraw at destination chain
func (w *TransferAssetWorker) ListenToRequestsFromChain(fromChain *TransferAssetWorker) {
	log.Printf("start ListenToRequests from %v to %v\n", fromChain.ChainName, w.ChainName)
	blockNumberChan := fromChain.GetBlockNumber()
	defer blockNumberChan.Stop()
	for {
		select {
		case toBlock := <-blockNumberChan.DataTranferChannel:
			fromBlock := fromChain.LastestReadedBlock + 1
			if fromBlock > toBlock {
				continue
			}

			log.Printf("listening to transfer assets request from %v to %v, fromBlock: %v, toBlock: %v\n", fromChain.ChainName, w.ChainName, fromBlock, toBlock)
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(fromBlock)),
				// FromBlock: big.NewInt(int64(88331)),
				ToBlock: big.NewInt(int64(toBlock)),
				// ToBlock:   big.NewInt(int64(88333)),
				Addresses: []common.Address{common.HexToAddress(fromChain.Config.ChainConfig.CrossChainBridgeAddress)},
				Topics:    [][]common.Hash{{crypto.Keccak256Hash([]byte(fromChain.Config.ChainConfig.EventSignature))}},
			}

			logs, err := fromChain.Client.FilterLogs(context.Background(), query)
			if err != nil {
				fromChain.LastestReadedBlock = fromBlock
				log.Printf("failed to FilterLogs from %v: %v\n", fromChain.ChainName, err)
				continue
			}

			for _, vlog := range logs {
				go w.TransferAsset(fromChain, vlog.TxHash)
			}

			fromChain.LastestReadedBlock = toBlock
		case <-w.stopSignal:
			return
		}
	}
}

func (w *TransferAssetWorker) TransferAsset(fromChain *TransferAssetWorker, txHash common.Hash) {
	log.Printf("start withdraw from %v, txHash: %v\n", fromChain.ChainName, txHash)
	// make proof at source chain
	var lastestNumber uint64
	blockNumber, _, err := utils.GetBlockNumberFromTransactionHash(fromChain.Client, txHash)
	if err != nil {
		log.Printf("cannot getBlockNumberFromTransactionHash, error: %v", err)
	}
	blockNumberChan := fromChain.GetBlockNumber()
	for {
		lastestNumber = <-blockNumberChan.DataTranferChannel
		if blockNumber.Int64()+fromChain.Config.ChainConfig.BlockConfirmations > int64(lastestNumber) {
			continue
		}
		blockNumberChan.Stop()
		break
	}

	_, proofRaw, _, _, err := utils.GenerateProofWithoutCheck(fromChain.Client, blockNumber.Int64(), int64(lastestNumber), txHash)
	if err != nil {
		log.Printf("failed to generateProof from %v: %v\n", fromChain.ChainName, err)
		return
	}

	// withdraw at destination chain
	err = w.Withdraw(proofRaw)
	if err != nil {
		log.Printf("failed to withdraw to %v: %v\n", w.ChainName, err)
		return
	}
	log.Println("finish withdraw: ", txHash)
}

// update latest block of BSC, return previous lastest of BSC
func (w *TransferAssetWorker) FetchLastestReadedBlock() (uint64, error) {
	lastestBlock, err := w.Client.BlockNumber(context.TODO())
	if err != nil {
		return 0, err
	}
	previousLastestBlock := w.LastestReadedBlock
	w.LastestReadedBlock = lastestBlock
	return previousLastestBlock, err
}

func (w *TransferAssetWorker) Withdraw(proof *utils.ProofRaw) error {
	opts, err := utils.GetTransactOpts(w.Client, w.Signer, big.NewInt(w.Config.ChainConfig.ChainId))
	if err != nil {
		return err
	}

	// withdraw
	tx, err := w.Bridge.Withdraw(opts, proof.BlockProofs, proof.RawReceipt, proof.ProofPath, proof.ProofSiblings)
	if err != nil {
		return err
	}
	return utils.WaitTxToBeMined(context.TODO(), w.Client, tx.Hash())
}
