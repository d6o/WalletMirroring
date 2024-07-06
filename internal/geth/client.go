package geth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthereumClient interface abstracts the Ethereum client methods.
type EthereumClient interface {
	TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
}

// Client wraps ethclient.Client to implement EthereumClient interface.
type Client struct {
	*ethclient.Client
}

// NewClient creates a new GethClient.
func NewClient(url string) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}
