package keyUtils

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)
type ECKeys struct{
	Private 	*ecdh.PrivateKey
	Public		*ecdh.PublicKey

	//Public key of instituition/user is retrieved from blockchain 
	//for generating shared secret for AES encryption
	MultiSig	*ecdh.PublicKey 
}

//Since ECDSA Keys are only used for signing.
//An ECDH public-private key pair is auto-generated to 
//encrypt the ipfs hash, so that only user and instituion can view the document
func(k *ECKeys)OnLogin(user string,passphrase string, errchan chan error,){
	pemKey,err:=DecryptPrivateKeyFile(user,passphrase);if err!=nil{
		errchan<-err
		return
	}
	k.Private,err=GetECDSAPrivateKeyFromPEM(string(pemKey)); if err!=nil{
		errchan<-err
		return
	}
	k.Public=k.Private.PublicKey()
	errchan<-nil
}


func (k *ECKeys)OnRegister(username,password string,publicKey chan string,errchan chan error){
	privECDSA,err:=ecdsa.GenerateKey(elliptic.P256(),rand.Reader);if err!=nil{
		errchan<- err
		return
	}
	k.Private,err=privECDSA.ECDH();if err!=nil{
		errchan<- err
		return
	}
	k.Public=k.Private.PublicKey()
	pemPrivateKey,err:=k.MarshalECDHPrivateKey();if err!=nil{
		errchan<- err
		return
	}
	pemPublicKey,err:=k.MarshalECDHPublicKey();if err!=nil{
		errchan<-err
	}
	_,err=EncryptPrivateKeyFile(pemPrivateKey,username,password)
	if err!=nil{
		errchan<- err
		return
	}
	publicKey<-pemPublicKey
}

func(k *ECKeys)SetMultiSigKey(multiSigKey string)error{
	var err error
	k.MultiSig,err=GetECDSAPublicKeyFromPEM(multiSigKey); if err!=nil{
		return err
	}
	return nil
}


//Shared Secret is used for AES encrytion/decryption 
func(k *ECKeys)GenerateSecret()([]byte,error){
	if k.MultiSig==nil || k.Private==nil{
		log.Println("multisig or private key is nil")
		return nil,fmt.Errorf("Multi-Sig-Public-Key not provided")
	}
	secret,err:=k.Private.ECDH(k.MultiSig); if err!=nil{
		return nil,err
	}
	return secret,nil
}

func(k * ECKeys)MarshalECDHPublicKey()(string,error){
	ecdhSKBytes, err := x509.MarshalPKIXPublicKey(k.Public)
	if err != nil {
			return "", fmt.Errorf("failed to marshal public key into PKIX format")
	}
	ecdhSKPEMBlock := pem.EncodeToMemory(
			&pem.Block{
					Type:  "PUBLIC KEY",
					Bytes: ecdhSKBytes,
			},
	)
	return string(ecdhSKPEMBlock), nil
}

func(k *ECKeys)MarshalECDHPrivateKey()(string,error){
	ecdhSKBytes, err := x509.MarshalPKCS8PrivateKey(k.Private)
    if err != nil {
        return "", fmt.Errorf("failed to marshal private key into PKIX format")
    }

    ecdhSKPEMBlock := pem.EncodeToMemory(
        &pem.Block{
            Type:  "PRIVATE KEY",
            Bytes: ecdhSKBytes,
        },
    )
    return string(ecdhSKPEMBlock), nil
}


