package blockchain

import (
	
	"log"

	"github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//todo: write the contract address in env after deployment
var TempContractAddress=""

func Deploy(txOpts *bind.TransactOpts, client *ethclient.Client)(string,*types.Transaction,error){
	contractAddr,tx,_,err:=verify.DeployVerify(txOpts,client)
	if err!=nil{
		log.Print("Error deploying contract")
		return "",nil,err
	}
	return contractAddr.Hex(),tx,err
}

