package main_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/joho/godotenv"
)

type app struct{
	conn	*blockchain.ClientConnection
	in		*blockchain.ContractVerifyOperations
}

var App_test=&app{
	conn: &blockchain.ClientConnection{},
	in: &blockchain.ContractVerifyOperations{},
}

	func TestRegisterUser(t *testing.T){
		err:=godotenv.Load()
		if err!=nil{
			t.Fatal(err)
		}
		contractAddr:=os.Getenv("CONTRACT_ADDR")
		if err:=App_test.conn.New("0x8789157eb94a5dfdc45dc8ae6bf4a39f40bd304a6d0139b81bb79329a2c7cfd9");err!=nil{
			t.Fatal(err)
		}
		App_test.in.Client=App_test.conn.Client
		if err:=App_test.in.New(contractAddr);err!=nil{
			t.Fatal(err)
		}
		err=App_test.in.RegisterUser(App_test.conn.TxOpts,"name","string@gmail.com")
		if err!=nil{
			log.Fatal(err)
		}
	}

func TestRegisterVerifier(t *testing.T){
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")

	if err:=App_test.conn.New("0xb4b1cd087d860bb9c9d286ee2b5fd50ffe2a1ffaefcc0a9b8657eefaa69289d8");err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	err=App_test.in.RegiserVerifier(App_test.conn.TxOpts,"123","kj")
	if err!=nil{
		log.Fatal(err)
	}
	c:=App_test.conn.TxOpts.From
	fmt.Println(c.Hex())


}

func TestApproveVerifier(t *testing.T){
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	privateKey:=os.Getenv("PRIVATE_KEY")
	if err:=App_test.conn.New(privateKey);err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	err=App_test.in.ApproveVerifier(App_test.conn.TxOpts,"kj")
	if err!=nil{
		t.Fatal(err)
	}
}


