package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
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
		log.Fatal("Error getting chainID : ",err)
	}
	conn.ChainId=chainID

	err=conn.setTxOpts(privateKey);if err!=nil{
		return err
	}


	return nil
}

func (conn *ClientConnection) setTxOpts(privateKeyString string) error {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !ok {
		return err
	}
	nonce, err := conn.Client.PendingNonceAt(conn.ctx, fromAddress)
	if err != nil {
		log.Fatal("nonce error : ", err)
	}

	fmt.Println("nonce : ", nonce)

	gasPrice, err := conn.Client.SuggestGasPrice(conn.ctx)
	if err != nil {
		log.Fatal("Error getting gas price : ", err)
	}
	chainID, err := conn.Client.ChainID(conn.ctx)
	if err != nil {
		log.Fatal("Error getting chainID : ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Error new key transaction : ", err)
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

			