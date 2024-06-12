package main_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Suy56/ProofChain/blockchain"
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
		if err:=App_test.conn.New("0xf8f0694170fdcaaa79fb7e15ee467d2c7e7c50c6144fad676911d230a83239c3");err!=nil{
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

	if err:=App_test.conn.New("0x13eb090e28f3d062d01e8caa1e4267b8b61cbd59fa1ba76c05bd052bf8bec3f3");err!=nil{
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


	err=App_test.in.ApproveVerifier(App_test.conn.TxOpts,"0x7E11d7B81D48c7b54A9D0b3deF1Ab8d12C433A4a")
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
	if err:=App_test.conn.New("0x96d59c83b86521e8a370dd419d09b59de724a71917d2be18f56a4ad7af234e49");err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	err=App_test.in.GetVerifierStatus(App_test.conn.CallOpts,"0x7E11d7B81D48c7b54A9D0b3deF1Ab8d12C433A4a")
	if err!=nil{
		t.Fatal(err)
	}
}

