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

func (w *Worker) FetchBlockNumber() *BlockNumberChannel {
	channel := make(chan uint64)
	refreshRate := time.NewTicker(w.Config.OtherConfig.Interval * time.Second)
	stopSignal := make(chan struct{})
	blockNumberChannel := &BlockNumberChannel{
		DataTranferChannel: channel,

		stopSignal: stopSignal,
	}
	go func() {
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
				refreshRate.Stop()
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
