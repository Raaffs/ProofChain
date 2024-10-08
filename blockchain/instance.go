package blockchain

import (

	verify "github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//GetDocument method in Verify.go returns an anonymous struct with corresponding
//fields.
type VerificationDocument struct{
	//id field is required in frontend for data row element
	ID			int
	Requester   string
	Verifier    string
	Institute	string
	Name        string
	Desc        string
	IpfsAddress string
	ShaHash		string
	Stats       uint8
}

type ContractVerifyOperations struct {
	Address  common.Address
	Instance *verify.Verify
	Client   *ethclient.Client
}
func (cv *ContractVerifyOperations)SetClient(c *ethclient.Client){
	cv.Client=c
}
func (cv *ContractVerifyOperations)New(contractAddr string) error {
	cv.Address = common.HexToAddress(contractAddr)
	instance, err := verify.NewVerify(cv.Address, cv.Client)
	if err != nil {
		return err
	}
	cv.Instance = instance
	return nil
}

func (cv *ContractVerifyOperations) RegisterUser(opts *bind.TransactOpts,publicKey string) error {
	_, err := cv.Instance.RegisterAsUser(opts,publicKey)

	if err != nil {
		return err
	}
	return nil
}
func (cv *ContractVerifyOperations)RegisterInstitution(opts *bind.TransactOpts, publicKey, institute string) error {
	_, err := cv.Instance.RegisterInstitution(opts,publicKey,institute )
	if err != nil {
		return err
	}
	return nil
}

func (cv *ContractVerifyOperations)ApproveVerifier(opts *bind.TransactOpts,_institute string)error{
	_,err:=cv.Instance.ApproveVerifier(opts, _institute)
	if err!=nil{
		return err
	}
	return nil
}

func (cv *ContractVerifyOperations) AddDocument(opts *bind.TransactOpts,shaHash , _encryptedIPFSHash , _institute, _name string) (error) {
	_, err := cv.Instance.AddDocument(opts, (shaHash),_encryptedIPFSHash,_institute,_name)
	if err != nil {
		return err
	}
	return err
}

func (cv *ContractVerifyOperations)VerifyDocument(opts *bind.TransactOpts, shaHash string, institute string,_status uint8) error {
	_, err := cv.Instance.VerifyDocument(opts, shaHash, institute, _status)
	if err != nil {
		return err
	}
	return nil
}

func (cv *ContractVerifyOperations)GetDocuments(opts *bind.CallOpts)([]VerificationDocument,error){
	var userDocs []VerificationDocument

	docs,err:=cv.Instance.GetDocuments(opts)
	if err!=nil{
		return nil,err
	}
	
	for i:=0;i<len(docs.Requester);i++{
		userDoc:=VerificationDocument{
			//ID field is required in frontend for data row element
			ID: 			i,				
			Requester: 		docs.Requester[i].Hex(),
			Verifier: 		docs.Verifer[i].Hex(),
			Institute: 		docs.Institute[i],
			Name: 			docs.Name[i],
			IpfsAddress:	docs.Ipfs[i],
			ShaHash:        docs.Hash[i],
			Stats: 			docs.Stats[i],
		}
		userDocs = append(userDocs,userDoc)
	}
	return userDocs,nil
}



