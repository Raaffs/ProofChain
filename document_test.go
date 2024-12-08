package main_test

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/Suy56/ProofChain/keyUtils"
)




func TestDoc(t *testing.T){
	if err:=blockchain.Init(App_test.conn,App_test.in,"84ef8fd680bf64990853c1a044090905cda74903b65b7fe668f281582bb2b434","0x23E4336fB78C884946090B70aa3Db6d1555ade34");err!=nil{
		t.Fatal(err)
	}
	errchan:=make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		App_test.keys.OnLogin("Maria","Maria",errchan)		
	}()
	go func(){
		wg.Wait()
		defer close(errchan)
	}()

	for err:=range errchan{
		if err!=nil{
			log.Println("Error retrieving user's keys  : ",err)
		 	log.Fatal("error retrieving account. Make sure the credentials are correct")
		}
	}
	pub,err:=App_test.in.Instance.GetInstituePublicKey(App_test.conn.CallOpts,"ins");if err!=nil{
		t.Fatal("Error getting public key: ",err)
	}
	log.Println("public key of ins: ",pub)
	if err:=App_test.keys.SetMultiSigKey(pub);err!=nil{
		log.Fatal("Error setting public key : ",err)
	}
	secret,err:=App_test.keys.GenerateSecret();if err!=nil{
		t.Fatal("Error generating secret key: ",err)
	}

	imagePath:="./assest/API.png"
	imageDoc,err:=os.ReadFile(imagePath);if err!=nil{
		t.Fatal("error reading image",err)
	}

	enc,err:=keyUtils.EncryptIPFSHash(secret,imageDoc);if err!=nil{
		t.Fatal("Error encrypting ",err)
	}
	decrypt,err:=keyUtils.DecryptIPFSHash(secret,enc)
	reader := bytes.NewReader(decrypt)
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}
	outFile, err := os.Create("output.jpg")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer outFile.Close()	
	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}

	fmt.Println("Image saved as output.jpg")
}