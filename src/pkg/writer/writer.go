package writer

import (
	"context"
	"encoding/json"
	"log"
	"os"

	relayer "relayer"
	"relayer/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WriteBloockProofs(client *ethclient.Client, blockNumber int64, config *relayer.Config) {
	if client == nil || config == nil {
		return
	}

	blockProofs, _, errData := utils.GenerateBlockProofs(client, blockNumber, blockNumber+config.ChainConfig.BlockConfirmations())
	if errData != nil {
		panic(errData)
	}
	blockProofsData, _ := json.MarshalIndent(blockProofs, "", "	")
	err := os.WriteFile(config.OtherConfig.DataPath+"blockProofs.json", blockProofsData, 0644)
	if err != nil {
		log.Println("error when WriteBloockProofs:", err)
	} else {
		log.Println("WriteBloockProofs successed at", config.OtherConfig.DataPath+"blockProofs.json")
	}
}

func WriteReceiptProofs(client *ethclient.Client, config *relayer.Config) {
	if client == nil || config == nil {
		return
	}

	receiptProof, _, block, receipt, _, errData := utils.GenerateReceiptProof(client, common.HexToHash(config.OtherConfig.TransactionHash))
	if errData != nil {
		panic(errData)
	}
	receiptProofData, _ := json.MarshalIndent(receiptProof, "", "	")
	err := os.WriteFile(config.OtherConfig.DataPath+"receiptProof.json", receiptProofData, 0644)
	if err != nil {
		log.Println("error when WriteReceiptProofs:", err)
	}

	headerData, _ := json.MarshalIndent(block.Header(), "", "	")
	err = os.WriteFile(config.OtherConfig.DataPath+"header.json", headerData, 0644)
	if err != nil {
		log.Println("error when WriteReceiptProofs:", err)
	}

	receiptData, _ := json.MarshalIndent(receipt, "", "	")
	err = os.WriteFile(config.OtherConfig.DataPath+"receipt.json", receiptData, 0644)

	if err != nil {
		log.Println("error when WriteReceiptProofs:", err)
	} else {
		log.Println("WriteReceiptProofs successed at", config.OtherConfig.DataPath)
	}
}

func WriteProof(client *ethclient.Client, config *relayer.Config) {
	if client == nil || config == nil {
		return
	}

	proofJson, _, block, receipt, errData := utils.GenerateProof(client, config.ChainConfig.BlockConfirmations(), common.HexToHash(config.OtherConfig.TransactionHash))
	if errData != nil {
		panic(errData)
	}
	proofJsonData, _ := json.MarshalIndent(proofJson, "", "	")
	err := os.WriteFile(config.OtherConfig.DataPath+"proof.json", proofJsonData, 0644)
	if err != nil {
		log.Println("error when WriteProof:", err)
	}

	headerData, _ := json.MarshalIndent(block.Header(), "", "	")
	err = os.WriteFile(config.OtherConfig.DataPath+"header.json", headerData, 0644)
	if err != nil {
		log.Println("error when WriteProof:", err)
	}

	receiptData, _ := json.MarshalIndent(receipt, "", "	")
	err = os.WriteFile(config.OtherConfig.DataPath+"receipt.json", receiptData, 0644)

	if err != nil {
		log.Println("error when WriteProof:", err)
	} else {
		log.Println("WriteProof successed at", config.OtherConfig.DataPath)
	}
}

func WriteAllProofs(config *relayer.Config, client *ethclient.Client) {
	if client == nil || config == nil {
		return
	}

	WriteBloockProofs(client, config.OtherConfig.BlockNumber, config)
	WriteReceiptProofs(client, config)
	WriteProof(client, config)

	latestBlockNumber, err := client.BlockNumber(context.TODO())
	if err != nil {
		panic(err)
	}
	log.Println("Latest block: ", latestBlockNumber)

	var latestEpochBlockNumber int64
	if config.OtherConfig.UseCheckpoint {
		latestEpochBlockNumber = config.OtherConfig.BlockNumber
	} else {
		latestEpochBlockNumber = (int64(latestBlockNumber) / config.ChainConfig.EpochLength) * config.ChainConfig.EpochLength
	}
	// create newest epoch block for register cross chain relay hub
	// create checkpoint epoch block for register cross chain relay hub
	if int64(latestBlockNumber)/config.ChainConfig.EpochLength == latestEpochBlockNumber/config.ChainConfig.EpochLength && int64(latestBlockNumber)%config.ChainConfig.EpochLength < config.ChainConfig.BlockConfirmations() {
		latestEpochBlockNumber -= config.ChainConfig.EpochLength
	}
	if latestEpochBlockNumber%config.ChainConfig.EpochLength != 0 {
		panic("bad epoch")
	}
	log.Println("Epoch block: ", latestEpochBlockNumber)
	WriteBloockProofs(client, int64(latestEpochBlockNumber), config)

	// create proof for withdraw
	WriteProof(client, config)

}
