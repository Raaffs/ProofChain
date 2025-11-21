package keyUtils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gtank/cryptopasta"
	"golang.org/x/crypto/pbkdf2"
)

func GetECDSAPrivateKeyFromPEM(pk string)(*ecdh.PrivateKey,error){
	block, _ := pem.Decode([]byte(pk))
if block == nil || block.Type!= "PRIVATE KEY" {
        log.Fatal("Failed to decode PEM block containing private key")
    }
    privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
    if err!= nil {
        log.Fatalf("Error parsing private key: %v", err)
    }
    ecdsaPrivateKey, ok := privateKey.(*ecdsa.PrivateKey)
    if!ok {
        log.Fatal("Expected *ecdh.PrivateKey, got another type : ",ok)
    }
	ecdhPrivateKey,err:=ecdsaPrivateKey.ECDH()
	if err!=nil{
		return nil,err
	}
	return ecdhPrivateKey,nil
}

//The public key string retrieved from smart contract is converted to public key object
func GetECDSAPublicKeyFromPEM(rawPublicKey string) (*ecdh.PublicKey, error) {
	block, _ := pem.Decode([]byte(rawPublicKey))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}
	publicKey, ok := parsedKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to convert to *ecdsa.PublicKey")
	}
	ecdhPubKey,err:=publicKey.ECDH()
	if err!=nil{
		log.Println("Error getting ecdh public key : ",err)
		return nil,err
	}
	return ecdhPubKey, nil
}

func generateKeyFromPassphrase(passphrase string,salt []byte) (*[32]byte,error) {
	key := pbkdf2.Key([]byte(passphrase), salt, 10000, 32, sha256.New)
	var keyArray [32]byte
	copy(keyArray[:], key)
	return &keyArray,nil
}


func EncryptPrivateKeyFile(privateKey, password, keyDir string) ( string, error) {
	salt := make([]byte, 10)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	key,err := generateKeyFromPassphrase(password,salt);if err!=nil{
		return "",err
	}
	encryptedData, err := cryptopasta.Encrypt([]byte(privateKey), key)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt data: %v", err)
	}
	encryptedDataWithSalt := append(salt, encryptedData...)
	filePath,err:=WKey(string(encryptedDataWithSalt),keyDir)
	if err!=nil{
		return "",err
	}
	return filePath, nil
}

func DecryptPrivateKeyFile(user, passphrase, path string) ([]byte, error) {
	encryptedDataWithSalt,err:=RKey(path); if err!=nil{
		return nil,err
	}
	salt := encryptedDataWithSalt[:10]
	encryptedData := encryptedDataWithSalt[10:]
	key, err := generateKeyFromPassphrase(passphrase, salt)
	if err != nil {
		return nil, err
	}
	decryptedData, err := cryptopasta.Decrypt([]byte(encryptedData), key)
    if err != nil {
		log.Println("Error decrypting file ",err)
        return nil, err
    }
    return decryptedData, nil
}


//Just a helper function to read files 
func RKey(filepath string)([]byte,error){
	keyBytes,err:=os.ReadFile(filepath)
	if err!=nil{
		return nil,err
	}
	return (keyBytes),nil
}


//helper function to write encrypted private key to file
//file name is generated based on time
func WKey(encryptedData,keyDir string)(string,error){
	homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }

	filePath:=filepath.Join(homeDir,keyDir,fmt.Sprintf("%d",time.Now().UnixNano()) ) 
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
			return "",err
	}
	defer f.Close()
	_,err=f.Write([]byte(encryptedData)); if err!=nil{
		return "",err
	}
	return filePath,nil
}




func EncryptIPFSHash(sharedKey []byte, plaintext []byte)([]byte,error){
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		return nil, err
	}

	// Generate a new AES-GCM instance
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Create a nonce. Nonce size should be equal to gcm.NonceSize()
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the data
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func DecryptIPFSHash(sharedKey []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		log.Println("error decrypting ipfs hash : ",err)
		return nil, err
	}

	// Generate a new AES-GCM instance
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("error decrypting ipfs hash : ",err)
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	// Extract the nonce and the actual ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println("error decrypting ipfs hash : ",err)
		return nil, err
	}

	return plaintext, nil
}