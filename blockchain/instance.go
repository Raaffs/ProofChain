package blockchain

import (
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

func(cv *ContractVerifyOperations)getDocuments(opts *bind.CallOpts){
	cv.Instance.GetDocumentList(opts)
}