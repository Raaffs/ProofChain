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
	"github.com/Suy56/ProofChain/users"
	"github.com/Suy56/ProofChain/wallet"
	"github.com/joho/godotenv"
)
// App struct
type App struct {
	ctx      		context.Context
	account			users.User
	keys			*keyUtils.ECKeys
	ipfs 			*nodeData.IPFSManager
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
	c:=&blockchain.ClientConnection{}
	i:=&blockchain.ContractVerifyOperations{}
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
	if err:=blockchain.Init(c,i,privateKey,app.envMap["CONTRACT_ADDR"].(string));err!=nil{
		log.Println("Error connecting to the blockchain : ",err)
		return fmt.Errorf("error connecting to the smart contract")
	}
	approved,err:=i.Instance.IsApprovedInstitute(c.CallOpts,username); if err!=nil{
		log.Println("Error getting the account verification status : ",err)
		return fmt.Errorf("error getting the account verification status")
	}
	if approved{
		app.account=&users.Verifier{Conn: c,Instance: i,Name: username}
	}else{
		app.account=&users.Requester{Conn: c,Instance: i}
	}
	return nil
}

func (app *App)Logout(){
	app=&App{}
}

func (app *App)IsLoggedIn()(bool){
	return app.account.GetTxOpts()!=nil
}

func (app *App)IsApprovedInstitute()bool{
	approved,err:=app.account.GetApprovalStatus();if err!=nil{
		log.Println("Error getting approval status : ",err)
		return false
	}
	return approved
}

func (app *App)Register(privateKeyString, name, password string, isInstitute bool) error {
	if len(privateKeyString)<64{
		log.Println("private key error")
		return fmt.Errorf("invalid private key")
	}
	c:=&blockchain.ClientConnection{}
	i:=&blockchain.ContractVerifyOperations{}
	if err:=blockchain.Init(c,i,privateKeyString[2:],app.envMap["CONTRACT_ADDR"].(string));err!=nil{
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
					verifier:=&users.Verifier{Conn: c,Instance: i,Name: name}
					app.account=verifier
					if err:=app.account.Register(pub,name);err!=nil{
						log.Println("error registering institution : ",err)
						return fmt.Errorf("error registering institution")
					}
					log.Println("registered successful")
				}else{
					requester:=&users.Requester{Conn: c,Instance: i}
					app.account=requester
					if err:=app.account.Register(pub,name);err!=nil{
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
	if err:=users.UpdateNonce(app.account);err!=nil{
		log.Println("Invalid transaction nonce: ",err)
		return fmt.Errorf("Invalid transaction nonce")
	}
	log.Println("Here ",institute,name,description)
	file,err:=app.GetFilePath();if err!=nil{
		log.Println("Error uploading File:",err)
		return fmt.Errorf("Error uploading file")
	}
	pubKey,err:=app.account.GetPublicKeys(strings.TrimSpace(institute),"");if err!=nil{
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
	shaHash,err:=Keccak256File(file); if err!=nil{
		log.Println("Error hashing file:",err)
		return fmt.Errorf("Error uploading file")
	}
	if account,ok:=app.account.(*users.Requester);ok{
		if err:=account.Instance.AddDocument(app.account.GetTxOpts(),shaHash,encryptedDocument,institute,name);err!=nil{
			return err
		}
		return nil
	}
	return fmt.Errorf("invalid account type")
}

func (app *App)GetAllDocs()([]blockchain.VerificationDocument,error){
	docs,err:=app.account.GetDocuments(); if err!=nil{
		log.Println("Error getAllDocs:",err)
		return nil,fmt.Errorf("error retrieving documents")
	}
	for i:=0;i<len(docs);i++{
		//We're calling contract to get public key of institute or requester, however if
		//loggedIn user's address doesn't match with either of them, we don't need to try and drcrypt ipfs cid.
		//This also avoids any unecessary calls to contract
		if docs[i].Verifier!=app.account.GetTxOpts().From.Hex() && docs[i].Requester!=app.account.GetTxOpts().From.Hex(){
			docs[i].IpfsAddress="not authorized"
			continue
		}
		docs[i].IpfsAddress=app.TryDecrypt2(docs[i].IpfsAddress,docs[i].Institute,docs[i].Requester)
	}
	return docs,nil
}

func (app *App) GetAcceptedDocs() ([]blockchain.VerificationDocument, error) {
	docs, err := app.account.GetDocuments()
	if err != nil {
		return nil, err
	}
	for i:=0;i<len(docs);i++{
		if docs[i].Verifier!=app.account.GetTxOpts().From.Hex() && docs[i].Requester!=app.account.GetTxOpts().From.Hex(){
			docs[i].IpfsAddress="not authorized"
			continue
		}
		docs[i].IpfsAddress=app.TryDecrypt2(docs[i].IpfsAddress,docs[i].Institute,docs[i].Requester)
	}
	verifiedDocs := app.account.GetAcceptedDocuments(docs)
	return verifiedDocs, nil
}

func (app *App) GetRejectedDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.account.GetDocuments();if err != nil {
		return nil, err
	}
	for i:=0;i<len(docs);i++{
		if docs[i].Verifier!=app.account.GetTxOpts().From.Hex() && docs[i].Requester!=app.account.GetTxOpts().From.Hex(){
			docs[i].IpfsAddress="not authorized"
			continue
		}
		docs[i].IpfsAddress=app.TryDecrypt2(docs[i].IpfsAddress,docs[i].Institute,docs[i].Requester)
	}
	rejectedDocs := app.account.GetRejectedDocuments(docs)
	return rejectedDocs, nil
}

func (app *App) GetPendingDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.account.GetDocuments(); if err != nil {
		return nil, err
	}
	for i:=0;i<len(docs);i++{
		if docs[i].Verifier!=app.account.GetTxOpts().From.Hex() && docs[i].Requester!=app.account.GetTxOpts().From.Hex(){
			docs[i].IpfsAddress="not authorized"
			continue
		}
		docs[i].IpfsAddress=app.TryDecrypt2(docs[i].IpfsAddress,docs[i].Institute,docs[i].Requester)
	}
	pendingDocs := app.account.GetPendingDocuments(docs)
	return pendingDocs, nil
}

func (app *App)ApproveDocument(status int,hash string)error{
	if err:=users.UpdateNonce(app.account);err!=nil{
		log.Println("Invalid transaction nonce: ",err)
		return fmt.Errorf("Invalid transaction nonce")
	}
	log.Println("document hash : ",hash)
	if verifier,ok:=app.account.(*users.Verifier);ok{
		if err:=verifier.Instance.VerifyDocument(app.account.GetTxOpts(),hash,verifier.Name,uint8(status));err!=nil{
			log.Println("Error approving document : ",err)
			return fmt.Errorf("error approving document")
		}
		return nil
	}
	return fmt.Errorf("invalid account type")
}
