package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Suy56/ProofChain/crypto/keyUtils"
	"github.com/Suy56/ProofChain/storage/models"
	"github.com/Suy56/ProofChain/users"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/crypto/sha3"
)



// func (app *App)autoResvolvePublicKey(target string)(string,error)

func (app *App)Encrypt(file []byte, entity string)([]byte,error){
	pubKey,err:=app.account.GetPublicKeys(entity);if err!=nil{
		return nil,err
	}
	if pubKey==""{
		log.Println("error retrieving public keys")
		return nil,fmt.Errorf("invalid institution")
	}
	if err:=app.keys.SetMultiSigKey(pubKey);err!=nil{
		return nil,err
	}
	secretKey,err:=app.keys.GenerateSecret();if err!=nil{
		return nil,err
	}
	encryptedDocument,err:=keyUtils.Encrypt(secretKey,file);if err!=nil{
		return nil,err
	}
	return encryptedDocument,nil
}	

func(app *App)TryDecrypt(encryptedDocument []byte,institute string,user string)([]byte,error){
	var targetEntity string
	if _,ok:=app.account.(*users.Requester); ok{
		targetEntity=institute
	}
	if _,ok:=app.account.(*users.Verifier); ok{
		targetEntity=user
	}

	pub,err:=app.account.GetPublicKeys(targetEntity); if err!=nil{
		return nil,fmt.Errorf("helper.go: error retrieving public keys %w",err)
	}
	log.Println("public key of ins: ",pub)
	if err:=app.keys.SetMultiSigKey(pub);err!=nil{
		log.Println("error setting multisigkey: ",err)
		return nil,fmt.Errorf("Error retrieving multi-sig keys")
	}
	sec,err:=app.keys.GenerateSecret();if err!=nil{
		log.Println("error generating secret: ",err)
		return nil,fmt.Errorf("Error generating secret key")
	}
	document,err:=keyUtils.Decrypt(sec,encryptedDocument);if err!=nil{
		log.Println("error decrypting ipfs hash here: ",err)
		return nil,fmt.Errorf("You're not authorized")
	}
	return document,nil
}

func (app *App)GetFileAndPath()([]byte, string, error){
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
		return nil,"",err
	}
	file,err:=os.ReadFile(filePath);if err!=nil{
		log.Println("Error reading file : ",err)
		return nil,"",err
	}
	return file,filePath,nil
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

func (app *App)IsApprovedInstitute()bool{
	approved,err:=app.account.GetApprovalStatus();if err!=nil{
		log.Println("Error getting approval status : ",err)
		return false
	}
	return approved
}

func (app *App)PrepareDigitalCopy(certificate models.CertificateData)(models.Document,string,error){
	publicCommit,saltedCertificate,err:=app.proof.GenerateRootProof(certificate);if err!=nil{
		log.Println(err)
		return models.Document{},"",fmt.Errorf("an error occurred while issuing certificate")
	}
	log.Println("public commit: ",publicCommit)
	json, err := json.Marshal(saltedCertificate);if err!=nil{
		log.Println("Error marshaling the certificate: ",err)
		return models.Document{},"",fmt.Errorf("invalid certificate format")
	}
	encryptedCertificate,err:=app.Encrypt(json,certificate.PublicAddress);if err!=nil{
		return models.Document{},"",fmt.Errorf("error encrypting document %w",err)
	}
	doc:=models.Document{
		Shahash: string(publicCommit),
		EncryptedDocument: encryptedCertificate,
		PublicAddress: certificate.PublicAddress,
	}
	return doc,string(publicCommit),nil
}

