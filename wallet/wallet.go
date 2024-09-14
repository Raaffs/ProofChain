package wallet

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

func NewWallet(privateKeyString string, username string, password string, errchan chan error) {
	accountMap,err:=godotenv.Read("accounts/accounts"); if err!=nil{
		errchan<-err
		return
	}
	accountPath:=accountMap[username]
	if accountPath!=""{
		fmt.Println(accountMap)
		errchan<-fmt.Errorf("account already exist")
		return
	}

	var relativePath string
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		errchan<-err
		return
	}
	ks := keystore.NewKeyStore("accounts", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		fmt.Println(accountMap)

		errchan<-err
		return
	}
	path := account.URL.Path
	fmt.Println("Path : ",path)
	index := strings.Index(path, "accounts")
	// If "accounts" is found
	if index != -1 {
		// Extract substring from "accounts" to the end
		relativePath = path[index:]
	} else {
		errchan <- errors.New("path not found")
	}
	writeAccountToFile(username, relativePath)

	errchan<-nil
}

func writeAccountToFile(username string, fileName string) error {
	accountDir := "accounts"
	accountFile := accountDir + "/accounts"

	err := os.MkdirAll(accountDir, 0755)
	if err != nil {
		return err
	}

	// Open or create the file for appending and writing
	f, err := os.OpenFile(accountFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%s=%s\n", username, fileName)
	if err != nil {
		return err
	}

	return nil
}


func RetriveAccount(username string, password string) (string, error) {
	// filePath := "accounts/accounts"
	accountMap,err := godotenv.Read("accounts/accounts")
	if err != nil {
		return "", err
	}
	accountPath := accountMap[username]
	if accountPath==""{
		return "",fmt.Errorf("account not found")
	}
	accountFile, err := os.ReadFile(accountPath)
	if err != nil {
		return "", err
	}

	privateKey, err := keystore.DecryptKey(accountFile, password)
	if err != nil {
		fmt.Println("Error in retriving Account")
		return "", err
	}
	privateKeyBytes := privateKey.PrivateKey.D.Bytes()
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	return privateKeyHex,nil
}
