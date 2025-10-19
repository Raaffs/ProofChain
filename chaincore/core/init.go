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
	client,err:=Client();if err!=nil{
		return err
	}
	c.SetClient(client)
	i.SetClient(client)
	if err:=c.New(privateKey);err!=nil{
		fmt.Println("Error initalizing the connection to blockchain : ", err)
		return err
	}

	if err:=i.New(contractAddr);err!=nil{
		fmt.Println("Error creating an instance : ", err)
		return err
	}
	return nil
}

func Client()(*ethclient.Client,error){
	err:=godotenv.Load()
	if err!=nil{
		log.Println("Error loading .env file : ",err)
		return nil,fmt.Errorf("Error connecting to client")
	}

	client_url:=os.Getenv("CLIENT_URL")
	if client_url==""{
		log.Println("Error retrieving client-url : ",err)
		return nil,fmt.Errorf("Error connecting to client")
	}

	client, err := ethclient.Dial(client_url)
	if err != nil {
		log.Println("Error connecting to the client : ", err)
		return nil,fmt.Errorf("Error connecting to client")
	}
	return client,nil
}