package flags

import (
	"os"

	"github.com/urfave/cli"
)

const (
	DEFAULT_CONFIG = "config.json"
)

var (
	Network = cli.StringFlag{
		Name:  "network",
		Value: "bas",
		Usage: "network to connect",
	}
	ConfigFlag = cli.StringFlag{
		Name:  "config",
		Usage: "Path to config file",
	}
)

func init() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ConfigFlag.Value = currentDirectory + "/" + DEFAULT_CONFIG
}
