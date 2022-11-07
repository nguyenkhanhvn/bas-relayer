package main

import (
	"log"
	"runtime"
	"sync"

	relayer "github.com/nguyenkhanhvn/bas-relayer"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/service"
)

var mainWaitGroup = sync.WaitGroup{}

func exit() {
	mainWaitGroup.Done()
}

func main() {
	// // create block proofs for registered cross chain bridge (init validator set)
	// config := relayer.GetBSCTestnetConfig()
	// client, _ := ethclient.Dial(config.ChainConfig.RpcUrl)
	// writer.WriteBloockProofs(client, config.OtherConfig.BlockNumber, config)
	// writer.WriteReceiptProofs(client, config)
	// writer.WriteProof(client, config)
	// return
	// test.TestWithdraw()
	// return
	// config := relayer.GetBSCTestnetConfig()
	// config := relayer.GetBSCMainnetConfig()
	// config := relayer.GetLocalNetworkConfig()

	// set maximum number of CPUs that can be executing simultaneously, default is number of physic CPUs
	runtime.GOMAXPROCS(0)
	bscConfig := relayer.GetBSCTestnetConfig()
	// bscConfig := getBSCMainnetConfig()
	basConfig := relayer.GetLocalNetworkConfig()

	mainWaitGroup.Add(1)
	updateValidatorService, err := service.NewUpdateValidatorService(bscConfig, basConfig)
	if err != nil {
		panic(err)
	}

	transferAssetService, err := service.NewTransferAssetService(bscConfig, basConfig)
	if err != nil {
		panic(err)
	}

	go func() {
		if err = updateValidatorService.Start(); err != nil {
			panic(err)
		}
	}()

	go func() {
		for {
			updating := <-updateValidatorService.UpdatingChannel
			if updating {
				log.Println("stop transfer asset to update validators")
				transferAssetService.Stop()
			} else {
				log.Println("start transfer asset")
				if err = transferAssetService.Start(); err != nil {
					panic(err)
				}
			}
		}
	}()

	// wait for exit signal
	mainWaitGroup.Wait()
}
