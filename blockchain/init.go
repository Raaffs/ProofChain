package blockchain

//Will implement this in future
import (
	"fmt"
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
	client:=Client()
	c.SetClient(client)
	i.SetClient(client)

	if err:=c.New(privateKey);err!=nil{
		fmt.Println(privateKey)
		fmt.Println("Error connecting to blockchain : ", err)
		return err
	}

	if err:=i.New(contractAddr);err!=nil{
		fmt.Println("Error creating an instance : ", err)
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