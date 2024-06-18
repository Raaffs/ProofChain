package main_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAddDocument(t *testing.T){
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

	// ipfsHash,err:=blockchain.IpfsHashTo32Byte("QmTzQ1hZ6sFS7yzG5fbEv68nW7uyDTP5JDyz1iLCv2gZjo")
	if err!=nil{t.Fatal(err)}
	if err:=App_test.in.AddDocument(App_test.conn.TxOpts,"0x111","kj","aadhar","desc");err!=nil{
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

	if err:=App_test.conn.New("0xb4b1cd087d860bb9c9d286ee2b5fd50ffe2a1ffaefcc0a9b8657eefaa69289d8");err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	// pubAddrHex:="0xb4b1cd087d860bb9c9d286ee2b5fd50ffe2a1ffaefcc0a9b8657eefaa69289d8"
	// pubAddr:=common.HexToAddress(pubAddrHex)
	if err!=nil{t.Fatal(err)}	
	App_test.in.VerifyDocument(App_test.conn.TxOpts,"kj","0x111",0)
}		



