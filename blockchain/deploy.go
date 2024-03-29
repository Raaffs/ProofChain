package blockchain

import (
	"github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
//todo: write the contract address in env after deployment
var TempContractAddress="0xEB941ea2F009d046a14261CcF7be147009F3A6B0"

func(conn *ClientConnection)Deploy()(common.Address,*types.Transaction,error){
	contract,tx,_,err:=verify.DeployVerify(conn.TxOpts,conn.Client)
	if err!=nil{
		return common.Address{},nil,err
	}
	TempContractAddress=contract.Hex()
	return contract,tx,nil
}
