package blockchain

import (
	"context"
	"crypto/ecdsa"
	"log"
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

func (conn *ClientConnection)New(privateKey string) error {
	conn.ctx=context.Background()
	chainID,err:=conn.Client.ChainID(conn.ctx);if err!=nil{
		log.Println("Error  getting chain ID:",err)
		return err
	}
	
	conn.ChainId=chainID

	if err:=conn.setTxOpts(privateKey); err!=nil{
		log.Println("Error setting txOpts:",err)
		return err
	}

	return nil
}

func (conn *ClientConnection) setTxOpts(privateKeyString string) error {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Println("Error parsing private key:", err)
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !ok {
		log.Println("Error casting public key to ECDSA")
		return err
	}
	nonce, err := conn.Client.PendingNonceAt(conn.ctx, fromAddress)
	if err != nil {
		log.Println("Error getting nonce:", err)
		return err 
	}

	gasPrice, err := conn.Client.SuggestGasPrice(conn.ctx)
	if err != nil {
		log.Println("Error getting gas price:", err)
		return err
	}
	chainID, err := conn.Client.ChainID(conn.ctx)
	if err != nil {
		log.Println("Error getting chain ID:", err)
		return err 
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Println("Error creating auth:", err)
		return err
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

			