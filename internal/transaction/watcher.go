package transaction

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Watcher watches for transactions involving a specific address.
type (
	Watcher struct {
		client        EthereumClient
		handler       mirrorer
		walletAddress common.Address
		logs          chan types.Log
	}

	mirrorer interface {
		MirrorTransaction(ctx context.Context, txHash common.Hash) error
	}
)

// NewWatcher creates a new TransactionWatcher.
func NewWatcher(client EthereumClient, mirrorer mirrorer, walletAddress common.Address) *Watcher {
	return &Watcher{
		client:        client,
		handler:       mirrorer,
		walletAddress: walletAddress,
		logs:          make(chan types.Log),
	}
}

// Watch listens for transactions involving the configured wallet address.
func (tw *Watcher) Watch(ctx context.Context) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tw.walletAddress},
	}

	sub, err := tw.client.SubscribeFilterLogs(ctx, query, tw.logs)
	if err != nil {
		return err
	}

	fmt.Println("Listening for transactions...")

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-tw.logs:
			if err := tw.handler.MirrorTransaction(ctx, vLog.TxHash); err != nil {
				return err
			}
		}
	}
}
