package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/crypto/sha3"
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
		return "private"
	}
	if err:=app.keys.SetMultiSigKey(pub);err!=nil{
		log.Println("error setting multisigkey: ",err)
		return "private"
	}
	sec,err:=app.keys.GenerateSecret();if err!=nil{
		log.Println("error generating secret: ",err)
		return "private"
	}
	ipfs,err:=keyUtils.DecryptIPFSHash(sec,[]byte(encryptedIPFS));if err!=nil{
		log.Println("error decrypting ipfs hash: ",err,ipfs)
		return "private"
	}
	return string(ipfs)
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

func Keccak256File(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a new Keccak-256 hasher
	hasher := sha3.New256()

	// Read the file's contents into the hasher
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("failed to hash file: %v", err)
	}

	// Compute the hash and convert it to a hex string
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}