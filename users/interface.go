package users
import (
	"github.com/Suy56/ProofChain/blockchain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)
type User interface{
	SetTxOpts(*blockchain.ClientConnection,*blockchain.ContractVerifyOperations)
	GetTxOpts() *bind.TransactOpts
	UpdateTxOpts(*bind.TransactOpts)*bind.TransactOpts
	Register(string,string)error
	GetApprovalStatus()(bool,error)
	GetPublicKeys(string,string)(string,error)
	GetDocuments()([]blockchain.VerificationDocument,error)
	GetAcceptedDocuments([]blockchain.VerificationDocument)([]blockchain.VerificationDocument)
	GetRejectedDocuments([]blockchain.VerificationDocument)([]blockchain.VerificationDocument)
	GetPendingDocuments([]blockchain.VerificationDocument)([]blockchain.VerificationDocument)
}