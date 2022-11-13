package service

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	relayer "github.com/nguyenkhanhvn/bas-relayer"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/abi"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/utils"
)

type UpdateValidatorWorker struct {
	Worker

	RelayHub        *abi.RelayHub
	updating        bool
	updatingChannel chan bool
	stopSignal      chan struct{}
}

func NewUpdateValidatorWorker(chainName string, client *ethclient.Client, config *relayer.Config, signer *ecdsa.PrivateKey, relayHub *abi.RelayHub) *UpdateValidatorWorker {
	return &UpdateValidatorWorker{
		Worker: Worker{
			ChainName: chainName,
			Client:    client,
			Config:    config,
			Signer:    signer,
		},
		RelayHub:        relayHub,
		updating:        false,
		updatingChannel: make(chan bool),
		stopSignal:      make(chan struct{}),
	}
}

func (w *UpdateValidatorWorker) Stop() {
	w.stopSignal <- struct{}{}
}

// return updating status
func (w *UpdateValidatorWorker) Updating() bool {
	return w.updating
}

// return channel which will emit when updateing status change
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

func (w *UpdateValidatorWorker) UpdateValidatorSetFromChain(fromChain *UpdateValidatorWorker) {
	log.Printf("start UpdateValidatorSet from %v to %v\n", fromChain.ChainName, w.ChainName)
	blockNumberChan := fromChain.GetBlockNumber()
	defer blockNumberChan.Stop()
	for {
		select {
		case lastestBlock := <-blockNumberChan.DataTranferChannel:
			lastestEpoch := lastestBlock / fromChain.Config.ChainConfig.EpochLength
			lastestKnownEpochRaw, err := w.RelayHub.GetLatestEpochNumber(nil, big.NewInt(fromChain.Config.ChainConfig.ChainId))
			if err != nil {
				log.Printf("UpdateValidatorSet: failed to get GetLatestEpochNumber from %v: %v\n", w.ChainName, err)
				continue
			}
			nextUpdatedEpoch := lastestKnownEpochRaw.Uint64() + 1
			if nextUpdatedEpoch > lastestEpoch+1 {
				panic("unexpected epoch")
			}
			if nextUpdatedEpoch >= lastestEpoch {
				w.changeUpdatingStatus(false)
			}
			for nextUpdatedEpoch <= lastestEpoch {
				if nextUpdatedEpoch < lastestEpoch {
					w.changeUpdatingStatus(true)
				}

				fromBlock := nextUpdatedEpoch * fromChain.Config.ChainConfig.EpochLength
				toBlock := fromBlock + uint64(fromChain.Config.ChainConfig.BlockConfirmations)
				if toBlock <= lastestBlock {
					_, blockProofs, err := utils.GenerateBlockProofs(fromChain.Client, int64(fromBlock), int64(toBlock))
					if err != nil {
						log.Printf("UpdateValidatorSet: failed to generateBlockProofs from %v: %v\n", fromChain.ChainName, err)
						break
					}

					opts, err := utils.GetTransactOpts(w.Client, w.Signer, big.NewInt(w.Config.ChainConfig.ChainId))
					if err != nil {
						log.Printf("UpdateValidatorSet: failed to getTransactOpts for %v: %v\n", w.ChainName, err)
						break
					}

					// update validator set
					log.Printf("updateing validator set from %v to %v, epoch: %v\n", fromChain.ChainName, w.ChainName, nextUpdatedEpoch)
					tx, err := w.RelayHub.UpdateValidatorSetUsingEpochBlocks(opts, big.NewInt(fromChain.Config.ChainConfig.ChainId), blockProofs.HeaderRlps)
					if err != nil {
						log.Printf("UpdateValidatorSet: failed to UpdateValidatorSetUsingEpochBlocks to %v: %v\n", w.ChainName, err)
						break
					}

					err = utils.WaitTxToBeMined(context.TODO(), w.Client, tx.Hash())
					if err != nil {
						log.Printf("UpdateValidatorSet: failed to get transaction receipt in %v: %v\n", w.ChainName, err)
						break
					}
				}

				nextUpdatedEpoch += 1
			}
		case <-w.stopSignal:
			w.changeUpdatingStatus(false)
			return
		}
	}
}

func (w *UpdateValidatorWorker) GetLatestEpoch() (*big.Int, error) {
	return w.RelayHub.GetLatestEpochNumber(nil, big.NewInt(w.Config.ChainConfig.ChainId))
}
