package blockchain

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Connect interface{
	New(string)error
	SetClient(*ethclient.Client)
}

func Init(c Connect, i Connect, privateKey string, contractAddr string)error{

	c.SetClient(Client())
	i.SetClient(Client())

	if err:=c.New(privateKey);err!=nil{
		return err
	}

	if err:=i.New(contractAddr);err!=nil{
		return err
	}
	return nil
}

func Client()*ethclient.Client{
	err:=godotenv.Load()
	if err!=nil{
		panic("Error loading env")
	}

	client_url:=os.Getenv("CLIENT_URL")
	if client_url==""{
		panic("INVALID CLIENT URL")
	}


	client, err := ethclient.Dial(client_url)
	if err != nil {
		log.Fatal("Error connecting to the client : ", err)
	}
	return client
}