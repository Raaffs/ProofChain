package main_test

import (
	"os"
	"testing"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/joho/godotenv"
)
func TestDeploy(t *testing.T){
	app:=&struct{
		conn *blockchain.ClientConnection
		in *blockchain.ContractVerifyOperations
	}{
		conn: &blockchain.ClientConnection{},
		in: &blockchain.ContractVerifyOperations{},
	}
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal("Error loading env")
	}
	privateKey:=os.Getenv("PRIVATE_KEY")
	blockchain.Init(app.conn,app.in,privateKey[2:],"")
	if err!=nil{
		t.Fatal("Error connecting to chain")
	}
	_,_,err=blockchain.Deploy(app.conn.TxOpts,app.conn.Client);if err!=nil{
		t.Fatal(err)
	}
}