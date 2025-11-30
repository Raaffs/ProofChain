package keyUtils

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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
//encrypt the document, so that only user and instituion can view the document

func(k *ECKeys)OnLogin(user ,passphrase,  path string)error{
	pemKey,err:=DecryptPrivateKeyFile(user,passphrase, path);if err!=nil{
		return err
	}
	k.Private,err=GetECDSAPrivateKeyFromPEM(string(pemKey)); if err!=nil{
		return err
	}
	k.Public=k.Private.PublicKey()
	return err
}

func(k *ECKeys)OnRegister(password, keyDir string)(string, string, error){
	privECDSA,err:=ecdsa.GenerateKey(elliptic.P256(),rand.Reader);if err!=nil{
		return "","",err
	}
	k.Private,err=privECDSA.ECDH();if err!=nil{
		return "","",err
	}
	k.Public=k.Private.PublicKey()
	pemPrivateKey,err:=k.MarshalECDHPrivateKey();if err!=nil{
		return "","",err
	}
	pemPublicKey,err:=k.MarshalECDHPublicKey();if err!=nil{
		return  "","",err
	}
	filePath,err:=EncryptPrivateKeyFile(pemPrivateKey,password,keyDir)
	if err!=nil{
		return "","",err
	}
	return pemPublicKey,filePath,nil
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


