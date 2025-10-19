package users

import (
	"github.com/Suy56/ProofChain/chaincore/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)
type Verifier struct{
	Conn 		*blockchain.ClientConnection
	Instance 	*blockchain.ContractVerifyOperations
	Name		string
}

func(v *Verifier)SetName(name string){
	v.Name = name
}

func (v *Verifier)GetName()string{
	return v.Name
}
func (v *Verifier)GetPublicAddress()string{
	return v.Conn.TxOpts.From.Hex()
}
func(v *Verifier)GetClient()*blockchain.ClientConnection{
	return v.Conn
}

func(v *Verifier)GetInstance()*blockchain.ContractVerifyOperations{
	return v.Instance
}


func(v *Verifier)SetTxOpts(c *blockchain.ClientConnection,i *blockchain.ContractVerifyOperations){
	v.Conn=c
	v.Instance=i
}

func(v *Verifier)GetTxOpts()*bind.TransactOpts{
	return v.Conn.TxOpts
}

func(v *Verifier)UpdateTxOpts(opts *bind.TransactOpts)*bind.TransactOpts{
	return v.Conn.TxOpts
}

func(v *Verifier)Register(publicKey string,name string)error{
	v.Instance.RegisterInstitution(v.Conn.TxOpts,publicKey,name)
	return nil
}

func(v *Verifier)GetApprovalStatus()(bool,error){
	return true,nil
}

func (v *Verifier)GetPublicKeys(institute string,user string)(string,error){
	pub,err:=v.Instance.Instance.GetUserPublicKey(v.Conn.CallOpts,common.HexToAddress(user));if err!=nil{
		return "",err
	}
	return pub,nil
}


func(v *Verifier)GetDocuments()([]blockchain.VerificationDocument,error){
	docs,err:=v.Instance.GetDocuments(v.Conn.CallOpts);if err!=nil{
		return nil,err
	}
	return docs,nil
}

func(v *Verifier)GetAcceptedDocuments(docs []blockchain.VerificationDocument)([]blockchain.VerificationDocument){
	accepted:=func(doc blockchain.VerificationDocument)bool{
		return doc.Institute==v.Name && doc.Stats==0
	}
	acceptedDocs:=blockchain.FilterDocument(docs,accepted)
	return acceptedDocs
}

func(v *Verifier)GetRejectedDocuments(docs []blockchain.VerificationDocument)([]blockchain.VerificationDocument){
	rejected:=func(doc blockchain.VerificationDocument)bool{
		return doc.Institute==v.Name && doc.Stats==1
	}
	rejectedDocs:=blockchain.FilterDocument(docs,rejected)
	return rejectedDocs

}

func(v *Verifier)GetPendingDocuments(docs []blockchain.VerificationDocument)([]blockchain.VerificationDocument){
	pending:=func(doc blockchain.VerificationDocument)bool{
		return doc.Institute==v.Name && doc.Stats==2
	}
	pendingDocs:=blockchain.FilterDocument(docs,pending)
	return pendingDocs
}