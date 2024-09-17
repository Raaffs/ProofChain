package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Suy56/ProofChain/keyUtils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/crypto/sha3"
)

func(app *App)TryDecrypt2(encryptedIPFS string,institute string,user string)string{
	pub,err:=app.account.GetPublicKeys(institute,user); if err!=nil{
		log.Println("error getting public key: ",err)
		return ""
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
		log.Println("error decrypting ipfs hash here: ",err,ipfs)
		return "private"
	}
	return string(ipfs)
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

func Keccak256File(path string) (string, error) {
	file, err := os.Open(path);if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	hasher := sha3.New256()

	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("failed to hash file: %v", err)
	}

	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}

