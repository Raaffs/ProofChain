package main

import (
	"context"
	"fmt"

	"github.com/Suy56/ProofChain/wallet"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

// Greet returns a greeting for the given name
func (app *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (app *App)Login(username string, password string)(string,error){
	privateKey,err:=wallet.RetriveAccount(username,password)
	if err!=nil{
		return "",err
	}
	return privateKey,nil
}

func (app *App)LoginUserTest() bool{
	return true
}

func (app *App)LoginVerifierTest()bool{
	return true
}

func (app *App)Register(privateKeyString, username,password string)error{
	err:=wallet.NewWallet(privateKeyString,username,password)
	if err!=nil{
		return err
	}
	return nil
}

func(app *App)getVerifiedDocuments(){
	
}

func(app *App)getRejectedDocuments(){
	
}