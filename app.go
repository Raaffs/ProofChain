package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/Suy56/ProofChain/chaincore/core"
	"github.com/Suy56/ProofChain/crypto/keyUtils"
	"github.com/Suy56/ProofChain/crypto/zkp"
	"github.com/Suy56/ProofChain/storage/models"
	storageclient "github.com/Suy56/ProofChain/storage/storage_client"
	"github.com/Suy56/ProofChain/users"
	"github.com/Suy56/ProofChain/wallet"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

const(
	GANACHE="Ganache"
	INFURA="Infura"
	CLOUDFLARE="CloudFlare"
	DRPC="dRPC"
) 
// App struct
type App struct {
	ctx      		context.Context
	account			users.User
	keys			*keyUtils.ECKeys
	envMap			map[string]any
	storage 		*storageclient.Client
	proof			zkp.ZKProof
	config			Config
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
	if err:=app.config.Load();err!=nil{
		log.Fatalf("Fatal error: loading config failed\n%v",err)
	}
	app.proof=zkp.NewMerkleProof()
	app.storage=storageclient.New(app.config.Services.STORAGE)
}

func (app *App)Login(username string, password string)(error){
	c:=&blockchain.ClientConnection{}
	i:=&blockchain.ContractVerifyOperations{}
	g, _ := errgroup.WithContext(context.Background())
	profile,ok:=app.config.Profiles[username];if !ok{
		return fmt.Errorf("Profile doesn't exist")
	}
	g.Go(func() error {
		if err:=app.keys.OnLogin(username,password,profile.KeyPath);err!=nil{
			return err
		}
		return nil
	})

	g.Go(func() error {
		privateKey,err:=wallet.RetriveAccount(
				username,
				password,
				profile.AccountPath,
			)
			if err!=nil{
			log.Println("Error retrieving user's wallet in : ",err)
			return fmt.Errorf("error retrieving account. Make sure the credentials are correct")
		}
		if err:=blockchain.Init(
				c,
				i,
				privateKey,
				app.config.Services.CONTRACT_ADDR,
				app.config.Services.RPC_PROVIDERS_URLS.Local[GANACHE],
			);err!=nil{
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
	})
	if err := g.Wait(); err != nil {
   		return err
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

func(app *App)Register(privateKeyString,name,password string, isInstitute bool)error{
	if len(privateKeyString)<64{
		log.Println("private key error")
		return fmt.Errorf("invalid private key")
	}
	if _,exist:=app.config.Profiles[name];exist{
		return fmt.Errorf("Profile already exist, please use different name or private key")
	}
	c:=&blockchain.ClientConnection{}
	i:=&blockchain.ContractVerifyOperations{}
	var (
		publicKey string
		accountPath string
		keyPath string
		identityPath string
	) 
	if err:=blockchain.Init(
			c,
			i,
			privateKeyString[2:],
			app.config.Services.CONTRACT_ADDR,
			app.config.Services.RPC_PROVIDERS_URLS.Local[GANACHE],
		);err!=nil{
		return err
	}
	g,_:=errgroup.WithContext(context.Background())
	g.Go(func() error {
		pub, path,err:=app.keys.OnRegister(password,app.config.Dirs.Key); if err!=nil{
			return err
		}
		publicKey=pub
		keyPath=path
		return nil
	})
	g.Go(func() error {
		path,err:=wallet.NewWallet(
				privateKeyString[2:],
				name,password,
				app.config.Dirs.Account,
			)
			if err!=nil{
			return err
		}
		accountPath=path
		return nil
	})
	if err:=g.Wait();err!=nil{
		return err
	}
	
	if err:=app.config.AddProfile(name,accountPath,keyPath,identityPath);err!=nil{
		log.Fatalf("Error creating profile : %v",err)
		return fmt.Errorf("Failed to create profile")
	}

	if isInstitute{
		verifier:=&users.Verifier{Conn: c,Instance: i,Name: name}
		app.account=verifier
		if err:=app.account.Register(publicKey,name);err!=nil{
			log.Println("error registering institution : ",err)
			return fmt.Errorf("error registering institution")
		}
		app.account.SetName(name)
		log.Println("registered successful")
	}else{
		requester:=&users.Requester{Conn: c,Instance: i}
		app.account=requester
		if err:=app.account.Register(publicKey,name);err!=nil{
			log.Println("error registering requester : ",err)
			return fmt.Errorf("error registering institution")
		}
		app.account.SetName(name)
	}
	return nil
}

func (app *App)UploadDocument(institute,name,description string)error{
	var document models.Document
	if err:=users.UpdateNonce(app.account);err!=nil{
		log.Println("Invalid transaction nonce: ",err)
		return fmt.Errorf("invalid transaction nonce")
	}
	file,path,err:=app.GetFileAndPath();if err!=nil{
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
	encryptedDocument,err:=app.Encrypt(file,pubKey);if err!=nil{
		log.Println(err)
		return fmt.Errorf("An error occurred while encrypting document")
	}
	shaHash,err:=Keccak256File(path); if err!=nil{
		log.Println("Error hashing file:",err)
		return fmt.Errorf("Error uploading file")
	}
	document.EncryptedDocument=encryptedDocument
	document.Shahash=shaHash
	document.PublicAddress=app.account.GetPublicAddress()
	if account,ok:=app.account.(*users.Requester);ok{
		if err:=account.Instance.AddDocument(app.account.GetTxOpts(),shaHash,institute);err!=nil{
			return err
		}
		return nil
	}
	if err:=app.storage.UploadDocument(document);err!=nil{
		log.Println("Error uploading file to mongodb : ",err)
		return fmt.Errorf("Error uploading file")
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

func (app *App)CreateDigitalCopy(status int,hash string, certificate models.CertificateData)error{
	if status==2{
		if verifier,ok:=app.account.(*users.Verifier);ok{
			if err:=verifier.Instance.VerifyDocument(app.account.GetTxOpts(),hash,verifier.Name,uint8(status),hash);err!=nil{
				log.Println("Error approving document : ",err)
				return fmt.Errorf("An error occurred ")
			}
			return nil
		}
	}
	doc,publicCommit,err:=app.PrepareDigitalCopy(certificate);if err!=nil{
		log.Println(err)
		return fmt.Errorf("An error occurred while issuing document")
	}
	if verifier,ok:=app.account.(*users.Verifier);ok{
		if err:=verifier.Instance.VerifyDocument(app.account.GetTxOpts(),hash,verifier.Name,uint8(status),publicCommit);err!=nil{
			log.Println("Error approving document : ",err)
			return fmt.Errorf("error approving document")
		}
		return nil
	}

	if err:=app.storage.UploadDocument(doc);err!=nil{
		log.Println(err)
		return fmt.Errorf("Error creating certificate")
	}
	return nil
}

func (app *App)IssueCertificate(certificate models.CertificateData)error{
	doc,publicCommit,err:=app.PrepareDigitalCopy(certificate);if err!=nil{
		log.Println(err)
		return fmt.Errorf("An error occurred while issuing document")
	}
	if err:=app.account.AddDocument(
		string(publicCommit),
		app.account.GetName(),
	);err!=nil{
		log.Println(err)
		return fmt.Errorf("Error issuing certificate")
	}
	if err:=app.storage.UploadDocument(doc);err!=nil{
		log.Println(err)
		return fmt.Errorf("Error creating certificate")
	}
	return nil
}

func(app *App)ViewDocument(shahash,instituteName,requesterAddress string)(string,error){
	encryptedDocument,err:=app.storage.RetrieveDocument(shahash); if err!=nil{
		log.Println("Error retrieving document: ",err)
		return "",fmt.Errorf("Error retrieving document")
	}

	decryptedDoc,err:=app.TryDecrypt(encryptedDocument.EncryptedDocument,instituteName,requesterAddress);if err!=nil{
		log.Println("Error decrypting :",err)
		return "",fmt.Errorf("Error decrypting document")
	}
	encodedDocument:=base64.StdEncoding.EncodeToString(decryptedDoc)
	return encodedDocument, nil
}