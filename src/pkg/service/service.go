package service

import (
	"log"

	relayer "relayer"
)

var i uint64 = 0

type Service struct {
	BSC *ServiceWorker
	BAS *ServiceWorker

	relayerConfig *relayer.RelayerConfig
	status        bool
}

func NewService(relayerConfig *relayer.RelayerConfig, bscConfig *relayer.Config, basConfig *relayer.Config) (*Service, error) {
	log.Println("BSC RelayHub type", relayerConfig.RelayHubType.BSC)
	var bscRelayHubType RelayHubType
	switch relayerConfig.RelayHubType.BSC {
	case "alpha":
		bscRelayHubType = Alpha
	default:
		fallthrough
	case "beta":
		bscRelayHubType = Beta
	}

	log.Println("BAS RelayHub type", relayerConfig.RelayHubType.BAS)
	var basRelayHubType RelayHubType
	switch relayerConfig.RelayHubType.BAS {
	default:
		fallthrough
	case "alpha":
		basRelayHubType = Alpha
	case "beta":
		basRelayHubType = Beta
	}

	return &Service{
		BSC: &ServiceWorker{
			UpdateValidatorWorker: NewUpdateValidatorWorker("BSC", bscRelayHubType, bscConfig),
			TransferAssetWorker:   NewTransferAssetWorker("BSC", bscConfig),
		},
		BAS: &ServiceWorker{
			UpdateValidatorWorker: NewUpdateValidatorWorker("BAS", basRelayHubType, basConfig),
			TransferAssetWorker:   NewTransferAssetWorker("BAS", basConfig),
		},
		relayerConfig: relayerConfig,
		status:        false,
	}, nil
}

func (s *Service) Start() error {
	if !s.status {
		s.status = true

		if s.relayerConfig.Flow.BscToBas {
			err := s.BAS.startServiceFromChain(s.BSC)
			if err != nil {
				return err
			}
		}

		if s.relayerConfig.Flow.BasToBsc {
			err := s.BSC.startServiceFromChain(s.BAS)
			if err != nil {
				return err
			}
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

func (sw *ServiceWorker) startServiceFromChain(fromChain *ServiceWorker) error {
	// fetch lastest block to listen from lastest block
	sw.TransferAssetWorker.FetchLastestReadedBlock([]*TransferAssetWorker{fromChain.TransferAssetWorker})

	go sw.UpdateValidatorWorker.UpdateValidatorSetFromChain(fromChain.UpdateValidatorWorker, i)
	i++
	go func(fromChain *ServiceWorker) {
		for {
			select {
			case updating := <-sw.UpdateValidatorWorker.UpdatingChannel():
				if updating {
					log.Printf("stop transfer asset to update validators from %v to %v\n", fromChain.TransferAssetWorker.ChainName, sw.TransferAssetWorker.ChainName)
					sw.TransferAssetWorker.Stop()
				} else {
					log.Printf("id %v: start transfer asset from %v to %v\n", i, fromChain.TransferAssetWorker.ChainName, sw.TransferAssetWorker.ChainName)
					go sw.TransferAssetWorker.ListenToRequestsFromChain(fromChain.TransferAssetWorker, i)
					i++
				}
			case <-sw.stopSignal:
				return
			}
		}
	}(fromChain)
	return nil
}

func (sw *ServiceWorker) stop() {
	sw.UpdateValidatorWorker.Stop()
	sw.TransferAssetWorker.Stop()
	sw.stopSignal <- struct{}{}
}
