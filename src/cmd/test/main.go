package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	relayer "relayer"
	"relayer/pkg/abi"
	"relayer/pkg/flags"
	"relayer/pkg/service"
	"relayer/pkg/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/nguyenkhanhvn/ethereum-merkle-patricia-trie/trie"
	"github.com/urfave/cli"
)

var (
	app = cli.NewApp()
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	app.Flags = append(app.Flags, flags.ConfigFlag)
	app.Action = action
}

func action(c *cli.Context) error {
	relayer.InitConfig(c.String(flags.ConfigFlag.Name))

	if c.NArg() == 0 {
		return fmt.Errorf("missing argument")
	} else if c.NArg() > 0 {
		task := c.Args()[0]
		log.Printf("Do task: %v", task)

		bscConfig := relayer.GetBSCTestnetConfig()
		bscClient, err := ethclient.Dial(bscConfig.ChainConfig.RpcUrl[0])
		if err != nil {
			panic(err)
		}

		basConfig := relayer.GetBASNetworkConfig()
		basClient, err := ethclient.Dial(basConfig.ChainConfig.RpcUrl[0])
		if err != nil {
			panic(err)
		}

		switch task {
		case "withdraw":
			TestWithdraw(bscConfig, bscClient, basConfig, basClient)
			return nil
		case "transferasset":
			client1 := basClient
			config1 := basConfig
			client2 := bscClient
			config2 := bscConfig
			if c.NArg() > 1 {
				switch c.Args()[1] {
				case "bsc":
					client1 = bscClient
					config1 = bscConfig
					client2 = basClient
					config2 = basConfig
				}
			}
			TransferAsset(config1, client1, config2, client2)
			return nil
		case "checkvalidatorset":
			client1 := basClient
			config1 := basConfig
			client2 := bscClient
			config2 := bscConfig
			if c.NArg() > 1 {
				switch c.Args()[1] {
				case "bsc":
					client1 = bscClient
					config1 = bscConfig
					client2 = basClient
					config2 = basConfig
				}
			}

			set1, errData := utils.GetValidatorSetFromParliaBlock(client1, 2*config1.ChainConfig.EpochLength)
			if errData != nil {
				panic(errData)
			}
			log.Println("config2 relayhub address:", config2.ChainConfig.RelayHubAddress)
			relayHub, err := abi.NewRelayHubBeta(common.HexToAddress(config2.ChainConfig.RelayHubAddress), client2)
			if err != nil {
				panic(err)
			}

			log.Println("config1 chainid:", config1.ChainConfig.ChainId)
			set2, err := relayHub.GetLatestValidatorSet(nil, big.NewInt(config1.ChainConfig.ChainId))
			if err != nil {
				panic(err)
			}

			log.Println("set1 before")
			for _, v := range set1 {
				log.Println(v)
			}
			log.Println("set2 before")
			for _, v := range set2 {
				log.Println(v)
			}

			log.Println("compare:", utils.CompareValidatorSet(set1, set2))

			epoch, err := relayHub.GetLatestEpochNumber(nil, big.NewInt(config1.ChainConfig.ChainId))
			if err != nil {
				panic(err)
			}
			log.Println("latestEpoch:", epoch)
			return nil
		case "test":
			client1 := basClient
			config1 := basConfig
			// client2 := bscClient
			config2 := bscConfig
			if c.NArg() > 1 {
				switch c.Args()[1] {
				case "bsc":
					client1 = bscClient
					config1 = bscConfig
					// client2 = basClient
					config2 = basConfig
				}
			}

			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(126690 * 200),
				ToBlock:   big.NewInt(126710 * 200),
				Addresses: []common.Address{common.HexToAddress(config1.ChainConfig.CrossChainBridgeAddress)},
				Topics:    [][]common.Hash{{crypto.Keccak256Hash([]byte(config1.ChainConfig.EventSignature))}},
			}

			logs, err := client1.FilterLogs(context.TODO(), query)
			if err != nil {
				panic(err)
			}

			log.Printf("logs size: %v\n", len(logs))
			for _, vlog := range logs {
				if len(vlog.Topics) >= relayer.LOG_TOPICS_NUMBER && vlog.Topics[relayer.LOG_CHAIN_ID_INDEX].Big().Int64() == config2.ChainConfig.ChainId {
					log.Printf("log: %v\n", vlog.Topics)
				}
			}
			return nil
		default:
			return fmt.Errorf("command not support")
		}
	}
	return fmt.Errorf("invalid number of argument")
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Panic(err)
	}
}

