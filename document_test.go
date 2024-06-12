package main_test

import (
	"fmt"
	"os"
	"testing"

	// "github.com/Suy56/ProofChain/blockchain"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/joho/godotenv"
)

func TestAddDocument(t *testing.T){
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

	// ipfsHash,err:=blockchain.IpfsHashTo32Byte("QmTzQ1hZ6sFS7yzG5fbEv68nW7uyDTP5JDyz1iLCv2gZjo")
	if err!=nil{t.Fatal(err)}
	if err:=App_test.in.AddDocument(App_test.conn.TxOpts,"doc3","0x111","des");err!=nil{
		t.Fatal(err)
	}

}

func TestDoc(t *testing.T){
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
	_,err=App_test.in.GetDocuments(App_test.conn.CallOpts)
	if err!=nil{
		t.Fatal(err)
	}
}

func TestApproveDoc(t *testing.T) {
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
	// ipfsHash,err:=blockchain.IpfsHashTo32Byte("QmTzQ1hZ6sFS7yzG5fbEv68nW7uyDTP5JDyz1iLCv2gZjo")
	if err!=nil{t.Fatal(err)}	
	App_test.in.VerifyDocument(App_test.conn.TxOpts,"0x111",1)
}		


func TestGetVerifierDocuments(t *testing.T){
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
	verifierAddr:=App_test.conn.TxOpts.From.Hex()
	fmt.Println(verifierAddr)
	err=App_test.in.GetVerifierDocuments(App_test.conn.CallOpts)
	if err!=nil{
		t.Fatal(err)
	}
}


