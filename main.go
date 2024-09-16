package main

import (
	"embed"

	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/Suy56/ProofChain/nodeData"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := &App{
		keys: 	  		&keyUtils.ECKeys{},
		ipfs: 	  		&nodeData.IPFSManager{},
		envMap: 		make(map[string]any),
	}
	
	err := wails.Run(&options.App{
		Title:  "ProofChainV2",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
	if err:=godotenv.Load(".env","keys/keyMap","accounts/accounts");err!=nil{
		println("Error : ",err.Error())
	}
}
