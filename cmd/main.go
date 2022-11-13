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
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	bscConfig := relayer.GetBSCTestnetConfig()
	// bscConfig := relayer.GetBSCMainnetConfig()
	basConfig := relayer.GetLocalNetworkConfig()

	service, err := service.NewService(bscConfig, basConfig)
	if err != nil {
		panic(err)
	}

	mainWaitGroup.Add(1)
	err = service.Start()
	if err != nil {
		panic(err)
	}

	// wait for exit
	mainWaitGroup.Wait()
	service.Stop()
}
