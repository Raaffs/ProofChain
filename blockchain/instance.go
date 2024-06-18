package blockchain

import (
	"fmt"

	verify "github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mr-tron/base58"
	multihash "github.com/multiformats/go-multihash"
)

type VerificationDocument struct{
	Requester   common.Address
	Verifer     common.Address
	Name        string
	Desc        string
	IpfsAddress string
	Stats       uint8
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
func (cv *ContractVerifyOperations) RegiserVerifier(opts *bind.TransactOpts, publicKey, institute string) error {
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

func (cv *ContractVerifyOperations) AddDocument(opts *bind.TransactOpts, _encryptedIPFSHash, _institute, _name, _description string) (error) {
	institutePubKey,err:=cv.Instance.GetInstituePublicKey(&bind.CallOpts{From: opts.From},_institute)
	if err!=nil{
		return err
	}
	fmt.Println(institutePubKey)
	_, err = cv.Instance.AddDocument(opts, _encryptedIPFSHash,_institute,_name,_description)
	if err != nil {
		return err
	}
	fmt.Println("data : ",_encryptedIPFSHash)
	return err
}

func (cv *ContractVerifyOperations)VerifyDocument(opts *bind.TransactOpts,institute string, _docAddressOnIPFS string, _status uint8) error {
	_, err := cv.Instance.VerifyDocument(opts, institute, _docAddressOnIPFS, _status)
	if err != nil {
		return err
	}
	return nil
}



func (cv *ContractVerifyOperations)GetDocuments(opts *bind.CallOpts)([]VerificationDocument,error){
	var userDocs []VerificationDocument
	docs,err:=cv.Instance.GetDocuments(opts)
	fmt.Println("docs ipfs : ",docs.Ipfs)
	if err!=nil{
		return nil,err
	}
	for i:=0;i<len(docs.Requester);i++{

		var ipfsAddress string
		
		if len(docs.Ipfs) > 0 {
			ipfsAddress = docs.Ipfs[i]
			fmt.Println("IPFS: " ,ipfsAddress)
		}

		userDoc:=VerificationDocument{
			Requester: docs.Requester[i],
			Verifer: docs.Verifer[i],
			Name: docs.Name[i],
			Desc: docs.Desc[i],
			IpfsAddress: ipfsAddress,
			Stats: docs.Stats[i],
		}
		userDocs = append(userDocs,userDoc )
	}
	fmt.Println("user doc: ",userDocs)
	fmt.Println("doc : ",docs)
	return userDocs,nil
}



func IpfsHashTo32Byte(ipfsHash string) ([32]byte, error) {
	// Decode the base58 IPFS hash
	decoded, err := base58.Decode(ipfsHash)
	if err != nil {
		return [32]byte{}, err
	}

	// Decode the multihash
	mhash, err := multihash.Decode(decoded)
	if err != nil {
		return [32]byte{}, err
	}

	// Get the digest
	digest := mhash.Digest

	// Convert to [32]byte
	var byteArray [32]byte
	copy(byteArray[:], digest)

	return byteArray, nil
}

