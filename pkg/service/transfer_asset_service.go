package service

import (
	"context"
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

type TransferAssetService struct {
	BSC *TransferAssetWorker

	BAS *TransferAssetWorker

	stopSignal chan struct{}
}

func NewTransferAssetService(bscConfig *relayer.Config, basConfig *relayer.Config) (*TransferAssetService, error) {
	bscClient, err := ethclient.Dial(bscConfig.ChainConfig.RpcUrl)
	if err != nil {
		return nil, err
	}
	bscSigner, err := crypto.HexToECDSA(bscConfig.SignerConfig.PrivateKey)
	if err != nil {
		return nil, err
	}
	bscBridge, err := abi.NewCrossChainBridge(common.HexToAddress(bscConfig.ChainConfig.CrossChainBridgeAddress), bscClient)
	if err != nil {
		return nil, err
	}

	basClient, err := ethclient.Dial(basConfig.ChainConfig.RpcUrl)
	if err != nil {
		return nil, err
	}
	basSigner, err := crypto.HexToECDSA(basConfig.SignerConfig.PrivateKey)
	if err != nil {
		return nil, err
	}
	basBridge, err := abi.NewCrossChainBridge(common.HexToAddress(basConfig.ChainConfig.CrossChainBridgeAddress), basClient)
	if err != nil {
		return nil, err
	}

	return &TransferAssetService{
		BSC: &TransferAssetWorker{
			Worker: Worker{
				ChainName: "BSC",
				Client:    bscClient,
				Config:    bscConfig,
				Signer:    bscSigner,
			},
			LastestReadedBlock: 0,
			Bridge:             bscBridge,
		},
		BAS: &TransferAssetWorker{
			Worker: Worker{
				ChainName: "BAS",
				Client:    basClient,
				Config:    basConfig,
				Signer:    basSigner,
			},
			LastestReadedBlock: 0,
			Bridge:             basBridge,
		},
	}, nil
}

func (s *TransferAssetService) Start() error {
	// get lastest block
	_, err := s.BSC.UpdateLastestBlock()
	if err != nil {
		return err
	}

	_, err = s.BAS.UpdateLastestBlock()
	if err != nil {
		return err
	}

	log.Println("starting listen to asset transfer request")
	// listen requests from bsc to bas
	go s.ListenToRequests(s.BSC, s.BAS)
	go s.ListenToRequests(s.BAS, s.BSC)
	return nil
}
func (s *TransferAssetService) Stop() {
	s.stopSignal <- struct{}{}
}

// listen to deposit requests, get Proof and withdraw at destination chain
func (s *TransferAssetService) ListenToRequests(fromChain *TransferAssetWorker, toChain *TransferAssetWorker) {
	blockNumberChan := fromChain.FetchBlockNumber()
	log.Printf("start ListenToRequests from %v to %v\n", fromChain.ChainName, toChain.ChainName)
	go func() {
		for {
			select {
			case toBlock := <-blockNumberChan.DataTranferChannel:
				fromBlock := fromChain.LastestReadedBlock + 1
				if fromBlock > toBlock {
					continue
				}

				log.Printf("listening to transfer assets request from %v to %v, fromBlock: %v, toBlock: %v\n", fromChain.ChainName, toChain.ChainName, fromBlock, toBlock)
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
					log.Printf("ListenToRequests: failed to FilterLogs from %v: %v\n", fromChain.ChainName, err)
					continue
				}

				for _, vlog := range logs {
					go s.TransferAsset(fromChain, vlog.TxHash, toChain)
				}

				fromChain.LastestReadedBlock = toBlock
			case <-s.stopSignal:
				blockNumberChan.Stop()
				return
			}
		}
	}()
}

func (s *TransferAssetService) TransferAsset(fromChain *TransferAssetWorker, txHash common.Hash, toChain *TransferAssetWorker) {
	log.Printf("start withdraw from %v, txHash: %v\n", fromChain.ChainName, txHash)
	// make proof at source chain
	var lastestNumber uint64
	blockNumber, _, err := utils.GetBlockNumberFromTransactionHash(fromChain.Client, txHash)
	if err != nil {
		log.Printf("generateReceiptProof: cannot getBlockNumberFromTransactionHash, error: %v", err)
	}
	blockNumberChan := fromChain.FetchBlockNumber()
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
		log.Printf("TransferAsset: failed to generateProof from %v: %v\n", fromChain.ChainName, err)
		return
	}

	// withdraw at destination chain
	err = toChain.Withdraw(proofRaw)
	if err != nil {
		log.Printf("TransferAsset: failed to withdraw to %v: %v\n", toChain.ChainName, err)
		return
	}
	log.Println("finish withdraw: ", txHash)
}

type TransferAssetWorker struct {
	Worker

	LastestReadedBlock uint64

	Bridge *abi.CrossChainBridge
}

// update latest block of BSC, return previous lastest of BSC
func (sw *TransferAssetWorker) UpdateLastestBlock() (uint64, error) {
	lastestBlock, err := sw.Client.BlockNumber(context.TODO())
	if err != nil {
		return 0, err
	}
	previousLastestBlock := sw.LastestReadedBlock
	sw.LastestReadedBlock = lastestBlock
	return previousLastestBlock, err
}

func (sw *TransferAssetWorker) Withdraw(proof *utils.ProofRaw) error {
	opts, err := utils.GetTransactOpts(sw.Client, sw.Signer, big.NewInt(sw.Config.ChainConfig.ChainId))
	if err != nil {
		return err
	}

	// withdraw
	tx, err := sw.Bridge.Withdraw(opts, proof.BlockProofs, proof.RawReceipt, proof.ProofPath, proof.ProofSiblings)
	if err != nil {
		return err
	}
	return utils.WaitTxToBeMined(context.TODO(), sw.Client, tx.Hash())
}
