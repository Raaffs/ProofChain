package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ClientConnection struct {
	Client    *ethclient.Client
	ClientURL string
	TxOpts    *bind.TransactOpts
	CallOpts  *bind.CallOpts
	ChainId   *big.Int
	ctx       context.Context
}

func (conn *ClientConnection)SetClient(c *ethclient.Client) {
	conn.Client=c
}

func (conn *ClientConnection) New(privateKey string) error {
    conn.ctx = context.Background()
    chainID, err := conn.Client.ChainID(conn.ctx)
    if err != nil {
        return fmt.Errorf("Error getting chain ID: %v", err)
    }
    
    conn.ChainId = chainID

    if err := conn.setTxOpts(privateKey); err != nil {
        return fmt.Errorf("Error setting txOpts: %v", err)
    }

    return nil
}

func (conn *ClientConnection) setTxOpts(privateKeyString string) error {
    privateKey, err := crypto.HexToECDSA(privateKeyString)
    if err != nil {
        return fmt.Errorf("Error parsing private key: %v", err)
    }
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    if !ok {
        return fmt.Errorf("Error casting public key to ECDSA")
    }
    nonce, err := conn.Client.PendingNonceAt(conn.ctx, fromAddress)
    if err != nil {
        return fmt.Errorf("Error getting nonce: %v", err)
    }

    gasPrice, err := conn.Client.SuggestGasPrice(conn.ctx)
    if err != nil {
        return fmt.Errorf("Error getting gas price: %v", err)
    }
    chainID, err := conn.Client.ChainID(conn.ctx)
    if err != nil {
        return fmt.Errorf("Error getting chain ID: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        return fmt.Errorf("Error creating auth: %v", err)
    }
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(3000000)
    auth.GasPrice = gasPrice
    conn.TxOpts = auth
    conn.setCallOpts(fromAddress)
    return nil
}

func (conn *ClientConnection)setCallOpts(fromAddress common.Address)  {
	conn.CallOpts=&bind.CallOpts{
		From: fromAddress,
	}
}

			