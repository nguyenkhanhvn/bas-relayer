package writer

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	relayer "github.com/nguyenkhanhvn/bas-relayer"
	"github.com/nguyenkhanhvn/bas-relayer/pkg/utils"
)

func WriteBloockProofs(client *ethclient.Client, blockNumber int64, config *relayer.Config) {
	blockProofs, _, err := utils.GenerateBlockProofs(client, blockNumber, blockNumber+config.ChainConfig.BlockConfirmations)
	if err != nil {
		panic(err)
	}
	blockProofsData, _ := json.MarshalIndent(blockProofs, "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"blockProofs.json", blockProofsData, 0644)
}

func WriteReceiptProofs(client *ethclient.Client, config *relayer.Config) {
	receiptProof, _, block, receipt, _, err := utils.GenerateReceiptProof(client, common.HexToHash(config.OtherConfig.TransactionHash))
	if err != nil {
		panic(err)
	}
	receiptProofData, _ := json.MarshalIndent(receiptProof, "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"receiptProof.json", receiptProofData, 0644)

	headerData, _ := json.MarshalIndent(block.Header(), "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"header.json", headerData, 0644)

	receiptData, _ := json.MarshalIndent(receipt, "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"receipt.json", receiptData, 0644)

}

func WriteProof(client *ethclient.Client, config *relayer.Config) {
	proofJson, _, block, receipt, err := utils.GenerateProof(client, config.ChainConfig.BlockConfirmations, common.HexToHash(config.OtherConfig.TransactionHash))

	if err != nil {
		panic(err)
	}
	proofJsonData, _ := json.MarshalIndent(proofJson, "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"proof.json", proofJsonData, 0644)

	headerData, _ := json.MarshalIndent(block.Header(), "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"header.json", headerData, 0644)

	receiptData, _ := json.MarshalIndent(receipt, "", "	")
	_ = os.WriteFile(relayer.DATA_PATH+"receipt.json", receiptData, 0644)

}

func WriteAllProofs(config *relayer.Config, client *ethclient.Client) {
	WriteBloockProofs(client, config.OtherConfig.BlockNumber, config)
	WriteReceiptProofs(client, config)
	WriteProof(client, config)

	latestBlockNumber, err := client.BlockNumber(context.TODO())
	if err != nil {
		panic(err)
	}
	log.Println("Latest block: ", latestBlockNumber)

	var latestEpochBlockNumber uint64
	if relayer.USE_CHECKPOINT {
		latestEpochBlockNumber = uint64(config.OtherConfig.BlockNumber)
	} else {
		latestEpochBlockNumber = (latestBlockNumber / config.ChainConfig.EpochLength) * config.ChainConfig.EpochLength
	}
	// create newest epoch block for register cross chain relay hub
	// create checkpoint epoch block for register cross chain relay hub
	if latestBlockNumber/config.ChainConfig.EpochLength == latestEpochBlockNumber/config.ChainConfig.EpochLength && int64(latestBlockNumber%config.ChainConfig.EpochLength) < config.ChainConfig.BlockConfirmations {
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
