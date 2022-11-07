package utils

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/nguyenkhanhvn/ethereum-merkle-patricia-trie/trie"
)

type ReceiptProofRaw struct {
	Header      []byte
	RawReceipt  []byte
	ProofPath   []byte
	Proof       [][]byte
	ProofRlp    []byte
	ReceiptRoot []byte
}
type ReceiptProofJson struct {
	Header      string
	RawReceipt  string
	ProofPath   string
	Proof       []string
	ProofRlp    string
	ReceiptRoot string
}

type BlockProofsRaw struct {
	Headers    []*types.Header
	HeaderRlps [][]byte
	HeadersRlp []byte
}
type BlockProofsJson struct {
	Headers    []*types.Header
	HeaderRlps []string
	HeadersRlp string
}

type ProofRaw struct {
	BlockProofs   [][]byte
	RawReceipt    []byte
	ProofPath     []byte
	ProofSiblings []byte
}
type ProofJson struct {
	BlockProofs   []string
	RawReceipt    string
	ProofPath     string
	ProofSiblings string
}

func HexToString(data []byte) string {
	return "0x" + hex.EncodeToString(data)
}

func GetBlockNumberFromTransactionHash(client *ethclient.Client, transactionHash common.Hash) (*big.Int, *types.Receipt, error) {
	currentTransaction, isPending, err := client.TransactionByHash(context.TODO(), transactionHash)
	if err != nil {
		return nil, nil, fmt.Errorf("getBlockNumberFromTransactionHash: cannot get client.TransactionByHash, error: %v", err)
	} else if isPending {
		return nil, nil, fmt.Errorf("getBlockNumberFromTransactionHash: transaction is pending")
	}
	receipt, err := client.TransactionReceipt(context.TODO(), currentTransaction.Hash())
	if err != nil {
		return nil, nil, fmt.Errorf("getBlockNumberFromTransactionHash: cannot get client.TransactionReceipt, error: %v", err)
	}
	return receipt.BlockNumber, receipt, nil
}

func GenerateReceiptProof(client *ethclient.Client, transactionHash common.Hash) (*ReceiptProofJson, *ReceiptProofRaw, *types.Block, *types.Receipt, *trie.Trie, error) {
	blockNumber, _, err := GetBlockNumberFromTransactionHash(client, transactionHash)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: cannot getBlockNumberFromTransactionHash, error: %v", err)
	}
	block, err := client.BlockByNumber(context.TODO(), blockNumber)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: cannot get client.BlockByNumber, error: %v", err)
	}

	var rawReceipt, rKey []byte
	var receipt *types.Receipt

	receiptsTrie := trie.NewTrie()
	notFound := true
	for i, tx := range block.Transactions() {
		// key is the encoding of the index as the unsigned integer type
		key, err := rlp.EncodeToBytes(uint(i))
		if err != nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: cannot rlp.EncodeToBytes, error: %v", err)
		}

		transactionReceipt, err := client.TransactionReceipt(context.TODO(), tx.Hash())
		if err != nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: cannot get client.TransactionReceipt, error: %v", err)
		}

		// value is the RLP encoding of a receipt
		receiptRlp, err := rlp.EncodeToBytes(transactionReceipt)
		if err != nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: cannot rlp.EncodeToBytes, error: %v", err)
		}

		receiptsTrie.Put(key, receiptRlp)

		if tx.Hash() == transactionHash {
			rKey = key
			rawReceipt = receiptRlp
			receipt = transactionReceipt
			notFound = false
		}
	}
	if notFound {
		return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: not found TransactionHash: %v" + transactionHash.String())
	}

	proofData, keyNibble, err := receiptsTrie.CreateProof(rKey)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: cannot make proof, error: %v", err)
	}

	headerRlp, err := rlp.EncodeToBytes(block.Header())
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: cannot rlp.EncodeToBytes, error: %v", err)
	}

	proofRlp, err := rlp.EncodeToBytes(proofData)
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("generateReceiptProof: cannot rlp.EncodeToBytes, error: %v", err)
	}

	receiptProofRaw := &ReceiptProofRaw{
		Header:      headerRlp,
		RawReceipt:  rawReceipt,
		ProofPath:   keyNibble,
		Proof:       proofData,
		ProofRlp:    proofRlp,
		ReceiptRoot: block.Header().ReceiptHash.Bytes(),
	}

	var proofString []string
	for _, node := range proofData {
		proofString = append(proofString, HexToString(node))
	}

	receiptProofJson := &ReceiptProofJson{
		Header:      HexToString(headerRlp),
		RawReceipt:  HexToString(rawReceipt),
		ProofPath:   HexToString(keyNibble),
		Proof:       proofString,
		ProofRlp:    HexToString(proofRlp),
		ReceiptRoot: HexToString(block.Header().ReceiptHash.Bytes()),
	}

	return receiptProofJson, receiptProofRaw, block, receipt, receiptsTrie, nil
}

