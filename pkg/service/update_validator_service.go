package service

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	relayer "github.com/nguyenkhanhvn/bas-relayer"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/abi"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/utils"
)

type UpdateValidatorService struct {
	BSC *UpdateValidatorWorker

	BAS *UpdateValidatorWorker

	Updating        bool
	UpdatingChannel chan bool
}

func NewUpdateValidatorService(bscConfig *relayer.Config, basConfig *relayer.Config) (*UpdateValidatorService, error) {
	bscClient, err := ethclient.Dial(bscConfig.ChainConfig.RpcUrl)
	if err != nil {
		return nil, err
	}
	bscSigner, err := crypto.HexToECDSA(bscConfig.SignerConfig.PrivateKey)
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
	basRelayHub, err := abi.NewRelayHub(common.HexToAddress(basConfig.ChainConfig.RelayHubAddress), basClient)
	if err != nil {
		return nil, err
	}
	status := make(chan bool)

	return &UpdateValidatorService{
		BSC: &UpdateValidatorWorker{
			Worker: Worker{
				ChainName: "BSC",
				Client:    bscClient,
				Config:    bscConfig,
				Signer:    bscSigner,
			},
			relayHub: bscRelayHub,
		},
		BAS: &UpdateValidatorWorker{
			Worker: Worker{
				ChainName: "BAS",
				Client:    basClient,
				Config:    basConfig,
				Signer:    basSigner,
			},
			relayHub: basRelayHub,
		},
		Updating:        false,
		UpdatingChannel: status,
	}, nil
}

func (s *UpdateValidatorService) Start() error {
	go s.UpdateValidatorSet(s.BSC, s.BAS)
	go s.UpdateValidatorSet(s.BAS, s.BSC)
	return nil
}

func (s *UpdateValidatorService) UpdateValidatorSet(fromChain *UpdateValidatorWorker, toChain *UpdateValidatorWorker) {
	blockNumberChan := fromChain.FetchBlockNumber()
	log.Printf("start UpdateValidatorSet from %v to %v\n", fromChain.ChainName, toChain.ChainName)
	go func() {
		for {
			lastestBlock := <-blockNumberChan.DataTranferChannel
			lastestEpoch := lastestBlock / fromChain.Config.ChainConfig.EpochLength
			lastestKnownEpochRaw, err := toChain.relayHub.GetLatestEpochNumber(nil, big.NewInt(fromChain.Config.ChainConfig.ChainId))
			if err != nil {
				log.Printf("UpdateValidatorSet: failed to get GetLatestEpochNumber from %v: %v\n", toChain.ChainName, err)
				continue
			}
			nextUpdatedEpoch := lastestKnownEpochRaw.Uint64() + 1
			if nextUpdatedEpoch > lastestEpoch+1 {
				panic("unexpected epoch")
			}
			if nextUpdatedEpoch >= lastestEpoch && s.Updating {
				s.Updating = false
				s.UpdatingChannel <- false
			}
			for nextUpdatedEpoch <= lastestEpoch {
				if nextUpdatedEpoch < lastestEpoch {
					if !s.Updating {
						s.Updating = true
						s.UpdatingChannel <- true
					}
				}

				fromBlock := nextUpdatedEpoch * fromChain.Config.ChainConfig.EpochLength
				toBlock := fromBlock + uint64(fromChain.Config.ChainConfig.BlockConfirmations)
				if toBlock <= lastestBlock {
					_, blockProofs, err := utils.GenerateBlockProofs(fromChain.Client, int64(fromBlock), int64(toBlock))
					if err != nil {
						log.Printf("UpdateValidatorSet: failed to generateBlockProofs from %v: %v\n", fromChain.ChainName, err)
						break
					}

					opts, err := utils.GetTransactOpts(toChain.Client, toChain.Signer, big.NewInt(toChain.Config.ChainConfig.ChainId))
					if err != nil {
						log.Printf("UpdateValidatorSet: failed to getTransactOpts for %v: %v\n", toChain.ChainName, err)
						break
					}

					// update validator set
					log.Printf("updateing validator set from %v to %v, epoch: %v\n", fromChain.ChainName, toChain.ChainName, nextUpdatedEpoch)
					tx, err := toChain.relayHub.UpdateValidatorSetUsingEpochBlocks(opts, big.NewInt(fromChain.Config.ChainConfig.ChainId), blockProofs.HeaderRlps)
					if err != nil {
						log.Printf("UpdateValidatorSet: failed to UpdateValidatorSetUsingEpochBlocks to %v: %v\n", toChain.ChainName, err)
						break
					}

					err = utils.WaitTxToBeMined(context.TODO(), toChain.Client, tx.Hash())
					if err != nil {
						log.Printf("UpdateValidatorSet: failed to get transaction receipt in %v: %v\n", toChain.ChainName, err)
						break
					}
				}

				nextUpdatedEpoch += 1
			}
		}
	}()
}

type UpdateValidatorWorker struct {
	Worker

	relayHub *abi.RelayHub
}

func (sw *UpdateValidatorWorker) GetLatestEpoch() (*big.Int, error) {
	return sw.relayHub.GetLatestEpochNumber(nil, big.NewInt(sw.Config.ChainConfig.ChainId))
}
