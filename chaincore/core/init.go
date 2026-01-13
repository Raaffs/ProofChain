package blockchain

//Will implement this in future
import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)


func Init(c Connect, i Connect, privateKey string, contractAddr string, clientUrl string) error {
    client, err := Client(clientUrl)
    if err != nil {
        return err
    }
    c.SetClient(client)
    i.SetClient(client)
    if err := c.New(privateKey); err != nil {
        return fmt.Errorf("Error initalizing the connection to blockchain : %v", err)
    }

    if err := i.New(contractAddr); err != nil {
        return fmt.Errorf("Error creating an instance : %v", err)
    }
    return nil
}

//todo: turn this into a load balancer
func Client(clientUrl string) (*ethclient.Client, error) {
    client, err := ethclient.Dial(clientUrl)
    if err != nil {
        return nil, fmt.Errorf("Error connecting to the client : %v", err)
    }
    return client, nil
}