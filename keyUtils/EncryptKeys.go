package keyUtils

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gtank/cryptopasta"
	"github.com/joho/godotenv"
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

func EncryptPrivateKeyFile(data, account, passphrase string) ([]byte, error) {
	salt := make([]byte, 10)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	key,err := generateKeyFromPassphrase(passphrase,salt);if err!=nil{
		return nil,err
	}
	encryptedData, err := cryptopasta.Encrypt([]byte(data), key)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt data: %v", err)
	}
	encryptedDataWithSalt := append(salt, encryptedData...)
	WKey(account, string(encryptedDataWithSalt))
	return encryptedDataWithSalt, nil
}

func RKey(filepath string)(string,error){
	keyBytes,err:=os.ReadFile(filepath)
	if err!=nil{
		return "",err
	}
	return string(keyBytes),nil
}
func ReadKeyMap(user string)(string,error){
	keyMapPath:="keys"
	if err:=godotenv.Load(keyMapPath);err!=nil{
		return "",err
	}
	KeyFile:=os.Getenv(user)
	encryptedKey,err:=RKey(KeyFile)
	if err!=nil{
		return "",err
	}
	return encryptedKey,nil
}
func WriteKeyMap(user,keyFile string)error{
	keyFileMap:="keys"
	err:=godotenv.Load(keyFileMap)
	if err!=nil{
		return err
	}
	value:=os.Getenv(user)
	if value!=""{
		return errors.New("user already exist")
	}
	f, err := os.OpenFile(keyFileMap, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
			return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "%s=%s\n", user, keyFile)
	if err != nil {
			return err
	}
	return nil
}

func WKey(user,encryptedData string)(error){
	filePath:=fmt.Sprintf("%d",time.Now().UnixNano()) 
	err:=WriteKeyMap(user,filePath)
	if err!=nil{
		return err
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
			return err
	}
	defer f.Close()
	_,err=f.Write([]byte(encryptedData)); if err!=nil{
		return err
	}
	return nil
}

func DecryptPrivateKeyFile(user,passphrase string) ([]byte, error) {
	encryptedDataWithSalt,err:=ReadKeyMap(user); if err!=nil{
		return nil,err
	}
	salt := encryptedDataWithSalt[:10]
	encryptedData := encryptedDataWithSalt[10:]
	key, err := generateKeyFromPassphrase(passphrase, []byte(salt))
	if err != nil {
		return nil, err
	}
	decryptedData, err := cryptopasta.Decrypt([]byte(encryptedData), key)
    if err != nil {
        return nil, err
    }
	fmt.Println(decryptedData)
    return decryptedData, nil
}

