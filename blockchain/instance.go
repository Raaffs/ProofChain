package blockchain

import (
	"log"

	verify "github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractVerifyOperations struct {
	Address  common.Address
	Instance *verify.Verify
	Client   *ethclient.Client
}

func (cv *ContractVerifyOperations) New(contractAddr string, newContract bool) error {
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
	log.Println("Opts : ", opts)
	_, err := cv.Instance.RegisterAsVerifier(opts, name, email, aadhar, institute)
	if err != nil {
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


func (cv *ContractVerifyOperations)GetDocuments(opts *bind.CallOpts)([]verify.VerificationDocument,error){
	docs,err:=cv.Instance.GetDocumentList(opts)
	if err!=nil{
		return nil,err
	}
	return docs,nil
}



