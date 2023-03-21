package utils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockNumberFromTransactionHash(client *ethclient.Client, transactionHash common.Hash) (*big.Int, *types.Receipt, Error) {
	if client == nil {
		return nil, nil, NewError(ClientError, fmt.Errorf("getBlockNumberFromTransactionHash: client is nil"))
	}
	currentTransaction, isPending, err := client.TransactionByHash(context.TODO(), transactionHash)
	if err != nil {
		return nil, nil, NewError(ClientError, fmt.Errorf("getBlockNumberFromTransactionHash: cannot get client.TransactionByHash, error: %v", err))
	} else if isPending {
		return nil, nil, NewError(OtherError, fmt.Errorf("getBlockNumberFromTransactionHash: transaction is pending"))
	}
	receipt, err := client.TransactionReceipt(context.TODO(), currentTransaction.Hash())
	if err != nil {
		return nil, nil, NewError(ClientError, fmt.Errorf("getBlockNumberFromTransactionHash: cannot get client.TransactionReceipt, error: %v", err))
	}
	return receipt.BlockNumber, receipt, nil
}

func GetValidatorSetFromParliaBlock(client *ethclient.Client, blockNumber int64) ([]common.Address, Error) {
	if client == nil {
		return nil, NewError(ClientError, fmt.Errorf("GetValidatorSetFromParliaBlock: client is nil"))
	}
	block, err := client.BlockByNumber(context.TODO(), big.NewInt(blockNumber))
	if err != nil {
		return nil, NewError(ClientError, fmt.Errorf("GetValidatorSetFromParliaBlock: cannot get client.BlockByNumber, error: %v", err))
	}

	validatorSet, err := ParseValidatorSet(block.Header())
	if err != nil {
		return nil, NewError(OtherError, fmt.Errorf("GetValidatorSetFromParliaBlock: %v", err))
	}
	return validatorSet, nil
}
