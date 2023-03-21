package relayer

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const (
	DEFAULT_USE_CHECKPOINT = true
	DEFAULT_DATA_PATH      = "./"
	DEFAULT_INTERVAL       = 3 * time.Second
	LOG_TOPICS_NUMBER      = 4
	LOG_CHAIN_ID_INDEX     = 1
)

var (
	initialized = false
	AllConfig   *AppConfig
)

func InitConfig(configPath string) {
	if initialized {
		log.Println("Already initialized!")
		return
	}
	initialized = true

	log.Println("Initialize application with config:", configPath)
	jsonFile, err := os.Open(configPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic(err)
		return
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if AllConfig == nil {
		AllConfig = &AppConfig{
			Testnet: Config{
				OtherConfig: OtherConfig{
					Interval:      DEFAULT_INTERVAL,
					DataPath:      DEFAULT_DATA_PATH,
					UseCheckpoint: DEFAULT_USE_CHECKPOINT,
					BlockNumber:   100,
				},
			},
			Mainnet: Config{
				OtherConfig: OtherConfig{
					Interval:      DEFAULT_INTERVAL,
					DataPath:      DEFAULT_DATA_PATH,
					UseCheckpoint: DEFAULT_USE_CHECKPOINT,
					BlockNumber:   100,
				},
			},
			BAS: Config{
				OtherConfig: OtherConfig{
					Interval:      DEFAULT_INTERVAL,
					DataPath:      DEFAULT_DATA_PATH,
					UseCheckpoint: DEFAULT_USE_CHECKPOINT,
					BlockNumber:   100,
				},
			},
			Localhost: Config{
				OtherConfig: OtherConfig{
					Interval:      DEFAULT_INTERVAL,
					DataPath:      DEFAULT_DATA_PATH,
					UseCheckpoint: DEFAULT_USE_CHECKPOINT,
					BlockNumber:   100,
				},
			},
		}
	}
	err = json.Unmarshal(byteValue, AllConfig)
	if err != nil {
		log.Panic(err)
		return
	}

	path := AllConfig.Testnet.OtherConfig.DataPath
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	path = AllConfig.BAS.OtherConfig.DataPath
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	path = AllConfig.Mainnet.OtherConfig.DataPath
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	path = AllConfig.Localhost.OtherConfig.DataPath
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func GetRelayerConfig() *RelayerConfig {
	return &AllConfig.Relayer
}

func GetBSCTestnetConfig() *Config {
	return &AllConfig.Testnet
}

func GetBSCMainnetConfig() *Config {
	return &AllConfig.Mainnet
}

func GetBASNetworkConfig() *Config {
	return &AllConfig.BAS
}

func GetLocalNetworkConfig() *Config {
	return &AllConfig.Localhost
}

type AppConfig struct {
	Relayer   RelayerConfig `json:"relayer"`
	Testnet   Config        `json:"testnet"`
	Mainnet   Config        `json:"mainnet"`
	BAS       Config        `json:"bas"`
	Localhost Config        `json:"localhost"`
}

type Config struct {
	ChainConfig  ChainConfig  `json:"chainConfig"`
	SignerConfig SignerConfig `json:"signerConfig"`
	OtherConfig  OtherConfig  `json:"otherConfig"`
}

type ChainConfig struct {
	ChainName                   string   `json:"chainName"`
	RpcUrl                      []string `json:"rpcUrl"`
	WebsocketRpc                []string `json:"websocketRpc"`
	ChainId                     int64    `json:"chainId"`
	blockConfirmations          int64
	blockConfirmationLock       sync.RWMutex
	EpochLength                 int64  `json:"epochLength"`
	RelayHubAddress             string `json:"relayHubAddress"`
	CrossChainBridgeAddress     string `json:"crossChainBridgeAddress"`
	FilterLogsMaximumBlockRange int64  `json:"filterLogsMaximumBlockRange"`
	EventSignature              string `json:"eventSignature"`
}

func (c *ChainConfig) SetBlockConfirmations(newBlockConfirmation int64) {
	if newBlockConfirmation != c.blockConfirmations {
		c.blockConfirmationLock.Lock()
		defer c.blockConfirmationLock.Unlock()
		c.blockConfirmations = newBlockConfirmation
	}
}

func (c *ChainConfig) BlockConfirmations() int64 {
	c.blockConfirmationLock.RLock()
	defer c.blockConfirmationLock.RUnlock()
	return c.blockConfirmations
}

type SignerConfig struct {
	PrivateKey string `json:"privateKey"`
}

type OtherConfig struct {
	Interval        time.Duration `json:"interval,omitempty"`
	TransactionHash string        `json:"transactionHash"`
	DataPath        string        `json:"dataPath,omitempty"`
	UseCheckpoint   bool          `json:"useCheckpoint,omitempty"`
	BlockNumber     int64         `json:"blockNumber"`
}

func (oc *OtherConfig) UnmarshalJSON(data []byte) error {
	var rawOtherConfig struct {
		Interval        time.Duration `json:"interval,omitempty"`
		TransactionHash string        `json:"transactionHash"`
		DataPath        string        `json:"dataPath,omitempty"`
		UseCheckpoint   bool          `json:"useCheckpoint,omitempty"`
		BlockNumber     int64         `json:"blockNumber"`
	}
	if err := json.Unmarshal(data, &rawOtherConfig); err != nil {
		return err
	}
	rawOtherConfig.Interval *= time.Second
	if rawOtherConfig.Interval <= 0 {
		rawOtherConfig.Interval = DEFAULT_INTERVAL
	}
	*oc = rawOtherConfig
	return nil
}

type RelayerConfig struct {
	Flow struct {
		BasToBsc bool `json:"basToBsc"`
		BscToBas bool `json:"bscToBas"`
	}
	RelayHubType struct {
		BSC string `json:"bsc"`
		BAS string `json:"bas"`
	}
}
