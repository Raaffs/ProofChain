package users

import (
	"log"

	"github.com/Suy56/ProofChain/blockchain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)
type Requester struct{
	Conn 		*blockchain.ClientConnection
	Instance 	*blockchain.ContractVerifyOperations
}

func(r *Requester)SetTxOpts(c *blockchain.ClientConnection, i *blockchain.ContractVerifyOperations){
	r.Conn=c
	r.Instance=i
}

func(r *Requester)GetTxOpts()*bind.TransactOpts{
	return r.Conn.TxOpts
}

func(r *Requester)UpdateTxOpts(opts *bind.TransactOpts)*bind.TransactOpts{
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
	acceptedDocs:=blockchain.FilterDocument(docs,accepted)
	return acceptedDocs
}

func(r *Requester)GetRejectedDocuments(docs []blockchain.VerificationDocument)([]blockchain.VerificationDocument){
	rejected:=func(doc blockchain.VerificationDocument)bool{
		return doc.Requester==r.Conn.CallOpts.From.Hex() && doc.Stats==1
	}
	rejectedDocs:=blockchain.FilterDocument(docs,rejected)
	return rejectedDocs
}

func(r *Requester)GetPendingDocuments(docs []blockchain.VerificationDocument)([]blockchain.VerificationDocument){
	pending:=func(doc blockchain.VerificationDocument)bool{
		return doc.Requester==r.Conn.CallOpts.From.Hex() && doc.Stats==2
	}
	pendingDocs:=blockchain.FilterDocument(docs,pending)
	return pendingDocs
}
