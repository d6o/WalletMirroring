package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Mirrorer handles the logic for processing transactions.
type Mirrorer struct {
	client     EthereumClient
	privateKey *ecdsa.PrivateKey
}

// NewMirrorer creates a new TransactionHandler.
func NewMirrorer(client EthereumClient, privateKey *ecdsa.PrivateKey) *Mirrorer {
	return &Mirrorer{
		client:     client,
		privateKey: privateKey,
	}
}

func (th *Mirrorer) MirrorTransaction(ctx context.Context, txHash common.Hash) error {
	tx, isPending, err := th.client.TransactionByHash(ctx, txHash)
	if err != nil {
		return err
	}

	if isPending {
		return nil
	}

	return th.sendIdenticalTransaction(ctx, tx)
}

func (th *Mirrorer) sendIdenticalTransaction(ctx context.Context, tx *types.Transaction) error {
	fromAddress := crypto.PubkeyToAddress(th.privateKey.PublicKey)
	nonce, err := th.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := th.client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(th.privateKey, tx.ChainId())
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = tx.Value()
	auth.GasLimit = tx.Gas()
	auth.GasPrice = gasPrice
	auth.Context = ctx

	rawTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      tx.Gas(),
		To:       tx.To(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	})

	signedTx, err := auth.Signer(fromAddress, rawTx)
	if err != nil {
		return err
	}

	retryCount := 0
	for {
		err = th.client.SendTransaction(ctx, signedTx)
		if err == nil {
			break
		}
		if retryCount >= 3 {
			return err
		}

		retryCount++
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("Sent identical transaction: %s\n", signedTx.Hash().Hex())
	return nil
}