func TransferAsset(srcConfig *relayer.Config, srcClient *ethclient.Client, destConfig *relayer.Config, destClient *ethclient.Client) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(126690 * 200),
		ToBlock:   big.NewInt(126710 * 200),
		Addresses: []common.Address{common.HexToAddress(srcConfig.ChainConfig.CrossChainBridgeAddress)},
		Topics:    [][]common.Hash{{crypto.Keccak256Hash([]byte(srcConfig.ChainConfig.EventSignature))}},
	}

	logs, err := srcClient.FilterLogs(context.TODO(), query)
	if err != nil {
		log.Printf("failed to FilterLogs: %v\n", err)
	}

	log.Printf("logs size: %v\n", len(logs))
	for _, vlog := range logs {
		log.Printf("start withdraw, txHash: %v, waiting for confirmation...\n", vlog.TxHash)
		// make proof at source chain
		blockNumber, _, err := utils.GetBlockNumberFromTransactionHash(srcClient, vlog.TxHash)
		if err != nil {
			log.Printf("cannot getBlockNumberFromTransactionHash, error: %v", err)
		}

		_, proofRaw, _, _, err := utils.GenerateProofWithoutCheck(srcClient, blockNumber.Int64(), blockNumber.Int64()+srcConfig.ChainConfig.BlockConfirmations(), vlog.TxHash)
		if err != nil {
			log.Printf("failed to generateProof, txHash: %v, error: %v\n", vlog.TxHash, err)
			continue
		}

		// withdraw at destination chain
		_, errData := Withdraw(destConfig, destClient, proofRaw)
		if errData != nil {
			log.Printf("failed to withdraw, txHash: %v, error: %v\n", vlog.TxHash, errData)
			continue
		}
		log.Println("finish withdraw: ", vlog.TxHash)
	}
}

func Withdraw(config *relayer.Config, client *ethclient.Client, proof *utils.ProofRaw) (*types.Transaction, utils.Error) {
	signer, err := crypto.HexToECDSA(config.SignerConfig.PrivateKey)
	if err != nil {
		return nil, utils.NewError(utils.OtherError, err)
	}
	bridge, err := abi.NewCrossChainBridge(common.HexToAddress(config.ChainConfig.CrossChainBridgeAddress), client)
	if err != nil {
		panic(err)
	}
	s := &service.TransferAssetWorker{
		Worker: service.Worker{
			Client: client,
			Config: config,
			Signer: signer,
		},
		LastestReadedBlock: make(map[*service.TransferAssetWorker]int64),
		Bridge:             bridge,
	}
	return s.Withdraw(proof)
}

func GetData(blockNumber int64) (*types.Header, []byte, []byte, []byte, trie.Proof, [][]byte, *types.Receipt, *trie.Trie) {
	client, _ := ethclient.Dial("https://bsc-dataseed1.defibit.io")
	// client, _ := ethclient.Dial("https://data-seed-prebsc-2-s1.binance.org:8545/")
	block, _ := client.BlockByNumber(context.TODO(), big.NewInt(blockNumber))

	// transactionsTrie := NewTrie()
	receiptsTrie := trie.NewTrie()

	var rawReceipt, rKey, keyNibble []byte
	var proof trie.Proof
	var proofData [][]byte
	var choseRp *types.Receipt
	log.Println("total transactions: ", block.Transactions().Len())
	for i, tx := range block.Transactions() {
		// key is the encoding of the index as the unsigned integer type
		key, _ := rlp.EncodeToBytes(uint(i))

		// value is the RLP encoding of a transaction
		// transactionRlp, _ := rlp.EncodeToBytes(tx)
		// transactionsTrie.Put(key, transactionRlp)

		rp, _ := client.TransactionReceipt(context.TODO(), tx.Hash())

		// value is the RLP encoding of a transaction
		receiptRlp, _ := rlp.EncodeToBytes(rp)
		receiptsTrie.Put(key, receiptRlp)

		if i == 0 {
			choseRp = rp
			rawReceipt = receiptRlp
			rKey = key
			file, _ := json.MarshalIndent(rp, "", "	")
			_ = os.WriteFile("testData/receipt.json", file, 0644)
		}
	}
	proof, _ = receiptsTrie.Prove(rKey)
	proofData, keyNibble, _ = receiptsTrie.CreateProof(rKey)

	log.Println()
	log.Println()
	log.Println()
	log.Println()
	hash, _ := receiptsTrie.Hash()
	log.Printf("receiptsTrieRoot: 0x%x\n", hash)
	log.Println("key: ", rKey)
	return block.Header(), rawReceipt, rKey, keyNibble, proof, proofData, choseRp, receiptsTrie
}

type CustomData struct {
	DataInt     uint64
	DataBigInt  *big.Int
	DataAddress *common.Address
	DataBloom   types.Bloom
	DataBytes   []byte
}

