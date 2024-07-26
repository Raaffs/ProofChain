package main_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/Suy56/ProofChain/nodeData"
	"github.com/Suy56/ProofChain/wallet"
	"github.com/joho/godotenv"
)

type app struct{
	conn	*blockchain.ClientConnection
	in		*blockchain.ContractVerifyOperations
	keys	*keyUtils.ECKeys
	dataNode 	*nodeData.IPFSManager
}

var App_test=&app{
	conn: &blockchain.ClientConnection{},
	in: &blockchain.ContractVerifyOperations{},
	keys: &keyUtils.ECKeys{},
	dataNode: &nodeData.IPFSManager{},
}

func TestRegisterUser(t *testing.T) {
	username:=""
	privateKey:=""
	if err := godotenv.Load(".env", "accounts/accounts", "keys/keyMap"); err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	errchan := make(chan error, 1)
	publicKeychan := make(chan string, 1)
	contractAddr := os.Getenv("CONTRACT_ADDR")
	if err := blockchain.Init(App_test.conn, App_test.in, privateKey[2:], contractAddr); err != nil {
		t.Fatal("error connecting to contract: ", err)
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		App_test.keys.OnRegister(username, username, publicKeychan, errchan)
	}()
	go func() {
		defer wg.Done()
		wallet.NewWallet(privateKey[2:], username, username, errchan)
	}()
	go func() {
		wg.Wait()
		close(errchan)
		close(publicKeychan)	
	}()
	for{
		select {
		case val, ok := <-publicKeychan:
			if !ok{
				continue
			}else{
				if err:=App_test.in.RegisterUser(App_test.conn.TxOpts,val);err!=nil{
					t.Fatal("Error registering user : ",err)
				}
			}
		case err,ok:= <-errchan:
			if !ok{
				return
			}
			if err!=nil{
				fmt.Println("Error : ",err)
				t.Fatal(err)
			}

		}
	}
}

func TestLoginUser(t *testing.T){
	var wg sync.WaitGroup
	errchan:=make(chan error)
	if err:=godotenv.Load(".env","accounts/accounts","keys/keyMap");err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	wg.Add(1)
	go func() {
		defer wg.Done()
		App_test.keys.OnLogin("Maria","Maria",errchan)
	}()
	go func(){
		wg.Wait()
		defer close(errchan)
	}()

	for err:=range errchan{
		if err!=nil{
			t.Fatal("error logging in : ",err)
		}
	}
	pk,err:=wallet.RetriveAccount("",""); if err!=nil{
		t.Fatal(err)
	}
	if err:=blockchain.Init(App_test.conn,App_test.in,pk,contractAddr);err!=nil{
		t.Fatal(err)
	}
}


func TestRegisterVerifier(t *testing.T){
	institute:=""
	privateKey:=""
	if err := godotenv.Load(".env", "accounts/accounts", "keys/keyMap"); err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	errchan := make(chan error, 1)
	publicKeychan := make(chan string, 1)
	contractAddr := os.Getenv("CONTRACT_ADDR")
	if err := blockchain.Init(App_test.conn, App_test.in, privateKey[2:], contractAddr); err != nil {
		t.Fatal("error connecting to contract: ", err)
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		App_test.keys.OnRegister(institute, institute, publicKeychan, errchan)
	}()
	go func() {
		defer wg.Done()
		wallet.NewWallet(privateKey[2:], institute, institute, errchan)
	}()
	go func ()  {
		wg.Wait()
		close(errchan)
		close(publicKeychan)	
	}()
	for{
		select {
		case val, ok := <-publicKeychan:
			if !ok {
				continue
			} else {
				fmt.Println(val)
				if err := App_test.in.RegisterInstitution(App_test.conn.TxOpts, val, ""); err != nil {
					t.Fatal(err)
				}
			}
		case err,ok:= <-errchan:
			if !ok{
				return
			}
			if err!=nil{
				t.Fatal(err)
			}
		}
	}

}

func TestApproveVerifier(t *testing.T){
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	privateKey:=os.Getenv("PRIVATE_KEY")
	fmt.Println(privateKey,contractAddr)
	if err:=blockchain.Init(App_test.conn,App_test.in,privateKey[2:],contractAddr);err!=nil{
		t.Fatal(err)
	}
	if err:=App_test.in.ApproveVerifier(App_test.conn.TxOpts,"kj"); err!=nil{
		t.Fatal(err)
	}
}

func TestGetInstitute(t *testing.T) {
	err:=godotenv.Load()
	if err!=nil{
		t.Fatal(err)
	}
	contractAddr:=os.Getenv("CONTRACT_ADDR")
	privateKey:=os.Getenv("PRIVATE_KEY")

	if err:=blockchain.Init(App_test.conn,App_test.in,privateKey[2:],contractAddr);err!=nil{
		t.Fatal(err)
	}

	key,_:=App_test.in.Instance.GetInstituePublicKey(App_test.conn.CallOpts,"kj")
	fmt.Println(key)
}