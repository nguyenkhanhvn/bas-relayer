package utils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTransactOpts(client *ethclient.Client, privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	opts.Nonce = big.NewInt(int64(nonce))
	// opts.Value = big.NewInt(0) // in wei
	// opts.GasLimit = uint64(30000000) // in units
	opts.GasPrice = gasPrice

	return opts, nil
}

func WaitTxToBeMined(ctx context.Context, eth *ethclient.Client, txHash common.Hash) error {
	tries := 30
	for tries > 0 {
		receipt, _ := eth.TransactionReceipt(ctx, txHash)
		if receipt != nil {
			return nil
		}
		tries--
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("transaction (%s) can't be mined in time", txHash.Hex())
}
