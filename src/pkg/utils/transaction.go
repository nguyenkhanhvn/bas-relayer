package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTransactOpts(client *ethclient.Client, privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, Error) {
	if client == nil {
		return nil, NewError(ClientError, fmt.Errorf("GetTransactOpts: client is nil"))
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, NewError(OtherError, fmt.Errorf("GetTransactOpts: error casting public key to ECDSA"))
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.TODO(), fromAddress)
	if err != nil {
		return nil, NewError(ClientError, err)
	}
	gasPrice, err := client.SuggestGasPrice(context.TODO())
	if err != nil {
		log.Println(err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, NewError(OtherError, fmt.Errorf("GetTransactOpts: %v", err))
	}
	opts.Nonce = big.NewInt(int64(nonce))
	// opts.Value = big.NewInt(0) // in wei
	// opts.GasLimit = uint64(30000000) // in units
	opts.GasPrice = gasPrice

	return opts, nil
}

func WaitTxToBeMined(ctx context.Context, client *ethclient.Client, txHash common.Hash) Error {
	if client == nil {
		return NewError(ClientError, fmt.Errorf("GetTransactOpts: client is nil"))
	}

	tries := 30
	for tries > 0 {
		receipt, _ := client.TransactionReceipt(ctx, txHash)
		if receipt != nil {
			return nil
		}
		tries--
		time.Sleep(time.Second)
	}
	return NewError(OtherError, fmt.Errorf("transaction (%s) can't be mined in time", txHash.Hex()))
}
