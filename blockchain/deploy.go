package blockchain

import (
	"github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
//todo: write the contract address in env after deployment
var TempContractAddress="0x3b6180deae925647c0ce07cb92328b8635918212739989a2348bfa8a5458168d"

func(conn *ClientConnection)Deploy()(common.Address,*types.Transaction,error){
	contract,tx,_,err:=verify.DeployVerify(conn.TxOpts,conn.Client)
	if err!=nil{
		return common.Address{},nil,err
	}
	TempContractAddress=contract.Hex()
	return contract,tx,nil
}
