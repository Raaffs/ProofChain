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

	err=app.conn.New(privateKey)
	if err!=nil{
		t.Fatal("Error connecting to chain")
	}
	blockchain.Deploy(app.conn.TxOpts,app.conn.Client)
}