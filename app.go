package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/Suy56/ProofChain/wallet"
	"github.com/ethereum/go-ethereum/common"
)

// App struct
type App struct {
	ctx      		context.Context
	conn		  	*blockchain.ClientConnection
	instance 		*blockchain.ContractVerifyOperations
	keys			*keyUtils.ECKeys
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		conn: 		&blockchain.ClientConnection{},
		instance: 	&blockchain.ContractVerifyOperations{},
		keys: 		&keyUtils.ECKeys{},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

func (app *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (app *App) Login(username string, password string) (error) {
	privateKey, err := wallet.RetriveAccount(username, password)
	if err != nil {
		return err
	}
	err = app.conn.New(privateKey)
	if err != nil {
		return  err
	}
	app.instance.Client = app.conn.Client
	err = app.instance.New("")
	if err != nil {
		return err
	}

	return nil
}

func (app *App) LoginUserTest() bool {
	return true
}

func (app *App) LoginVerifierTest() bool {
	return true
}

// Concurrent go routines needs to be fixed
// Also add proper error handling
func (app *App) Register(privateKeyString, username, password string) error {
	var wg sync.WaitGroup
	errchan:=make(chan error)
	wg.Add(2)
	fmt.Println(username,password)
	go func() {
		defer wg.Done()
		app.keys.OnRegister(username,password,errchan)
		
	}()
	go func(){
		defer wg.Done()
		wallet.NewWallet(privateKeyString,username,password,errchan)
	}()

	go func(){
		wg.Wait()
		defer close(errchan)
	}()

	for err:=range errchan{
		if err!=nil{
			return err
		}
	}

	
	// if err:=app.conn.New(privateKeyString);err!=nil{
	// 	return err
	// }

	// // Note: The current approach for setting the Client instance is not be optimal.
	// // I tried  to use a method returning *ethclient.Client,
	// // but it resulted in a nil pointer error.
	// // Future improvements may involve revisiting the method approach.
	// app.instance.Client = app.conn.Client

	// if err:=app.conn.New(privateKeyString);err!=nil{
	// 	return err
	// }

	// if err:=app.instance.RegisterUser(app.conn.TxOpts,username,password);err!=nil{
	// 	return err
	// }
	return nil
}

func (app *App) GetAcceptedDocs() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
		return doc.Requester == requester && doc.Stats == 0
	}, app.conn.CallOpts.From)
	fmt.Println("Verified docs : ", verifiedDocs)
	return verifiedDocs, nil
}

func (app *App) GetRejectedDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
		return doc.Requester == requester && doc.Stats == 1
	}, app.conn.CallOpts.From)
	fmt.Println("Verified docs : ", verifiedDocs)
	return verifiedDocs, nil

}

func (app *App) GetPendingDocuments() ([]blockchain.VerificationDocument, error) {
	docs, err := app.instance.GetDocuments(app.conn.CallOpts)
	if err != nil {
		return nil, err
	}
	verifiedDocs := blockchain.FilterDocument(docs, func(doc blockchain.VerificationDocument, requester common.Address) bool {
		return doc.Requester == requester && doc.Stats == 3
	}, app.conn.CallOpts.From)
	fmt.Println("Verified docs : ", verifiedDocs)
		return verifiedDocs, nil
}

