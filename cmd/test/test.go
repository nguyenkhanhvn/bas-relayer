package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	relayer "github.com/nguyenkhanhvn/bas-relayer"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/abi"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/service"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/utils"
	"github.com/nguyenkhanhvn/ethereum-merkle-patricia-trie/trie"
)

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
		// log.Println("kcheck loop: ",i)
		// log.Println("kcheck tx: ",tx.Hash())
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

func TestWithdraw() {
	bscConfig := relayer.GetBSCTestnetConfig()
	bscClient, _ := ethclient.Dial(bscConfig.ChainConfig.RpcUrl)
	// testGenerateBloockProofs(client, config.OtherConfig.BlockNumber, config)

	proofJson, proof, block, receipt, err := utils.GenerateProof(bscClient, bscConfig.ChainConfig.BlockConfirmations, common.HexToHash(bscConfig.OtherConfig.TransactionHash))
	if err != nil {
		panic(err)
	}

	proofJsonData, _ := json.MarshalIndent(proofJson, "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"proof.json", proofJsonData, 0644)

	headerData, _ := json.MarshalIndent(block.Header(), "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"header.json", headerData, 0644)

	receiptData, _ := json.MarshalIndent(receipt, "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"receipt.json", receiptData, 0644)

	basConfig := relayer.GetLocalNetworkConfig()
	basClient, err := ethclient.Dial(basConfig.ChainConfig.RpcUrl)
	if err != nil {
		panic(err)
	}

	basSigner, err := crypto.HexToECDSA(basConfig.SignerConfig.PrivateKey)
	if err != nil {
		panic(err)
	}
	basBridge, err := abi.NewCrossChainBridge(common.HexToAddress(basConfig.ChainConfig.CrossChainBridgeAddress), basClient)
	if err != nil {
		panic(err)
	}
	s := &service.TransferAssetWorker{
		Worker: service.Worker{
			Client: basClient,
			Config: basConfig,
			Signer: basSigner,
		},
		LastestReadedBlock: 0,
		Bridge:             basBridge,
	}
	err = s.Withdraw(proof)
	if err != nil {
		panic(err)
	}
}

func testGoChan() {
	sig := make(chan int)
	i := 1
	refreshRate := time.NewTicker(5 * time.Second)
	go func() {
		for {
			<-refreshRate.C
			sig <- i
			i++
		}
	}()
	go func(chan1 chan int) {
		for {
			data := <-chan1
			log.Println("func1: ", data)
			time.Sleep(7 * time.Second)
		}
	}(sig)
}
