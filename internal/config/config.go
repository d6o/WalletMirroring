package config

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Config holds configuration values.
type Config struct {
	NodeURL       string
	WalletAddress common.Address
	PrivateKey    *ecdsa.PrivateKey
}

func Load() (*Config, error) {
	nodeURL := os.Getenv("ETH_NODE_URL")
	if nodeURL == "" {
		return nil, errors.New("ETH_NODE_URL environment variable not set")
	}

	walletAddress := os.Getenv("WALLET_ADDRESS")
	if walletAddress == "" {
		return nil, errors.New("WALLET_ADDRESS environment variable not set")
	}

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		return nil, errors.New("PRIVATE_KEY environment variable not set")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %v", err)
	}

	return &Config{
		NodeURL:       nodeURL,
		WalletAddress: common.HexToAddress(walletAddress),
		PrivateKey:    privateKey,
	}, nil
}
