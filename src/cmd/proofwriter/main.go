package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"

	"relayer"
	"relayer/pkg/flags"
	"relayer/pkg/utils"
	"relayer/pkg/writer"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
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

	app.Flags = []cli.Flag{
		flags.Network,
	}
	app.Action = action
}

func action(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("missing argument")
	} else if c.NArg() == 1 {
		return doTask(c.Args()[0], c.String("network"))
	}
	return fmt.Errorf("too many argument")
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Panic(err)
	}
}

func doTask(task string, network string) error {
	log.Printf("Do task: %v, in network: %v", task, network)

	var config *relayer.Config
	switch network {
	case "bas":
		config = relayer.GetBASNetworkConfig()
	case "mainnet":
		config = relayer.GetBSCMainnetConfig()
	case "testnet":
		config = relayer.GetBSCTestnetConfig()
	case "local":
		config = relayer.GetLocalNetworkConfig()
	default:
		config = relayer.GetLocalNetworkConfig()
	}
	var client *ethclient.Client
	for _, rpcUrl := range config.ChainConfig.RpcUrl {
		var err error
		client, err = ethclient.Dial(rpcUrl)
		if err == nil {
			break
		}
	}
	switch task {
	case "genblockproofs":
		writer.WriteBloockProofs(client, config.OtherConfig.BlockNumber, config)
	case "genreceiptproofs":
		writer.WriteReceiptProofs(client, config)
	case "genproof":
		writer.WriteProof(client, config)
	case "genheader":
		block, err := client.BlockByNumber(context.TODO(), big.NewInt(config.OtherConfig.BlockNumber))
		if err != nil {
			return err
		}
		headerRlp, _ := rlp.EncodeToBytes(block.Header())
		header := struct {
			Header    *types.Header
			HeaderRlp string
		}{
			Header:    block.Header(),
			HeaderRlp: utils.HexToString(headerRlp),
		}
		headerData, _ := json.MarshalIndent(header, "", "	")
		_ = os.WriteFile(config.OtherConfig.DataPath+"header.json", headerData, 0644)

	default:
		return fmt.Errorf("command not support")
	}

	return nil
}
