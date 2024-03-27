package blockchain

import (
	"github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var TempContractAddress=""

func(conn *ClientConnection)Deploy()(common.Address,*types.Transaction,error){
	contract,tx,_,err:=verify.DeployVerify(conn.TxOpts,conn.Client)
	if err!=nil{
		return common.Address{},nil,err
	}
	TempContractAddress=contract.Hex()
	return contract,tx,nil
}
