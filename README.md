# WalletMirroring

WalletMirroring is a Go-based application that monitors a specified Ethereum wallet for transactions and sends identical
transactions from a different wallet. This tool can be used for purposes such as mirroring transaction activity or
creating backups of transaction data.

## Features

- Monitors an Ethereum wallet for new transactions.
- Sends identical transactions using a different wallet.
- Configurable via environment variables.
- Robust error handling and retry logic.
- Modular design with clear separation of concerns.

## Requirements

- Go 1.16 or later
- An Ethereum node URL (e.g., Infura project URL)
- A private key for the wallet that will send the mirrored transactions

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/WalletMirroring.git
cd WalletMirroring
```

2. Export the following variables:

```ini
export ETH_NODE_URL=https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID
export WALLET_ADDRESS=0xWalletAddress
export PRIVATE_KEY=YourPrivateKey
```

Replace `YOUR_INFURA_PROJECT_ID`, `0xWalletAddress`, and `YourPrivateKey` with your actual values.

## Usage

1. Build the project:

```sh
go build -o walletmirroring
```

2. Run the executable:

```sh
./walletmirroring
```

The application will start monitoring the specified wallet for new transactions and send identical transactions from
the wallet associated with the provided private key.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License.
