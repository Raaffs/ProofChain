package main

import (
	"log"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)


func (app *App)FilterStatus(status int)func(blockchain.VerificationDocument)bool{
	return func(doc blockchain.VerificationDocument)bool{
		//doc.Institute is basically the name of institute.
		//So instead of common.Address we just need the name that is stored in app.user
		//when user is logged in
		if app.isApproved{
			return doc.Institute==app.user && int(doc.Stats)==status
		}
		return doc.Requester==app.conn.CallOpts.From.Hex() && int(doc.Stats)==status
	}
}


func(app *App)TryDecrypt(encryptedIPFS string,user common.Address,institute string )string{
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

func (app *App)GetFilePath()(string,error){
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
	return filePath,nil
}
