package main

import (
	"log"
	"os"
	"testing"

	"github.com/Suy56/ProofChain/chaincore/core"
	"github.com/joho/godotenv"
)
var ETH_CLIENT_URL="http://localhost:7545"

func TestDeploy(t *testing.T){
	if err:=godotenv.Load(".env");err!=nil{
		t.Fatal("Error reading .env: ",err)
	}
	app:=&struct{
		conn *blockchain.ClientConnection
		in *blockchain.ContractVerifyOperations
	}{
		conn: &blockchain.ClientConnection{},
		in: &blockchain.ContractVerifyOperations{},
	}

	privateKey:=os.Getenv("PRIVATE_KEY")
	if err:=blockchain.Init(app.conn,app.in,privateKey[2:],"",ETH_CLIENT_URL);err!=nil{
		t.Fatal(err)
	}
	contract,_,err:=blockchain.Deploy(app.conn.TxOpts,app.conn.Client);if err!=nil{
		log.Println("contract:",contract)
		t.Fatal(err)
	}
	log.Println(contract)
}