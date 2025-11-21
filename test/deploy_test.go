package main

import (
	"fmt"
	"testing"

	"github.com/Suy56/ProofChain/chaincore/core"
)
var ETH_CLIENT_URL="http://localhost:7545"

func TestDeploy(t *testing.T){
	app:=&struct{
		conn *blockchain.ClientConnection
		in *blockchain.ContractVerifyOperations
	}{
		conn: &blockchain.ClientConnection{},
		in: &blockchain.ContractVerifyOperations{},
	}

	privateKey:="0x552d079fe5a76c34f075ddd694784a0f9a3c19aeadc62632fde406492073b48d"
	fmt.Println("p: ",privateKey)
	if err:=blockchain.Init(app.conn,app.in,privateKey[2:],"",ETH_CLIENT_URL);err!=nil{
		t.Fatal(err)
	}
	_,_,err:=blockchain.Deploy(app.conn.TxOpts,app.conn.Client);if err!=nil{
		t.Fatal(err)
	}
}