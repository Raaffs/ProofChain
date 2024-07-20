package main_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func TestAddDocument(t *testing.T){
	err:=godotenv.Load();if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	errchan:=make(chan error)
	var wg sync.WaitGroup

	if err:=blockchain.Init(App_test.conn,App_test.in,"",contractAddr);err!=nil{
		t.Fatal(err)
	}
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		App_test.keys.OnLogin("","",errchan)		
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()
	App_test.dataNode.New("5001")
	cid,err:=App_test.dataNode.Upload("testFiles/test1.txt");if err!=nil{
		t.Fatal("Error uploading file on ipfs : ",err)
	}
	pub,err:=App_test.in.Instance.GetInstituePublicKey(App_test.conn.CallOpts,"");if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("institute public key : ",pub)
	if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
		t.Fatal(err)
	}
	sharedSecret,err:=App_test.keys.GenerateSecret();if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(cid)
	encryptedText,err:=keyUtils.EncryptIPFSHash(sharedSecret,[]byte(cid));if err!=nil{
		t.Fatal(err)
	}
	if err:=App_test.in.AddDocument(App_test.conn.TxOpts,encryptedText,"","","");err!=nil{
		t.Fatal(err)
	}
	fmt.Println("lenght : ",len(encryptedText))
	
	_,err=keyUtils.DecryptIPFSHash(sharedSecret,[]byte(encryptedText));if err!=nil{
		t.Fatal(err)
	}
}

func TestDoc(t *testing.T){
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}	
	contractAddr:=os.Getenv("CONTRACT_ADDR")

	if err:=blockchain.Init(App_test.conn,App_test.in,"",contractAddr);err!=nil{
		t.Fatal(err)
	}
	docs,err:=App_test.in.GetDocuments(App_test.conn.CallOpts);if err!=nil{
		t.Fatal(err)
	}

	errchan:=make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		App_test.keys.OnLogin("","",errchan)
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()
	for _,doc :=range docs{
		pub,err:=App_test.in.Instance.GetUserPublicKey(App_test.conn.CallOpts,common.HexToAddress(doc.Requester));if err!=nil{
			fmt.Println("error : ",err)
			continue
		}
		fmt.Println("user pub : ",pub)
		if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
			fmt.Println("error : ",err)
			continue
		}
		sec,err:=App_test.keys.GenerateSecret();if err!=nil{
			fmt.Println("error : ",err)
			continue
		}
		ipfs,err:=keyUtils.DecryptIPFSHash(sec,[]byte(doc.IpfsAddress));if err!=nil{
			fmt.Println("error : ",err,ipfs)
			continue
		}
		fmt.Println("ipfs: ",ipfs)
	}
}

func TestApproveDoc(t *testing.T) {
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")

	if err:=App_test.conn.New("");err!=nil{
		t.Fatal(err)
	}
	App_test.in.Client=App_test.conn.Client
	if err:=App_test.in.New(contractAddr);err!=nil{
		t.Fatal(err)
	}
	if err!=nil{t.Fatal(err)}	
	App_test.in.VerifyDocument(App_test.conn.TxOpts,"","",0)
}		



