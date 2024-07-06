package main

import (
	"context"
	"log"

	"github.com/d6o/WalletMirroring/internal/config"
	"github.com/d6o/WalletMirroring/internal/geth"
	"github.com/d6o/WalletMirroring/internal/transaction"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	client, err := geth.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	transactionMirrorer := transaction.NewMirrorer(client, cfg.PrivateKey)

	watcher := transaction.NewWatcher(client, transactionMirrorer, cfg.WalletAddress)
	if err := watcher.Watch(ctx); err != nil {
		log.Fatalf("Error mirroring wallet: %v", err)
	}
}
