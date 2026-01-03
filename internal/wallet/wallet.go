package wallet

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)


func NewWallet(privateKeyString, username, password, accountDir string )(string,error){
	homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }

    // Ensure the directory exists
    fullDir := filepath.Join(homeDir, accountDir)
    if err := os.MkdirAll(fullDir, 0700); err != nil {
        return "", err
    }

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return "", err
	}
	ks := keystore.NewKeyStore(fullDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return "", err
	}
	return account.URL.Path,nil
}

func RetriveAccount(username, password, path string) (string, error) {
	if path==""{
		return "",fmt.Errorf("account not found")
	}
	accountFile, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	privateKey, err := keystore.DecryptKey(accountFile, password)
	if err != nil {
		return "", fmt.Errorf("Error in retriving Account %w",err)
	}
	privateKeyBytes := privateKey.PrivateKey.D.Bytes()
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	return privateKeyHex,nil
}