type CustomJson struct {
	Data      CustomData
	DataBytes string
	Raw       string
}

func Test() {
	var blockNumber int64 = 2000000
	header, rawReceipt, key, keyNibble, proof, proofData, receipt, _ := GetData(blockNumber)
	log.Println()
	log.Println()
	log.Println()
	log.Println(header.ReceiptHash)
	log.Println()
	log.Println(rawReceipt)
	log.Println()
	log.Println()
	log.Println("this is key")
	log.Println(key)
	log.Println(keyNibble)
	log.Println()
	log.Println()
	proofRlp, _ := rlp.EncodeToBytes(proofData)
	var proofString []string
	for _, node := range proofData {
		proofString = append(proofString, utils.HexToString(node))
	}
	log.Println()
	log.Println()
	log.Println()
	log.Println()
	log.Println()
	val, err := trie.VerifyProof(header.ReceiptHash.Bytes(), key, proof)
	if err != nil {
		log.Println(err)
	} else if val != nil {
		log.Println(val)
		log.Println()
		log.Println()
		log.Println()
		log.Println()
		headerRlp, _ := rlp.EncodeToBytes(header)

		fileHeader, _ := json.MarshalIndent(header, "", "	")
		_ = os.WriteFile("testData/header.json", fileHeader, 0644)
		receiptProofJson := utils.ReceiptProofJson{
			Header:      utils.HexToString(headerRlp),
			RawReceipt:  utils.HexToString(rawReceipt),
			ProofPath:   utils.HexToString(keyNibble),
			Proof:       proofString,
			ProofRlp:    utils.HexToString(proofRlp),
			ReceiptRoot: utils.HexToString(header.ReceiptHash.Bytes()),
		}

		file, _ := json.MarshalIndent(receiptProofJson, "", " ")
		path := "testData/test" + fmt.Sprint(blockNumber) + ".json"
		_ = os.WriteFile(path, file, 0644)
		_ = os.WriteFile("../../scripts/test.json", file, 0644)

		// dataInt,_:=rlp.EncodeToBytes(1099)
		// dataBigInt,_:=rlp.EncodeToBytes(big.NewInt(19999999999))
		addressValue := common.HexToAddress("0x193be62f1f40766f7e535b8f09063ad5345c4f96")
		bigIntValue, _ := new(big.Int).SetString("2188824200011112223", 10)
		// dataAddress,_:=rlp.EncodeToBytes(common.HexToAddress("0x193be62f1f40766f7e535b8f09063ad5345c4f96"))
		// dataBytes,_:=rlp.EncodeToBytes([]byte("Here is a string...."))
		c := CustomData{
			DataInt:     1099,
			DataBigInt:  bigIntValue,
			DataAddress: &addressValue,
			DataBloom:   receipt.Bloom,
			DataBytes:   []byte("Here is a string...."),
		}
		cRLP, _ := rlp.EncodeToBytes(c)
		cJ := CustomJson{
			Data:      c,
			DataBytes: utils.HexToString(c.DataBytes),
			Raw:       utils.HexToString(cRLP),
		}
		fileCustom, _ := json.MarshalIndent(cJ, "", " ")
		_ = os.WriteFile("testData/custom.json", fileCustom, 0644)

	} else {
		log.Println("unexpected")
	}

	log.Println("receipt 		: ", utils.HexToString(rawReceipt))
	log.Println("receipt hash	: ", common.BytesToHash(rawReceipt))

}

func TestWithdraw(srcConfig *relayer.Config, srcClient *ethclient.Client, destConfig *relayer.Config, destClient *ethclient.Client) {
	// testGenerateBloockProofs(client, config.OtherConfig.BlockNumber, config)

	proofJson, proof, block, receipt, err := utils.GenerateProof(srcClient, srcConfig.ChainConfig.BlockConfirmations(), common.HexToHash(srcConfig.OtherConfig.TransactionHash))
	if err != nil {
		panic(err)
	}

	proofJsonData, _ := json.MarshalIndent(proofJson, "", "	")
	_ = os.WriteFile(srcConfig.OtherConfig.DataPath+"proof.json", proofJsonData, 0644)

	headerData, _ := json.MarshalIndent(block.Header(), "", "	")
	_ = os.WriteFile(srcConfig.OtherConfig.DataPath+"header.json", headerData, 0644)

	receiptData, _ := json.MarshalIndent(receipt, "", "	")
	_ = os.WriteFile(srcConfig.OtherConfig.DataPath+"receipt.json", receiptData, 0644)

	if err != nil {
		panic(err)
	}

	_, err = Withdraw(destConfig, destClient, proof)
	if err != nil {
		panic(err)
	}
}
