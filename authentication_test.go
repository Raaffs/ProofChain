package main_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common"
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
		if err:=App_test.conn.New("0x5906abc1b858a4dd0da8d309cb9c43c9fa2652e73832fc6805a50754c42e93fd");err!=nil{
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

	if err:=App_test.conn.New("0x7aba46cec1861fcbe98665858d2030b83dc7810d7173c313c0b93071d861a050");err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	err=App_test.in.RegiserVerifier(App_test.conn.TxOpts,"name","string@gmail.com","fd","kj")
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

	pubAddrHex:="0xE582078b233fDA011eab94bc9833dBA72A6fd21D"
	pubAddr:=common.HexToAddress(pubAddrHex)
	err=App_test.in.ApproveVerifier(App_test.conn.TxOpts,pubAddr)
	if err!=nil{
		t.Fatal(err)
	}
}

func TestCheckVerifierStatus(t *testing.T){
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	if err:=App_test.conn.New("0x7aba46cec1861fcbe98665858d2030b83dc7810d7173c313c0b93071d861a050");err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	err=App_test.in.GetVerifierStatus(App_test.conn.CallOpts)
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(App_test.conn.CallOpts)
}

