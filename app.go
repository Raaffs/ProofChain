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
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	app.user=username
	return nil
}

func (app *App)Logout(){
	app=&App{}
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

func (app *App)GetFilePath(next func(string)error)(string,error){
	filePath,err:=runtime.OpenFileDialog(app.ctx,runtime.OpenDialogOptions{
		Title: "Select Document",
		Filters: []runtime.FileFilter{	
			{
				DisplayName: "Documents (*.pdf; *.jpg; *.png)",
				Pattern: "*.pdf;*.jpg;*.png",
			},
		},
	})
	if err!=nil{
		return "",err
	}
	if err:=next(filePath);err!=nil{
		return "",err
	}
	return filePath,nil
}

func (app *App)UploadDocument(institute,name,description string)error{
	app.GetFilePath(
		func(file string) error {
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
		},
	)
	return nil
}

func (app *App)GetAllDocs()([]blockchain.VerificationDocument,error){
	docs, err := app.instance.GetDocuments(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.tryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester),docs[i].Institute)
	}
	return docs,nil
}

func (app *App) GetAcceptedDocs() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	accepted:=app.rt(0)
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.tryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester),docs[i].Institute)
	}
	verifiedDocs := blockchain.FilterDocument(docs,accepted,app.conn.CallOpts.From.Hex())
	return verifiedDocs, nil
}

func (app *App) GetRejectedDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	rejected:=app.rt(1)
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.tryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester),docs[i].Institute)
	}
	rejectedDocs := blockchain.FilterDocument(docs,rejected, app.conn.CallOpts.From.Hex())
	return rejectedDocs, nil
}

func (app *App) GetPendingDocuments()([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	pending:=app.rt(2)
	for i:=0;i<len(docs);i++{
		docs[i].IpfsAddress=app.tryDecrypt(docs[i].IpfsAddress,common.HexToAddress(docs[i].Requester),docs[i].Institute)
	}
	pendingDocs := blockchain.FilterDocument(docs,pending,app.conn.CallOpts.From.Hex())
	return pendingDocs, nil
}

func (app *App)rt(status int)func(blockchain.VerificationDocument, string)bool{
	return func(doc blockchain.VerificationDocument,req string)bool{
		if app.isApproved{
			return doc.Institute==req && int(doc.Stats)==status
		}
		return doc.Requester==req && int(doc.Stats)==status
	}
}


func(app *App)tryDecrypt(encryptedIPFS string,user common.Address,institute string )string{
	pub,err:=app.GetPublicKeys(user,institute);if err!=nil{
		log.Println("error getting public key: ",err)
		return ""
	}
	if err:=app.keys.SetMultiSigKey(pub);err!=nil{
		log.Println("error setting multisigkey: ",err)
		return ""
	}
	sec,err:=app.keys.GenerateSecret();if err!=nil{
		log.Println("error generating secret: ",err)
		return ""
	}
	ipfs,err:=keyUtils.DecryptIPFSHash(sec,[]byte(encryptedIPFS));if err!=nil{
		log.Println("error decrypting ipfs hash: ",err,ipfs)
		return ""
	}
	return ipfs
}

func(app *App)GetPublicKeys(user common.Address,institute string)(string,error){
	if app.isApproved{
		pub,err:=app.instance.Instance.GetUserPublicKey(app.conn.CallOpts,user);if err!=nil{
			log.Println("Error getting public key of institute : ",err)
			return "",err
		}
		return pub,nil
	}
	pub,err:=app.instance.Instance.GetInstituePublicKey(app.conn.CallOpts,institute);if err!=nil{
		log.Println("Error getting public key of user : ",err)
		return "",err
	}
	return pub,nil
}

