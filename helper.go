package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/Suy56/ProofChain/crypto/keyUtils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/crypto/sha3"
)



func (app *App)Encrypt(file []byte, publicKey string)([]byte,error){
	if err:=app.keys.SetMultiSigKey(publicKey);err!=nil{
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
	pub,err:=app.account.GetPublicKeys(institute,user); if err!=nil{
		log.Println("error getting public key: ",err)
		return nil,fmt.Errorf("Error retrieving public keys")
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

func Walk[S any](s S) func(yield func(string, any) bool) {
	v := reflect.ValueOf(s)

	// Dereference pointer if needed
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return func(yield func(string, any) bool) {}
	}

	return func(yield func(string, any) bool) {
		t := v.Type()
		numFields := v.NumField()

		for i := range numFields {
			field := t.Field(i)
			value := v.Field(i)

			if !field.IsExported() {
				continue
			}

			switch value.Kind() {
			case reflect.Map:
				// Iterate map keys
				for _, key := range value.MapKeys() {
					val := value.MapIndex(key)
					if !yield(fmt.Sprint(key.Interface()), val.Interface()) {
						return
					}
				}
			default:
				if !yield(field.Name, value.Interface()) {
					return
				}
			}
		}
	}
}

func SaveProof(identityDir, certificateName string, proofs map[string]any) error {
	if certificateName == "" {
		return fmt.Errorf("certificate name cannot be empty")
	}
	if err := os.MkdirAll(identityDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	path := filepath.Join(identityDir, certificateName+".json")
	
	data, err := json.MarshalIndent(proofs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal proofs: %w", err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("failed to write proofs file: %w", err)
	}
	return nil
}
