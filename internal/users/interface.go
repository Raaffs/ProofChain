package users

import (
	"log"
	"math/big"

	"github.com/Suy56/ProofChain/chaincore/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type User interface{
	SetTxOpts(*blockchain.ClientConnection,*blockchain.ContractVerifyOperations)
	GetTxOpts() *bind.TransactOpts
	GetClient()*blockchain.ClientConnection
	GetInstance()*blockchain.ContractVerifyOperations
	Register(string,string)error
	GetApprovalStatus()(bool,error)
	GetPublicKeys(entity string)(string,error)
	GetDocuments()([]blockchain.VerificationDocument,error)
	GetAcceptedDocuments([]blockchain.VerificationDocument)([]blockchain.VerificationDocument)
	GetRejectedDocuments([]blockchain.VerificationDocument)([]blockchain.VerificationDocument)
	GetPendingDocuments([]blockchain.VerificationDocument)([]blockchain.VerificationDocument)
	SetName(string)
	GetName()string
	GetPublicAddress()string
	AddDocument(	
		hash, 
		institute string,
		add func ()error,
		approve func()error,
	)error
}

func UpdateNonce(u User)error{
	nonce,err:=u.GetClient().Client.PendingNonceAt(u.GetTxOpts().Context,u.GetTxOpts().From);if err!=nil{
		log.Println("Error getting nonce :",err)
		return err
	}
	u.GetTxOpts().Nonce=big.NewInt(int64(nonce))
	return nil
}