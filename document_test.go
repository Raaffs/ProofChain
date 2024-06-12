package main_test

import (
	"os"
	"testing"

	// "github.com/Suy56/ProofChain/blockchain"
	"github.com/joho/godotenv"
)

func TestAddDocument(t *testing.T){
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

	if err:=App_test.conn.New("0x13eb090e28f3d062d01e8caa1e4267b8b61cbd59fa1ba76c05bd052bf8bec3f3");err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	_,err=App_test.in.GetDocuments(App_test.conn.CallOpts,App_test.conn.TxOpts.From.Hex())
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

	if err:=App_test.conn.New("0xd605d658a3d05074ed97c7a617c3023a2808d8a492f7d480e418413dd3905d6f");err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	verifierAddr:=App_test.conn.TxOpts.From.Hex()
	err=App_test.in.GetVerifierDocuments(App_test.conn.CallOpts,verifierAddr)
	if err!=nil{
		t.Fatal(err)
	}
}


