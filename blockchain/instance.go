package blockchain

import (
	"math/big"

	verify "github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VerificationDocument struct{
	Requester   common.Address
	Verifer     common.Address
	Name        string
	Desc        string
	IpfsAddress string
	Stats       uint8
	UserDocId   *big.Int
}

type ContractVerifyOperations struct {
	Address  common.Address
	Instance *verify.Verify
	Client   *ethclient.Client
}

func (cv *ContractVerifyOperations) New(contractAddr string) error {
	cv.Address = common.HexToAddress(contractAddr)
	instance, err := verify.NewVerify(cv.Address, cv.Client)
	if err != nil {
		return err
	}
	cv.Instance = instance
	return nil
}

func (cv *ContractVerifyOperations) RegisterUser(opts *bind.TransactOpts, name string, email string) error {
	_, err := cv.Instance.RegisterAsUser(opts, name, email)

	if err != nil {
		return err
	}
	return nil
}
func (cv *ContractVerifyOperations) RegiserVerifier(opts *bind.TransactOpts, name string, email string, aadhar string, institute string) error {
	_, err := cv.Instance.RegisterAsVerifier(opts, name, email, aadhar, institute)
	if err != nil {
		return err
	}
	return nil
}

func (cv *ContractVerifyOperations)ApproveVerifier(opts *bind.TransactOpts,address common.Address)error{
	_,err:=cv.Instance.ApproveVerifier(opts, address)
	if err!=nil{
		return err
	}
	return nil
}

func (cv *ContractVerifyOperations) AddDocument(opts *bind.TransactOpts, _name string, _description string, _docAddressOnIPFS string) error {
	_, err := cv.Instance.AddDocument(opts, _name, _description, _docAddressOnIPFS)
	if err != nil {
		return err
	}
	return nil
}

func (cv *ContractVerifyOperations) VerifyDocument(opts *bind.TransactOpts, _docAddressOnIPFS string, _status uint8) error {
	_, err := cv.Instance.VerifyDocuments(opts, _docAddressOnIPFS, _status)
	if err != nil {
		return err
	}
	return nil
}


func (cv *ContractVerifyOperations)GetDocuments(opts *bind.CallOpts)([]VerificationDocument,error){
	var userDocs []VerificationDocument
	docs,err:=cv.Instance.GetDocumentList(opts)
	if err!=nil{
		return nil,err
	}
	for i:=0;i<len(docs.Requester);i++{
		userDoc:=VerificationDocument{
			Requester: docs.Requester[i],
			Verifer: docs.Verifer[i],
			Name: docs.Name[i],
			Desc: docs.Desc[i],
			IpfsAddress: docs.IpfsAddress[i],
			Stats: docs.Stats[i],
			UserDocId: docs.UserDocId[i],
		}
		userDocs = append(userDocs,userDoc )
	}
	return userDocs,nil
}