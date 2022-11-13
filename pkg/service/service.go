package service

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	relayer "github.com/nguyenkhanhvn/bas-relayer"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/abi"
)

type Service struct {
	BSC *ServiceWorker
	BAS *ServiceWorker

	status bool
}

func NewService(bscConfig *relayer.Config, basConfig *relayer.Config) (*Service, error) {
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
	bscRelayHub, err := abi.NewRelayHub(common.HexToAddress(bscConfig.ChainConfig.RelayHubAddress), bscClient)
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
	basRelayHub, err := abi.NewRelayHub(common.HexToAddress(basConfig.ChainConfig.RelayHubAddress), basClient)
	if err != nil {
		return nil, err
	}

	return &Service{
		BSC: &ServiceWorker{
			UpdateValidatorWorker: NewUpdateValidatorWorker("BSC", bscClient, bscConfig, bscSigner, bscRelayHub),
			TransferAssetWorker:   NewTransferAssetWorker("BSC", bscClient, bscConfig, bscSigner, bscBridge),
		},
		BAS: &ServiceWorker{
			UpdateValidatorWorker: NewUpdateValidatorWorker("BAS", basClient, basConfig, basSigner, basRelayHub),
			TransferAssetWorker:   NewTransferAssetWorker("BAS", basClient, basConfig, basSigner, basBridge),
		},
		status: false,
	}, nil
}

func (s *Service) Start() error {
	if !s.status {
		s.status = true
		err := s.BAS.startServiceWithChain(s.BSC)
		if err != nil {
			return err
		}
		err = s.BSC.startServiceWithChain(s.BAS)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Service) Stop() {
	if s.status {
		s.status = false
		s.BSC.stop()
		s.BAS.stop()
	}
}

type ServiceWorker struct {
	UpdateValidatorWorker *UpdateValidatorWorker
	TransferAssetWorker   *TransferAssetWorker
	stopSignal            chan struct{}
}

func (sw *ServiceWorker) startServiceWithChain(fromChain *ServiceWorker) error {
	sw.TransferAssetWorker.FetchLastestReadedBlock()
	go sw.UpdateValidatorWorker.UpdateValidatorSetFromChain(fromChain.UpdateValidatorWorker)
	go func(fromChain *ServiceWorker) {
		for {
			select {
			case updating := <-sw.UpdateValidatorWorker.UpdatingChannel():
				if updating {
					log.Printf("stop transfer asset to update validators from %v to %v", fromChain.TransferAssetWorker.ChainName, sw.TransferAssetWorker.ChainName)
					sw.TransferAssetWorker.Stop()
				} else {
					log.Println("start transfer asset")
					go sw.TransferAssetWorker.ListenToRequestsFromChain(fromChain.TransferAssetWorker)
				}
			case <-sw.stopSignal:
				return
			}
		}
	}(fromChain)
	go sw.TransferAssetWorker.ListenToRequestsFromChain(fromChain.TransferAssetWorker)
	return nil
}

func (sw *ServiceWorker) stop() {
	sw.UpdateValidatorWorker.Stop()
	sw.TransferAssetWorker.Stop()
	sw.stopSignal <- struct{}{}
}
