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

	if err:=blockchain.Init(App_test.conn,App_test.in,"9d99735c0d9902c6d4baba9c66f611a45705c52e65f6b037c6fb2e06293273bb",contractAddr);err!=nil{
		t.Fatal(err)
	}
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		App_test.keys.OnLogin("Maria","Maria",errchan)		
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()
	App_test.dataNode.New("5001")
	cid,err:=App_test.dataNode.Upload("testFiles/test1.txt");if err!=nil{
		t.Fatal("Error uploading file on ipfs : ",err)
	}
	pub,err:=App_test.in.Instance.GetInstituePublicKey(App_test.conn.CallOpts,"ins");if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("institute public key : ",pub)
	if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
		t.Fatal(err)
	}
	fmt.Println(App_test.keys.MultiSig)

	sharedSecret,err:=App_test.keys.GenerateSecret();if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("cid ",cid)
	fmt.Println("shared secret",sharedSecret)
	encryptedText,err:=keyUtils.EncryptIPFSHash(sharedSecret,[]byte(cid));if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("encrypted text",encryptedText)
	// if err:=App_test.in.AddDocument(App_test.conn.TxOpts,encryptedText,"ins","name","desc");err!=nil{
	// 	t.Fatal(err)
	// }
	fmt.Println("lenght : ",len(encryptedText))
	
	dec,err:=keyUtils.DecryptIPFSHash(sharedSecret,[]byte(encryptedText));if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("decrypted text",dec)
}


func TestDoc(t *testing.T){
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}	
	contractAddr:=os.Getenv("CONTRACT_ADDR")

	if err:=blockchain.Init(App_test.conn,App_test.in,"3a688747a7ea218d2db2d17c5ae2d7b04f9e6fd57284a1616b07c4bd87821ec2",contractAddr);err!=nil{
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
		App_test.keys.OnLogin("ins","ins",errchan)
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
		fmt.Println(ipfs)
	}

}

func TestApproveDoc(t *testing.T) {
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	errchan:=make(chan error)
	var wg sync.WaitGroup

	err=blockchain.Init(App_test.conn,App_test.in,"3a688747a7ea218d2db2d17c5ae2d7b04f9e6fd57284a1616b07c4bd87821ec2",contractAddr);if err!=nil{t.Fatal(err)}	
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		App_test.keys.OnLogin("ins","ins",errchan)		
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()

	pub,err:=App_test.in.Instance.GetUserPublicKey(App_test.conn.CallOpts,common.HexToAddress("0x442783011f4A7e04eb00CEa8Ae508b65E7A0283C"));if err!=nil{
		t.Fatal("error getting user's pub key: ",err)
	}

	if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
		t.Fatal("error getting multi sig:",err)
	}
	sec,err:=App_test.keys.GenerateSecret();if err!=nil{
		t.Fatal("err sec: ",err)
	}
	enc,err:=keyUtils.EncryptIPFSHash(sec,[]byte("QmQVH5s5nk8b4UoC7cNZS6xB1cFMs4jWQJiUbjQvnq8KfF"));if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("enc",enc)
	if err:=App_test.in.VerifyDocument(App_test.conn.TxOpts,"ins",enc,0);err!=nil{
		t.Fatal("Error approving document : ",err)
	}

}		

func TestEncryptionEquality(t *testing.T) {
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	errchan:=make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		App_test.keys.OnLogin("Maria","Maria",errchan)		
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()	

	err=blockchain.Init(App_test.conn,App_test.in,"9d99735c0d9902c6d4baba9c66f611a45705c52e65f6b037c6fb2e06293273bb",contractAddr);if err!=nil{t.Fatal(err)}	
	docs,err:=App_test.in.GetDocuments(App_test.conn.CallOpts);if err!=nil{
		t.Fatal(err)
	}
	doc:=docs[0]
	pub,err:=App_test.in.Instance.GetInstituePublicKey(App_test.conn.CallOpts,doc.Institute); if err!=nil{
		t.Fatal("error getting institute's pub key: ",err)
	}
	fmt.Println("institute public key : ",doc.Institute,pub)
	if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
		t.Fatal("error setting multi sig key: ",err)
	}
	fmt.Println(App_test.keys.MultiSig)
	sec,err:=App_test.keys.GenerateSecret();if err!=nil{
		t.Fatal("error generating sec:",err)
	}
	fmt.Println("sec: ",sec)
	dec,err:=keyUtils.DecryptIPFSHash(sec,[]byte(doc.IpfsAddress));if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("dec",dec)
	// err=blockchain.Init(App_test.conn,App_test.in,"607a72b2808feb3cd11f582cc1253c373da1e95c8af26322419763b852e1cb94",contractAddr);if err!=nil{t.Fatal(err)}	
	// pub,err=App_test.in.Instance.GetInstituePublicKey(App_test.conn.CallOpts,"ins"); if err!=nil{
	// 	t.Fatal(err)
	// }
	// if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
	// 	t.Fatal(err)
	// }
	// sec,err=App_test.keys.GenerateSecret();if err!=nil{
	// 	t.Fatal(err)
	// }
	// enc2,err:=keyUtils.DecryptIPFSHash(sec,[]byte("QmTQpEySom7HNb4HJrZPEqTQnC7prd9Z9eiWLrj4moYVD3"));if err!=nil{
	// 	t.Fatal(err)
	// }
	// fmt.Println("enc : ",enc2==enc)
}

func TestGetUserDocs(t *testing.T){
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	errchan:=make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		App_test.keys.OnLogin("Maria","Maria",errchan)		
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()	

	err=blockchain.Init(App_test.conn,App_test.in,"9d99735c0d9902c6d4baba9c66f611a45705c52e65f6b037c6fb2e06293273bb",contractAddr);if err!=nil{t.Fatal(err)}	
	docs,err:=App_test.in.GetDocuments(App_test.conn.CallOpts);if err!=nil{
		t.Fatal("Error getting docs:",err)
	}
	for _,doc :=range docs{
		pub,err:=App_test.in.Instance.GetInstituePublicKey(App_test.conn.CallOpts,doc.Institute);if err!=nil{
			fmt.Println("error : ",err)
			continue
		}
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
		fmt.Println(ipfs)
	}

}