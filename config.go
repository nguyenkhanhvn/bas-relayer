package relayer

import "time"

const (
	USE_CHECKPOINT = true
	DATA_PATH      = "../data/"
	INTERVAL       = 3 // in second
)

func GetBSCTestnetConfig() *Config {
	return &Config{
		ChainConfig: ChainConfig{
			ChainName:               "BSC Testnet",
			RpcUrl:                  "https://data-seed-prebsc-2-s1.binance.org:8545/",
			ChainId:                 97,
			BlockConfirmations:      7,
			EpochLength:             200,
			RelayHubAddress:         "0x33366db4C1C593B3d83c0B3532548b31112748DD",
			CrossChainBridgeAddress: "0x8e92275FA0D61C6B039688cDe252dEbfF0ab11eE",
			EventSignature:          "DepositToken(uint256,uint256,address,address,address,address,uint256,(string,string,uint256,address))",
		},
		SignerConfig: SignerConfig{
			PrivateKey: "dbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
		},
		OtherConfig: OtherConfig{
			Interval:        INTERVAL,
			TransactionHash: "0x8dc2796d81e5c729fc5b15a630999bb513dee75fa55f05a12138c9a14c9c083e",
			DataPath:        DATA_PATH,
			UseCheckpoint:   USE_CHECKPOINT,
			BlockNumber:     24264200,
		},
	}
}

func GetBSCMainnetConfig() *Config {
	return &Config{
		ChainConfig: ChainConfig{
			ChainName:               "BSC Mainnet",
			RpcUrl:                  "https://bsc-dataseed1.defibit.io",
			ChainId:                 56,
			BlockConfirmations:      150,
			EpochLength:             200,
			RelayHubAddress:         "",
			CrossChainBridgeAddress: "",
			EventSignature:          "DepositToken(uint256,uint256,address,address,address,address,uint256,(string,string,uint256,address))",
		},
		SignerConfig: SignerConfig{
			PrivateKey: "dbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
		},
		OtherConfig: OtherConfig{
			Interval:        INTERVAL,
			TransactionHash: "",
			DataPath:        DATA_PATH,
			UseCheckpoint:   USE_CHECKPOINT,
			BlockNumber:     22095750,
		},
	}
}

func GetLocalNetworkConfig() *Config {
	return &Config{
		ChainConfig: ChainConfig{
			ChainName:               "Local network",
			RpcUrl:                  "http://172.16.0.150:8545/",
			WebsocketRpc:            "ws://127.0.0.1:8545",
			ChainId:                 14001,
			BlockConfirmations:      3,
			EpochLength:             1200,
			RelayHubAddress:         "0xF901cC9dc5E1fDE56ED38740492Aeec00324b195",
			CrossChainBridgeAddress: "0xDd09C3cB257D99eb41A99784Ef5d1fc1B323dF8E",
			EventSignature:          "DepositToken(uint256,uint256,address,address,address,address,uint256,(string,string,uint256,address))",
		},
		SignerConfig: SignerConfig{
			PrivateKey: "4cb95630b02e4a1b3fa28da9138cda1c62677d121a324d025f00cec6a5610828",
		},
		OtherConfig: OtherConfig{
			Interval:        INTERVAL,
			TransactionHash: "0x76359a01c9c94dd29824bbc3ba34c1eb8eae124148795b77155bb13b6ef7b87d",
			DataPath:        DATA_PATH,
			UseCheckpoint:   USE_CHECKPOINT,
			BlockNumber:     87600,
		},
	}
}

type Config struct {
	ChainConfig  ChainConfig
	SignerConfig SignerConfig
	OtherConfig  OtherConfig
}

type ChainConfig struct {
	ChainName               string
	RpcUrl                  string
	WebsocketRpc            string
	ChainId                 int64
	BlockConfirmations      int64
	EpochLength             uint64
	RelayHubAddress         string
	CrossChainBridgeAddress string
	EventSignature          string
}

type SignerConfig struct {
	PrivateKey string
}

type OtherConfig struct {
	Interval        time.Duration
	TransactionHash string
	DataPath        string
	UseCheckpoint   bool
	BlockNumber     int64
}
