package service

import (
	"context"
	"crypto/ecdsa"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	relayer "github.com/nguyenkhanhvn/bas-relayer"
)

type Worker struct {
	ChainName string
	Client    *ethclient.Client
	Config    *relayer.Config

	Signer *ecdsa.PrivateKey
}

func (w *Worker) GetBlockNumber() *BlockNumberChannel {
	channel := make(chan uint64)
	stopSignal := make(chan struct{})
	blockNumberChannel := &BlockNumberChannel{
		DataTranferChannel: channel,
		stopSignal:         stopSignal,
	}
	go func() {
		refreshRate := time.NewTicker(w.Config.OtherConfig.Interval * time.Second)
		defer refreshRate.Stop()
		for {
			select {
			case <-refreshRate.C:
				blockNumber, err := w.Client.BlockNumber(context.TODO())
				if err != nil {
					log.Printf("FetchBlockNumber: failed to get BlockNumber from %v: %v\n", w.ChainName, err)
					continue
				}
				channel <- blockNumber
			case <-stopSignal:
				return
			}
		}
	}()
	return blockNumberChannel
}

type BlockNumberChannel struct {
	DataTranferChannel <-chan uint64
	stopSignal         chan<- struct{}
}

func (c *BlockNumberChannel) Stop() {
	c.stopSignal <- struct{}{}
}
