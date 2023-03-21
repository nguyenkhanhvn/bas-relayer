package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"relayer/pkg/utils"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	relayer "relayer"
)

type Worker struct {
	ChainName string
	Client    *ethclient.Client
	Config    *relayer.Config

	Signer *ecdsa.PrivateKey

	running bool
}

func (w *Worker) start() {
	w.running = true
}

func (w *Worker) isRunning() bool {
	return w.running
}

func (w *Worker) ConnectToRpc() error {
	if len(w.Config.ChainConfig.RpcUrl) <= 0 {
		return fmt.Errorf("ConnectToRpc: url not found")
	}
	var err error
	for _, url := range w.Config.ChainConfig.RpcUrl {
		w.Client, err = ethclient.Dial(url)
		if err == nil {
			w.Signer, err = crypto.HexToECDSA(w.Config.SignerConfig.PrivateKey)
			if err != nil {
				return fmt.Errorf("ConnectToRpc: %v", err)
			}
			return nil
		}
	}
	return fmt.Errorf("ConnectToRpc: cannot connect to any url")
}

func (w *Worker) GetBlockNumber() *BlockNumberChannel {
	if w.Client == nil {
		return nil
	}
	channel := make(chan int64)
	errorChannel := make(chan utils.ErrorType)
	stopSignal := make(chan struct{})
	blockNumberChannel := &BlockNumberChannel{
		Running:            false,
		DataTranferChannel: channel,
		ErrorChannel:       errorChannel,
		stopSignal:         stopSignal,
		worker:             w,
	}
	return blockNumberChannel
}

type BlockNumberChannel struct {
	Running            bool
	DataTranferChannel chan int64
	ErrorChannel       chan utils.ErrorType
	stopSignal         chan struct{}
	worker             *Worker
}

func (c *BlockNumberChannel) Start() {
	if c.Running {
		return
	}
	c.Running = true

	go func() {
		refreshRate := time.NewTicker(c.worker.Config.OtherConfig.Interval)
		defer func() {
			refreshRate.Stop()
			c.Running = false
		}()

		for {
			select {
			case <-c.stopSignal:
				return
			case <-refreshRate.C:
				blockNumber, err := c.worker.Client.BlockNumber(context.TODO())
				if err != nil {
					log.Printf("failed to get BlockNumber from %v: %v\n", c.worker.ChainName, err)
					c.ErrorChannel <- utils.ClientError
					return
				} else {
					c.DataTranferChannel <- int64(blockNumber)
				}
			}
		}
	}()
}

func (c *BlockNumberChannel) Stop() {
	if !c.Running {
		return
	}

	go func() {
		c.stopSignal <- struct{}{}
	}()

	go func() {
		// process all exist signal
		for {
			select {
			case <-c.DataTranferChannel:
				log.Printf("flush BlockNumberChannel DataTranferChannel")
			default:
				return
			}
		}
	}()
}
