package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/Suy56/ProofChain/nodeData"
	"github.com/Suy56/ProofChain/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

// App struct
type App struct {
	ctx      		context.Context
	conn		  	*blockchain.ClientConnection
	instance 		*blockchain.ContractVerifyOperations
	keys			*keyUtils.ECKeys
	ipfs 			*nodeData.IPFSManager
	user			string
	isApproved		bool
	envMap			map[string]any
}
// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx=ctx
	keyMap,err:=godotenv.Read(".env");if err!=nil{
		log.Println("Error loading .env : " ,err)
		return
	}
	for key,val:=range keyMap{
		app.envMap[key]=val
	}
	app.ipfs.New("5001")	
}

func (app *App)Login(username string, password string) (error) {
	var wg sync.WaitGroup
	errchan:=make(chan error)
	if err:=godotenv.Load(); err!=nil{
		return err
	}
	wg.Add(1)
	go func(){
		defer wg.Done()
		app.keys.OnLogin(username,password,errchan)
	}()
	go func(){
		wg.Wait()
		defer close(errchan)
	}()
	for err:=range errchan{
		if err!=nil{
			log.Println("Error retrieving user's keys  : ",err)
			return fmt.Errorf("error retrieving account. Make sure the credentials are correct")
		}
	}

	privateKey,err:=wallet.RetriveAccount(username,password);if err!=nil{
		log.Println("Error retrieving user's wallet in : ",err)
		return fmt.Errorf("error retrieving account. Make sure the credentials are correct")

	}

	if err:=blockchain.Init(app.conn,app.instance,privateKey,app.envMap["CONTRACT_ADDR"].(string));err!=nil{
		log.Println("Error connecting to the blockchain : ",err)
		return fmt.Errorf("error connecting to the smart contract")
	}

	app.isApproved,err=app.instance.Instance.IsApprovedInstitute(app.conn.CallOpts,username);if err!=nil{
		log.Println("Error getting the account verification status : ",err)
		return fmt.Errorf("error getting the account verification status")
	}
	log.Println("is approved : ",app.isApproved)
	app.user=username
	return nil
}

func (app *App)Logout(){
	app=&App{}
}

func (app *App)IsLoggedIn()bool{
	return app.conn.TxOpts!=nil
}

func (app *App)IsApprovedInstitute()bool{
	return app.isApproved
}

func (app *App)Register(privateKeyString, name, password string, isInstitute bool) error {
	if len(privateKeyString)<64{
		log.Println("private key error")
		return fmt.Errorf("invalid private key")
	}
	if err:=blockchain.Init(app.conn,app.instance,privateKeyString[2:],app.envMap["CONTRACT_ADDR"].(string));err!=nil{
		return err
	}
	var wg sync.WaitGroup
	errchan:=make(chan error)
	publicKeychan:=make(chan string)	

	wg.Add(2)
	go func() {
		defer wg.Done()
		app.keys.OnRegister(name,password,publicKeychan,errchan)
		
	}()
	go func(){
		defer wg.Done()
		wallet.NewWallet(privateKeyString[2:],name,password,errchan)
	}()
	go func(){
		wg.Wait()
		close(publicKeychan)
		close(errchan)
	}()

	for{
		select{
		case pub,ok:=<-publicKeychan:
			if !ok{
				continue
			}else{
				if isInstitute{
					if err:=app.instance.RegisterInstitution(app.conn.TxOpts,pub,name);err!=nil{
						log.Println("error registering institution : ",err)
						return fmt.Errorf("error registering institution")
					}
					log.Println("Registeration successful")
				}else{
					if err:=app.instance.RegisterUser(app.conn.TxOpts,pub);err!=nil{
						log.Println("error registering user : ",err)
						return fmt.Errorf("error registering institution")
					}
				}
			}
		case err,ok:=<-errchan:
			if err!=nil{
				return err
			}
			if !ok{
				return nil
			}
		}
	}
	
}


func (app *App)UploadDocument(institute,name,description string)error{
	log.Println("Here ",institute,name,description)
	file,err:=app.GetFilePath();if err!=nil{
		log.Println("Error uploading File:",err)
		return fmt.Errorf("Error uploading file")
	}
	log.Println("here file ",file)
	pubKey,err:=app.instance.Instance.GetInstituePublicKey(app.conn.CallOpts,strings.TrimSpace(institute));if err!=nil{
		return err
	}
	if pubKey==""{
		log.Println("error retrieving the name of institute")
		return fmt.Errorf("invalid institution")
	}
	cid,err:=app.ipfs.Upload(file);if err!=nil{
		return err					
	}
	if err:=app.keys.SetMultiSigKey(pubKey);err!=nil{
		return err
	}
	secretKey,err:=app.keys.GenerateSecret();if err!=nil{
		return err
	}
	encryptedDocument,err:=keyUtils.EncryptIPFSHash(secretKey,[]byte(cid));if err!=nil{
		return err
	}
	if err:=app.instance.AddDocument(app.conn.TxOpts,encryptedDocument,institute,name,description); err!=nil{
		return err
	}		
	return nil
}

func (app *App)GetAllDocs()([]blockchain.VerificationDocument,error){
	docs, err := app.instance.GetDocuments(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.TryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester),docs[i].Institute)
	}
	return docs,nil
}

func (app *App) GetAcceptedDocs() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	log.Println("doc id : ",docs[0].ID)
	accepted:=app.FilterStatus(0)
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.TryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester),docs[i].Institute)
		fmt.Println("docs",docs[i])
	}
	verifiedDocs := blockchain.FilterDocument(docs,accepted)
	log.Println("verified : ",verifiedDocs)
	return verifiedDocs, nil
}

func (app *App) GetRejectedDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	rejected:=app.FilterStatus(1)
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.TryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester),docs[i].Institute)
	}
	rejectedDocs := blockchain.FilterDocument(docs,rejected,)

	return rejectedDocs, nil
}

func (app *App) GetPendingDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts); if err != nil {
		return nil, err
	}
	log.Println("doc id : ",docs[0].ID)
	pending:=app.FilterStatus(2)
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.TryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester),docs[i].Institute)
	}
	pendingDocs := blockchain.FilterDocument(docs,pending)
	log.Println("pending : ",pendingDocs)
	return pendingDocs, nil
}

func (app *App)ApproveDocument(status int)error{
	if !app.isApproved{
		return fmt.Errorf("action not approved")
	}
	if err:=app.instance.VerifyDocument(app.conn.TxOpts,app.user,"",uint8(status)); err!=nil{
		log.Println("Error approving document : ",err)
		return fmt.Errorf("error approving document")
	}
	return nil
}
