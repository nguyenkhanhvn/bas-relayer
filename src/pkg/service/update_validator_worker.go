package service

import (
	"context"
	"fmt"
	"log"
	"math/big"

	relayer "relayer"
	"relayer/pkg/abi"
	"relayer/pkg/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type RelayHubType uint64

const (
	// Update ValidatorSet each epoch
	Alpha RelayHubType = iota

	// Update ValidatorSet only when changed
	Beta
)

type RelayHubInterface interface {
	GetLatestEpochNumber(opts *bind.CallOpts, chainId *big.Int) (*big.Int, error)
	UpdateValidatorSetUsingEpochBlocks(opts *bind.TransactOpts, chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error)
	GetLatestValidatorSet(opts *bind.CallOpts, chainId *big.Int) ([]common.Address, error)
}

type UpdateValidatorWorker struct {
	Worker

	RelayHubType    RelayHubType
	RelayHub        RelayHubInterface
	updating        bool
	updatingChannel chan bool
	stopSignal      chan struct{}
}

func NewUpdateValidatorWorker(chainName string, chainType RelayHubType, config *relayer.Config) *UpdateValidatorWorker {
	if config == nil {
		return nil
	}
	worker := &UpdateValidatorWorker{
		Worker: Worker{
			ChainName: chainName,
			Config:    config,
		},

		RelayHubType:    chainType,
		updating:        false,
		updatingChannel: make(chan bool),
		stopSignal:      make(chan struct{}),
	}
	err := worker.ConnectToRpc()
	if err != nil {
		log.Panic(err)
	}
	return worker
}

func (w *UpdateValidatorWorker) ConnectToRpc() error {
	log.Printf("UpdateValidatorWorker ConnectToRpc")
	err := w.Worker.ConnectToRpc()
	if err != nil {
		return err
	}

	switch w.RelayHubType {
	case Alpha:
		w.RelayHub, err = abi.NewRelayHubAlpha(common.HexToAddress(w.Config.ChainConfig.RelayHubAddress), w.Client)
		if err != nil {
			return fmt.Errorf("ConnectToRpc: %v", err)
		}
	default:
		fallthrough
	case Beta:
		w.RelayHub, err = abi.NewRelayHubBeta(common.HexToAddress(w.Config.ChainConfig.RelayHubAddress), w.Client)
		if err != nil {
			return fmt.Errorf("ConnectToRpc: %v", err)
		}
	}
	return nil
}

func (w *UpdateValidatorWorker) Stop() {
	w.stopSignal <- struct{}{}
}

// return updating status
func (w *UpdateValidatorWorker) Updating() bool {
	return w.updating
}

// return channel which will emit when updating status change
func (w *UpdateValidatorWorker) UpdatingChannel() <-chan bool {
	return w.updatingChannel
}

func (w *UpdateValidatorWorker) changeUpdatingStatus(status bool) bool {
	if w.updating != status {
		w.updating = status
		w.updatingChannel <- status
		return true
	}
	return false
}

func (w *UpdateValidatorWorker) stopped() {
	w.running = false
	w.changeUpdatingStatus(false)
}

func (w *UpdateValidatorWorker) UpdateValidatorSetFromChain(fromChain *UpdateValidatorWorker, id uint64) {
	if w.isRunning() {
		return
	}

	w.start()
	defer w.stopped()

	w.changeUpdatingStatus(true)
	log.Printf("id %v: start UpdateValidatorSet from %v to %v\n", id, fromChain.ChainName, w.ChainName)
	defer log.Printf("id %v: stop UpdateValidatorSet from %v to %v\n", id, fromChain.ChainName, w.ChainName)

	switch w.RelayHubType {
	case Alpha:
		w.alphaRelayHubUpdateValidatorSetFromChain(fromChain, id)
	case Beta:
		w.betaRelayHubUpdateValidatorSetFromChain(fromChain, id)
	default:
		log.Printf("id %v: RelayHubType not support: %v", id, w.RelayHubType)
	}
}

func (w *UpdateValidatorWorker) alphaRelayHubUpdateValidatorSetFromChain(fromChain *UpdateValidatorWorker, id uint64) {
	var blockNumberChan *BlockNumberChannel
	defer func() {
		if blockNumberChan != nil {
			blockNumberChan.Stop()
		}
	}()

MAIN_LOOP_UPDATE_VALIDATOR_ALPHA:
	for {
		if blockNumberChan == nil {
			blockNumberChan = fromChain.GetBlockNumber()
			if blockNumberChan == nil {
				err := w.ConnectToRpc()
				if err != nil {
					log.Println(err.Error())
					continue MAIN_LOOP_UPDATE_VALIDATOR_ALPHA
				}
			}
		}
		if blockNumberChan != nil {
			blockNumberChan.Start()
			select {
			case <-w.stopSignal:
				return

			case <-blockNumberChan.ErrorChannel:
				log.Printf("id %v: updated validator set from %v to %v blockNumberChan error!\n", id, fromChain.ChainName, w.ChainName)
				blockNumberChan.Stop()
				blockNumberChan = nil
				err := w.ConnectToRpc()
				if err != nil {
					log.Println(err.Error())
				}

			case lastestBlock := <-blockNumberChan.DataTranferChannel:
				lastestEpoch := lastestBlock / fromChain.Config.ChainConfig.EpochLength

				lastestKnownEpochRaw, err := w.RelayHub.GetLatestEpochNumber(nil, big.NewInt(fromChain.Config.ChainConfig.ChainId))
				if err != nil {
					log.Printf("id %v: failed to get GetLatestEpochNumber from %v: %v\n", id, w.ChainName, err)
					continue
				}
				nextUpdatedEpoch := lastestKnownEpochRaw.Int64() + 1
				if nextUpdatedEpoch > lastestEpoch+1 {
					log.Panic("id ", id, ": unexpected epoch")
				}
				if nextUpdatedEpoch >= lastestEpoch {
					w.changeUpdatingStatus(false)
				}
				for nextUpdatedEpoch <= lastestEpoch {
					if nextUpdatedEpoch < lastestEpoch {
						w.changeUpdatingStatus(true)
					}

					log.Printf("id %v: start update validator set from %v to %v, epoch: %v", id, fromChain.ChainName, w.ChainName, nextUpdatedEpoch)
					success, err := w.processUpdateValidatorSet(fromChain, nextUpdatedEpoch, lastestBlock)
					if err != nil {
						log.Printf("id %v: update validator set failed from %v to %v: %v\n", id, fromChain.ChainName, w.ChainName, err)
						break
					}
					if success {
						log.Printf("id %v: successfully updated validator set from %v to %v, epoch: %v\n", id, fromChain.ChainName, w.ChainName, nextUpdatedEpoch)
					}

					nextUpdatedEpoch += 1
				}
			}
		}
	}
}

func (w *UpdateValidatorWorker) betaRelayHubUpdateValidatorSetFromChain(fromChain *UpdateValidatorWorker, id uint64) {
	latestValidatorSet, err := w.RelayHub.GetLatestValidatorSet(nil, big.NewInt(fromChain.Config.ChainConfig.ChainId))
	if err != nil {
		log.Printf("id %v: failed to get GetLatestEpochNumber from %v: %v\n", id, w.ChainName, err)
		return
	}

	lastestKnownEpochRaw, err := w.RelayHub.GetLatestEpochNumber(nil, big.NewInt(fromChain.Config.ChainConfig.ChainId))
	if err != nil {
		log.Printf("id %v: failed to get GetLatestEpochNumber from %v: %v\n", id, w.ChainName, err)
		return
	}
	lastestKnownEpoch := fromChain.Config.OtherConfig.BlockNumber / fromChain.Config.ChainConfig.EpochLength
	if lastestKnownEpoch < lastestKnownEpochRaw.Int64() {
		lastestKnownEpoch = lastestKnownEpochRaw.Int64()
	}
	nextUpdatedEpoch := lastestKnownEpoch + 1

	var blockNumberChan *BlockNumberChannel
	defer func() {
		if blockNumberChan != nil {
			blockNumberChan.Stop()
		}
	}()

MAIN_LOOP_UPDATE_VALIDATOR_BETA:
	for {
		if blockNumberChan == nil {
			blockNumberChan = fromChain.GetBlockNumber()
			if blockNumberChan == nil {
				err := w.ConnectToRpc()
				if err != nil {
					log.Println(err.Error())
					continue MAIN_LOOP_UPDATE_VALIDATOR_BETA
				}
			}
		}
		if blockNumberChan != nil {
			blockNumberChan.Start()
			select {
			case lastestBlock := <-blockNumberChan.DataTranferChannel:
				lastestEpoch := lastestBlock / fromChain.Config.ChainConfig.EpochLength
				if nextUpdatedEpoch >= lastestEpoch {
					w.changeUpdatingStatus(false)
				}
				for nextUpdatedEpoch <= lastestEpoch {
					if nextUpdatedEpoch < lastestEpoch {
						w.changeUpdatingStatus(true)
					}
					log.Printf("id %v: checking validator set from %v to %v, epoch: %v", id, fromChain.ChainName, w.ChainName, nextUpdatedEpoch)

					newValidatorSet, err := utils.GetValidatorSetFromParliaBlock(fromChain.Client, int64(nextUpdatedEpoch)*fromChain.Config.ChainConfig.EpochLength)
					if err != nil {
						log.Printf("id %v: failed to get GetValidatorSetFromEpochBlock from %v: %v\n", id, w.ChainName, err.ErrorString())
						continue
					}
					if !utils.CompareValidatorSet(newValidatorSet, latestValidatorSet) {
						log.Printf("id %v: start update validatorvalidator set from %v to %v, epoch: %v", id, fromChain.ChainName, w.ChainName, nextUpdatedEpoch)
						success, err := w.processUpdateValidatorSet(fromChain, nextUpdatedEpoch, lastestBlock)
						if err != nil {
							log.Printf("id %v: update validator set failed from %v to %v: %v\n", id, fromChain.ChainName, w.ChainName, err)
							latestValidatorSet, err = w.RelayHub.GetLatestValidatorSet(nil, big.NewInt(fromChain.Config.ChainConfig.ChainId))
							if err != nil {
								log.Printf("id %v: failed to get GetLatestEpochNumber from %v: %v\n", id, w.ChainName, err)
							}
							continue
						}
						if success {
							latestValidatorSet = newValidatorSet
							log.Printf("id %v: successfully updated validator set from %v to %v, epoch: %v\n", id, fromChain.ChainName, w.ChainName, nextUpdatedEpoch)
						}
					}
					nextUpdatedEpoch += 1
				}

			case <-blockNumberChan.ErrorChannel:
				blockNumberChan.Stop()
				blockNumberChan = nil
				err := w.ConnectToRpc()
				if err != nil {
					log.Println(err.Error())
				}
			case <-w.stopSignal:
				return
			}
		}
	}
}

func (w *UpdateValidatorWorker) processUpdateValidatorSet(fromChain *UpdateValidatorWorker, epoch, lastestBlock int64) (bool, error) {
	fromBlock := epoch * fromChain.Config.ChainConfig.EpochLength

	newValidatorSet, err := utils.GetValidatorSetFromParliaBlock(fromChain.Client, fromBlock)
	if err != nil {
		return false, fmt.Errorf("failed to generateBlockProofs for get validatorSet from %v: %v", fromChain.ChainName, err.ErrorString())
	}
	if len(newValidatorSet) > 0 {
		newBlockConfirmation := int64(len(newValidatorSet))*2/3 + 1
		if newBlockConfirmation > 0 {
			fromChain.Config.ChainConfig.SetBlockConfirmations(newBlockConfirmation)
		}
	} else {
		return false, fmt.Errorf("failed to get validatorSet from %v: %v", fromChain.ChainName, err.ErrorString())
	}

	toBlock := fromBlock + fromChain.Config.ChainConfig.BlockConfirmations()
	success := false
	if toBlock <= lastestBlock {
		_, blockProofsRaw, err := utils.GenerateBlockProofs(fromChain.Client, int64(fromBlock), int64(toBlock))
		if err != nil {
			return false, fmt.Errorf("failed to generateBlockProofs from %v: %v", fromChain.ChainName, err.ErrorString())
		}

		opts, err := utils.GetTransactOpts(w.Client, w.Signer, big.NewInt(w.Config.ChainConfig.ChainId))
		if err != nil {
			return false, fmt.Errorf("failed to getTransactOpts for %v: %v", w.ChainName, err.ErrorString())
		}

		// update validator set
		log.Printf("updating validator set from %v to %v, epoch: %v", fromChain.ChainName, w.ChainName, epoch)

		tx, errS := w.RelayHub.UpdateValidatorSetUsingEpochBlocks(opts, big.NewInt(fromChain.Config.ChainConfig.ChainId), blockProofsRaw.HeaderRlps)
		if errS != nil {
			return false, fmt.Errorf("failed to UpdateValidatorSetUsingEpochBlocks to %v: %v", w.ChainName, errS)
		}

		err = utils.WaitTxToBeMined(context.TODO(), w.Client, tx.Hash())
		if err != nil {
			return false, fmt.Errorf("failed to get transaction receipt in %v: %v", w.ChainName, err.ErrorString())
		}
		success = true
	}
	return success, nil
}
