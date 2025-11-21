package blockchain

//Will implement this in future
import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Connect interface{
	New(string)error
	SetClient(*ethclient.Client)
}

func Init(c Connect, i Connect, privateKey string, contractAddr string,clientUrl string)error{
	client,err:=Client(clientUrl);if err!=nil{
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

func Client(clientUrl string)(*ethclient.Client,error){
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		log.Println("Error connecting to the client : ", err)
		return nil,fmt.Errorf("Error connecting to client")
	}
	return client,nil
}