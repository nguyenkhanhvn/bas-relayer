package main

import (
	"log"
	"os"
	"runtime"
	"sync"

	"relayer"
	"relayer/pkg/flags"
	"relayer/pkg/service"

	"github.com/urfave/cli"
)

var (
	mainWaitGroup = sync.WaitGroup{}
	app           = cli.NewApp()
)

func exit() {
	mainWaitGroup.Done()
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	app.Flags = append(app.Flags, flags.ConfigFlag)
	app.Action = action
}

func action(c *cli.Context) error {
	relayer.InitConfig(c.String(flags.ConfigFlag.Name))

	relayerConfig := relayer.GetRelayerConfig()
	// config := relayer.GetBSCTestnetConfig()
	// config := relayer.GetBSCMainnetConfig()
	// config := relayer.GetLocalNetworkConfig()
	// config := relayer.GetBASNetworkConfig()

	bscConfig := relayer.GetBSCTestnetConfig()
	// bscConfig := relayer.GetBSCMainnetConfig()
	// localConfig := relayer.GetLocalNetworkConfig()
	basConfig := relayer.GetBASNetworkConfig()

	service, err := service.NewService(relayerConfig, bscConfig, basConfig)
	if err != nil {
		log.Panic(err)
	}

	err = service.Start()
	if err != nil {
		log.Panic(err)
	}
	defer service.Stop()

	mainWaitGroup.Add(1)
	// wait for exit
	mainWaitGroup.Wait()
	return nil
}

func main() {
	// set maximum number of CPUs that can be executing simultaneously, default is number of physic CPUs
	runtime.GOMAXPROCS(0)

	err := app.Run(os.Args)
	if err != nil {
		log.Panic(err)
	}
}
