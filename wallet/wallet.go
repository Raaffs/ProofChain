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
	var relativePath string
	privateKey, err := crypto.HexToECDSA(privateKeyString[2:])
	if err != nil {
		errchan<-err
		return
	}
	ks := keystore.NewKeyStore("accounts", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		errchan<-err
		return
	}
	path := account.URL.Path
	index := strings.Index(path, "accounts")
	// If "accounts" is found
	if index != -1 {
		// Extract substring from "accounts" to the end
		relativePath = path[index:]
	} else {
		errchan <- errors.New("path not found")
	}
	writeAccountToFile(username, relativePath)
	fmt.Println(relativePath)
	errchan<-nil
}

func writeAccountToFile(username string, fileName string) error {
	accountFile := "accounts/accounts"
	f, err := os.OpenFile(accountFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write to file followed by a new line
	_, err = fmt.Fprintf(f, "%s=%s\n", username, fileName)
	if err != nil {
		return err
	}

	return nil
}

func RetriveAccount(username string, password string) (string, error) {
	filePath := "accounts/accounts"
	err := godotenv.Load(filePath)
	if err != nil {
		return "", err
	}
	accountPath := os.Getenv(username)
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
	return privateKeyHex, nil
}
