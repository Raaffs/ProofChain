package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/Suy56/ProofChain/chaincore/core"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/Suy56/ProofChain/storage"
	"github.com/Suy56/ProofChain/users"
	"github.com/Suy56/ProofChain/wallet"
	"github.com/joho/godotenv"
)

// App struct
type App struct {
	ctx      		context.Context
	account			users.User
	keys			*keyUtils.ECKeys
	envMap			map[string]any
	storage 		storage.Document
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
	if err:=blockchain.Init(c,i,privateKey,app.envMap[ENV_CONTRACT_ADDR].(string));err!=nil{
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
	app.account.SetName(username)
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
	if err:=blockchain.Init(c,i,privateKeyString[2:],app.envMap[ENV_CONTRACT_ADDR].(string));err!=nil{
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
					app.account.SetName(name)
					log.Println("registered successful")
				}else{
					requester:=&users.Requester{Conn: c,Instance: i}
					app.account=requester
					if err:=app.account.Register(pub,name);err!=nil{
						log.Println("error registering requester : ",err)
						return fmt.Errorf("error registering institution")
					}
					app.account.SetName(name)
				}
			}
		case err,ok:=<-errchan:
			if err!=nil{
				log.Println("Error registering user: ",err)
				return err
			}
			if !ok{
				return nil
			}
		}
	}
}


func (app *App)UploadDocument(institute,name,description string)error{
	var document storage.Document
	if err:=users.UpdateNonce(app.account);err!=nil{
		log.Println("Invalid transaction nonce: ",err)
		return fmt.Errorf("Invalid transaction nonce")
	}
	filePath,err:=app.GetFilePath();if err!=nil{
		log.Println("Error uploading File:",err)
		return fmt.Errorf("Error uploading file")
	}
	//empty string, cause we don't want public keys of requester
	pubKey,err:=app.account.GetPublicKeys(strings.TrimSpace(institute),"");if err!=nil{
		return err
	}
	if pubKey==""{
		log.Println("error retrieving the name of institute")
		return fmt.Errorf("invalid institution")
	}

	file,err:=os.ReadFile(filePath);if err!=nil{
		log.Println("Error reading file : ",err)
		return err
	}
	//set public key of institution for performing ECDH key exchange
	if err:=app.keys.SetMultiSigKey(pubKey);err!=nil{
		return err
	}
	secretKey,err:=app.keys.GenerateSecret();if err!=nil{
		return err
	}
	encryptedDocument,err:=keyUtils.EncryptIPFSHash(secretKey,file);if err!=nil{
		return err
	}
	shaHash,err:=Keccak256File(filePath); if err!=nil{
		log.Println("Error hashing file:",err)
		return fmt.Errorf("Error uploading file")
	}
	document.EncryptedDocument=encryptedDocument
	document.Shahash=shaHash

	//potential security risk. Since user can directly modify GetPublicAddress function
	//to set any public address
	//Though It won't effect third party verification as the api will directly be sending
	//user's public key stored on the chain, which third party needs to take a hash of to get
	//the public address of user. 
	
	document.PublicAddress=app.account.GetPublicAddress()
	
	if err:=storage.UploadDocument(document);err!=nil{
		log.Println("Error uploading file to mongodb : ",err)
		return fmt.Errorf("Error uploading file")
	}
	if account,ok:=app.account.(*users.Requester);ok{
		if err:=account.Instance.AddDocument(app.account.GetTxOpts(),shaHash,"1",institute,name);err!=nil{
			return err
		}
		return nil
	}
	return fmt.Errorf("invalid account type")
}

func (app *App)GetAllDocs()([]blockchain.VerificationDocument,error){
	var userDocs []blockchain.VerificationDocument
	docs,err:=app.account.GetDocuments(); if err!=nil{
		log.Println("Error getAllDocs:",err)
		return nil,fmt.Errorf("error retrieving documents")
	}
	for i:=range docs{
		//We're calling contract to get public key of institute or requester, however if
		//loggedIn user's address doesn't match with either of them, we don't need to try and drcrypt ipfs cid.
		//This also avoids any unecessary calls to contract
		if docs[i].Institute!=app.account.GetName() && docs[i].Requester!=app.account.GetTxOpts().From.Hex(){
			continue
		}
		userDocs = append(userDocs, docs[i])
	}
	return userDocs,nil
}

func (app *App) GetAcceptedDocs() ([]blockchain.VerificationDocument, error) {
	docs, err := app.account.GetDocuments()
	if err != nil {
		return nil, err
	}
	verifiedDocs := app.account.GetAcceptedDocuments(docs)
	return verifiedDocs, nil
}

func (app *App) GetRejectedDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.account.GetDocuments();if err != nil {
		return nil, err
	}
	rejectedDocs := app.account.GetRejectedDocuments(docs)
	return rejectedDocs, nil
}

func (app *App) GetPendingDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.account.GetDocuments(); if err != nil {
		return nil, err
	}
	pendingDocs := app.account.GetPendingDocuments(docs)
	return pendingDocs, nil
}

func (app *App)ApproveDocument(status int,hash string)error{
	if err:=users.UpdateNonce(app.account);err!=nil{
		log.Println("Invalid transaction nonce: ",err)
		return fmt.Errorf("Invalid transaction nonce")
	}
	//need to look at this again
	if verifier,ok:=app.account.(*users.Verifier);ok{
		if err:=verifier.Instance.VerifyDocument(app.account.GetTxOpts(),hash,verifier.Name,uint8(status));err!=nil{
			log.Println("Error approving document : ",err)
			return fmt.Errorf("error approving document")
		}
		return nil
	}
	return fmt.Errorf("invalid account type")
}

func(app *App)ViewDocument(shahash,instituteName,requesterAddress string)(string,error){
	encryptedDocument,err:=storage.RetrieveDocument(shahash); if err!=nil{
		log.Println("Error retrieving document: ",err)
		return "",fmt.Errorf("Error retrieving document")
	}

	
	decryptedDoc,err:=app.TryDecrypt2(encryptedDocument.EncryptedDocument,instituteName,requesterAddress);if err!=nil{
		log.Println("Error decrypting :",err)
		return "",fmt.Errorf("Error decrypting document")
	}

	encodedDocument:=base64.StdEncoding.EncodeToString(decryptedDoc)

	return encodedDocument, nil
}
