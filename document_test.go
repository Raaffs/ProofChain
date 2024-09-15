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

	if err:=blockchain.Init(App_test.conn,App_test.in,"5ae503b8f45687a0b1e13d7e2f41d1b830c864b47ecd486b86366eb389fc6a74",contractAddr);err!=nil{
		t.Fatal(err)
	}
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		App_test.keys.OnLogin("user","user",errchan)		
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()
	if err:=App_test.in.AddDocument(App_test.conn.TxOpts,"hi",("hi2"),"ins","name");err!=nil{
		t.Fatal(err)
	}
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


//errors might be related to how address are stored in accounts. Letters are not capitialized 
//in when wallets are made
func TestApproveDoc(t *testing.T) {
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	errchan:=make(chan error)
	var wg sync.WaitGroup

	err=blockchain.Init(App_test.conn,App_test.in,"f400f71d5132397be7708ef07285c4e9fe5a1b526e5034febf391ee6c4af28e0",contractAddr);if err!=nil{t.Fatal(err)}	
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		App_test.keys.OnLogin("ins","ins",errchan)		
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()

	// pub,err:=App_test.in.Instance.GetUserPublicKey(App_test.conn.CallOpts,common.HexToAddress("0x9F8cb1C4f6329CFB72C7034C010eD78bdc8e3976"));if err!=nil{
	// 	t.Fatal("error getting user's pub key: ",err)
	// }
	// if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
	// 	t.Fatal("error getting multi sig:",err)
	// }
	// sec,err:=App_test.keys.GenerateSecret();if err!=nil{
	// 	t.Fatal("err sec: ",err)
	// }
	// enc,err:=keyUtils.EncryptIPFSHash(sec,[]byte("QmX8asJ26eMXXCiuKxk7N8QFBKzEwYHmNp2Ss8GNwQYFKn"));if err!=nil{
	// 	t.Fatal(err)
	// }
	err=App_test.in.VerifyDocument(App_test.conn.TxOpts,"8cceabce917f7f34e4a6dddb2eb320b30041902b9417e1eab5569992e2196b29","ins",2);if err!=nil{
		t.Fatal("Error approving document : ",err)
	}		
	index,err:=App_test.in.Instance.GetDocIndexCounter(App_test.conn.CallOpts);if err!=nil{
		t.Fatal("error docsCounterIndex",err)
	}
	fmt.Println("docindexCounter : ",index)
	index2,err:=App_test.in.Instance.GetDocumentIndex(App_test.conn.CallOpts,("d3d095dc3fd845d259516e8fc78b4b56db2b89700a319a73817c1cf2e28115a3"));if err!=nil{
		t.Fatal("error documentLists[index]",err)
	}
	fmt.Println("documentLists[index]",index2)
	// stat,err:=App_test.in.Instance.VerifyDocumentTest2(App_test.conn.TxOpts,"ins",enc,0); if err!=nil{
	// 	t.Fatal("Error stat 2",err)
	// }
	// fmt.Println("stat2 ",stat.Value())
}		

func TestDocIndex(t *testing.T) {
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
		App_test.keys.OnLogin("ins","ins",errchan)		
	}()
	go func() {
		wg.Wait()
		close(errchan)
	}()	
	err=blockchain.Init(App_test.conn,App_test.in,"fba76c12f61bc57fa86348d2228676dd0cf207baa314541b67a2b34105b8f5d5",contractAddr);if err!=nil{
		t.Fatal(err)
	}	
	// pub,err:=App_test.in.Instance.GetUserPublicKey(App_test.conn.CallOpts,common.HexToAddress("0x9F8cb1C4f6329CFB72C7034C010eD78bdc8e3976"));if err!=nil{
	// 	t.Fatal("error getting user's pub key: ",err)
	// }
	// if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
	// 	t.Fatal("error getting multi sig:",err)
	// }
	// sec,err:=App_test.keys.GenerateSecret();if err!=nil{
	// 	t.Fatal("err sec: ",err)
	// }
	// enc,err:=keyUtils.EncryptIPFSHash(sec,[]byte("QmX8asJ26eMXXCiuKxk7N8QFBKzEwYHmNp2Ss8GNwQYFKn"));if err!=nil{
	// 	t.Fatal(err)
	// }



	index,err:=App_test.in.Instance.GetDocIndexCounter(App_test.conn.CallOpts);if err!=nil{
		t.Fatal("error docsCounterIndex",err)
	}
	fmt.Println("docindexCounter : ",index)
	index2,err:=App_test.in.Instance.GetDocumentIndex(App_test.conn.CallOpts,"hi2");if err!=nil{
		t.Fatal("error documentLists[index]",err)
	}
	fmt.Println(" documentLists[index]",index2)

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

	err=blockchain.Init(App_test.conn,App_test.in,"84ef8fd680bf64990853c1a044090905cda74903b65b7fe668f281582bb2b434",contractAddr);if err!=nil{t.Fatal(err)}	
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
		fmt.Println(doc.Stats,doc.ShaHash)
	}

}