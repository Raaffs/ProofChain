package users

import (
	"fmt"
	"log"
	"math/big"

	"github.com/Suy56/ProofChain/chaincore/core"
	"github.com/Suy56/ProofChain/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)
type Requester struct{
	Name		string
	Conn 		*blockchain.ClientConnection
	Instance 	*blockchain.ContractVerifyOperations
}
func(r *Requester)SetName(name string){
	r.Name = name
}

func (r *Requester)GetName()string{
	return r.Name
}
func (r *Requester)GetPublicAddress()string{
	return r.Conn.TxOpts.From.Hex()
}
func(r *Requester)UpdateNonce()error{
	nonce,err:=r.Conn.Client.PendingNonceAt(r.GetTxOpts().Context,r.GetTxOpts().From);if err!=nil{
		return fmt.Errorf("Error getting nonce %w",err)
	}
	r.Conn.TxOpts.Nonce=big.NewInt(int64(nonce))
	return nil
}

func(r *Requester)GetClient()*blockchain.ClientConnection{
	return r.Conn
}

func(r *Requester)GetInstance()*blockchain.ContractVerifyOperations{
	return r.Instance
}
func(r *Requester)SetTxOpts(c *blockchain.ClientConnection, i *blockchain.ContractVerifyOperations){
	r.Conn=c
	r.Instance=i
}

func(r *Requester)GetTxOpts()*bind.TransactOpts{
	return r.Conn.TxOpts
}


func(r *Requester)Register(publicKey string,name string)error{
	if err:=r.Instance.RegisterUser(r.Conn.TxOpts,publicKey); err!=nil{
		return err
	}
	return nil
}

func(r *Requester)GetApprovalStatus()(bool,error){
	approved,err:=r.Instance.Instance.IsApprovedInstitute(r.Conn.CallOpts,"");if err!=nil{
		return false,err
	}
	return approved,nil
}

func (r *Requester)GetPublicKeys(institute string,user string)(string,error){
	pub,err:=r.Instance.Instance.GetInstituePublicKey(r.Conn.CallOpts,institute);if err!=nil{
		return "",err
	}
	log.Println("public key: ",pub)
	return pub,nil
}

func(r *Requester)GetDocuments()([]blockchain.VerificationDocument,error){
	docs,err:=r.Instance.GetDocuments(r.Conn.CallOpts);if err!=nil{
		return nil,err
	}
	return docs,nil
}

func(r *Requester)GetAcceptedDocuments(docs []blockchain.VerificationDocument)([]blockchain.VerificationDocument){
	accepted:=func(doc blockchain.VerificationDocument)bool{
		return doc.Requester==r.Conn.CallOpts.From.Hex() && doc.Stats==0
	}
	acceptedDocs:=utils.FilterDocument(docs,accepted)
	return acceptedDocs
}

func(r *Requester)GetRejectedDocuments(docs []blockchain.VerificationDocument)([]blockchain.VerificationDocument){
	rejected:=func(doc blockchain.VerificationDocument)bool{
		return doc.Requester==r.Conn.CallOpts.From.Hex() && doc.Stats==1
	}
	rejectedDocs:=utils.FilterDocument(docs,rejected)
	return rejectedDocs
}

func(r *Requester)GetPendingDocuments(docs []blockchain.VerificationDocument)([]blockchain.VerificationDocument){
	pending:=func(doc blockchain.VerificationDocument)bool{
		return doc.Requester==r.Conn.CallOpts.From.Hex() && doc.Stats==2
	}
	pendingDocs:=utils.FilterDocument(docs,pending)
	return pendingDocs
}

func(r *Requester)AddDocument(hash, institute string)error{
	if err:=r.Instance.AddDocument(r.GetTxOpts(),hash,institute);err!=nil{
		return fmt.Errorf("error adding document: %w",err)
	}
	return nil
}