// generate block proofs [fromBlock, toBlock]
func GenerateBlockProofs(client *ethclient.Client, fromBlock, toBlock int64) (*BlockProofsJson, *BlockProofsRaw, error) {
	if toBlock <= fromBlock || fromBlock < 0 {
		return nil, nil, fmt.Errorf("generateBlockProofs: toBlock must be greater than fromBlock")
	}

	var headers []*types.Header
	var headerRlps [][]byte
	var headerRlpsString []string
	for blockNumber := fromBlock; blockNumber <= toBlock; blockNumber++ {
		block, err := client.BlockByNumber(context.TODO(), big.NewInt(blockNumber))
		if err != nil {
			return nil, nil, fmt.Errorf("generateBlockProofs: cannot get client.BlockByNumber, error: %v", err)
		}

		headers = append(headers, block.Header())

		headerRlp, err := rlp.EncodeToBytes(block.Header())
		if err != nil {
			return nil, nil, fmt.Errorf("generateBlockProofs: cannot rlp.EncodeToBytes, error: %v", err)
		}

		headerRlps = append(headerRlps, headerRlp)
		headerRlpsString = append(headerRlpsString, HexToString(headerRlp))
	}
	headersRlp, err := rlp.EncodeToBytes(headerRlps)
	if err != nil {
		return nil, nil, fmt.Errorf("generateBlockProofs: cannot rlp.EncodeToBytes, error: %v", err)
	}

	blockProofsRaw := &BlockProofsRaw{
		Headers:    headers,
		HeaderRlps: headerRlps,
		HeadersRlp: headersRlp,
	}

	blockProofsJson := &BlockProofsJson{
		Headers:    headers,
		HeaderRlps: headerRlpsString,
		HeadersRlp: HexToString(headersRlp),
	}

	return blockProofsJson, blockProofsRaw, nil
}

func GenerateProof(client *ethclient.Client, blockConfirmations int64, transactionHash common.Hash) (*ProofJson, *ProofRaw, *types.Block, *types.Receipt, error) {
	receiptProofJson, receiptProofRaw, block, receipt, _, err := GenerateReceiptProof(client, transactionHash)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("generateProof: cannot generateReceiptProof, error: %v", err)
	}

	lastestNumber, err := client.BlockNumber(context.TODO())
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("generateProof: cannot get client.BlockNumber, error: %v", err)
	}

	blockNumber := receipt.BlockNumber
	if blockNumber.Int64()+blockConfirmations > int64(lastestNumber) {
		return nil, nil, nil, nil, fmt.Errorf("generateProof: must wait until block confirmations, current block is %v, block confirm at block %v", lastestNumber, blockNumber.Int64()+blockConfirmations)
	}

	blockProofsJson, blockProofsRaw, err := GenerateBlockProofs(client, blockNumber.Int64(), blockNumber.Int64()+blockConfirmations)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("generateProof: cannot generateBlockProofs, error: %v", err)
	}

	proofRaw := &ProofRaw{
		BlockProofs:   blockProofsRaw.HeaderRlps,
		RawReceipt:    receiptProofRaw.RawReceipt,
		ProofPath:     receiptProofRaw.ProofPath,
		ProofSiblings: receiptProofRaw.ProofRlp,
	}
	proofJson := &ProofJson{
		BlockProofs:   blockProofsJson.HeaderRlps,
		RawReceipt:    receiptProofJson.RawReceipt,
		ProofPath:     receiptProofJson.ProofPath,
		ProofSiblings: receiptProofJson.ProofRlp,
	}
	return proofJson, proofRaw, block, receipt, nil
}

func GenerateProofWithoutCheck(client *ethclient.Client, fromBlock, toBlock int64, transactionHash common.Hash) (*ProofJson, *ProofRaw, *types.Block, *types.Receipt, error) {
	receiptProofJson, receiptProofRaw, block, receipt, _, err := GenerateReceiptProof(client, transactionHash)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("generateProof: cannot generateReceiptProof, error: %v", err)
	}

	blockProofsJson, blockProofsRaw, err := GenerateBlockProofs(client, fromBlock, toBlock)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("generateProof: cannot generateBlockProofs, error: %v", err)
	}

	proofRaw := &ProofRaw{
		BlockProofs:   blockProofsRaw.HeaderRlps,
		RawReceipt:    receiptProofRaw.RawReceipt,
		ProofPath:     receiptProofRaw.ProofPath,
		ProofSiblings: receiptProofRaw.ProofRlp,
	}
	proofJson := &ProofJson{
		BlockProofs:   blockProofsJson.HeaderRlps,
		RawReceipt:    receiptProofJson.RawReceipt,
		ProofPath:     receiptProofJson.ProofPath,
		ProofSiblings: receiptProofJson.ProofRlp,
	}
	return proofJson, proofRaw, block, receipt, nil
}
