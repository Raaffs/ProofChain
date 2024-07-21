// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verify

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VerifyMetaData contains all meta data concerning the Verify contract.
var VerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_EncryptedIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"addDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"approveVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocuments\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"requester\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"verifer\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"institute\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"ipfs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"name\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"desc\",\"type\":\"string[]\"},{\"internalType\":\"enumVerification.DocStatus[]\",\"name\":\"stats\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getInstituePublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getUserPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"institutions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"publicAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"isApprovedInstitute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"registerAsUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"registerInstitution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"verifyDocument\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600d5534801561001557600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611eec806100656000396000f3fe6080604052600436106100915760003560e01c8063ba78572511610059578063ba7857251461017e578063c6b5c5cc146101a7578063e74c0342146101d0578063ea6509f81461020d578063ef2d87001461024d57610091565b80630e72ca7a1461009657806311231fe0146100bf578063515ef106146100fc57806390752bb214610125578063ba1f3dc714610141575b600080fd5b3480156100a257600080fd5b506100bd60048036038101906100b891906115fc565b61027e565b005b3480156100cb57600080fd5b506100e660048036038101906100e19190611462565b61058c565b6040516100f39190611ade565b60405180910390f35b34801561010857600080fd5b50610123600480360381019061011e919061148b565b610660565b005b61013f600480360381019061013a919061157d565b61076e565b005b34801561014d57600080fd5b50610168600480360381019061016391906114d0565b61096b565b6040516101759190611ade565b60405180910390f35b34801561018a57600080fd5b506101a560048036038101906101a091906114d0565b610a1e565b005b3480156101b357600080fd5b506101ce60048036038101906101c99190611511565b610ae8565b005b3480156101dc57600080fd5b506101f760048036038101906101f291906114d0565b610be8565b6040516102049190611ac3565b60405180910390f35b34801561021957600080fd5b50610234600480360381019061022f91906114d0565b610c27565b60405161024494939291906119d0565b60405180910390f35b34801561025957600080fd5b50610262610daa565b6040516102759796959493929190611a23565b60405180910390f35b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690506001151581151514610316576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030d90611b20565b60405180910390fd5b600d5460038660405161032991906119b9565b9081526020016040518091039020819055506006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600760009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060098490806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061043d9291906112e3565b506008839080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906104799291906112e3565b50600a829080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906104b59291906112e3565b50600c600290806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff0219169083600281111561052c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550600b8590806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061056c9291906112e3565b50600d600081548092919061058090611d65565b91905055505050505050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000180546105db90611d33565b80601f016020809104026020016040519081016040528092919081815260200182805461060790611d33565b80156106545780601f1061062957610100808354040283529160200191610654565b820191906000526020600020905b81548152906001019060200180831161063757829003601f168201915b50505050509050919050565b604051806020016040528083838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815250600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001908051906020019061070e9291906112e3565b509050506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6001151560018460405161078291906119b9565b908152602001604051809103902060030160009054906101000a900460ff16151514801561081c57503373ffffffffffffffffffffffffffffffffffffffff166001846040516107d291906119b9565b908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b61082557600080fd5b600060038360405161083791906119b9565b908152602001604051809103902054905081600c8281548110610883577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090602091828204019190066101000a81548160ff021916908360028111156108dd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550336007828154811061091d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b606060018260405161097d91906119b9565b9081526020016040518091039020600101805461099990611d33565b80601f01602080910402602001604051908101604052809291908181526020018280546109c590611d33565b8015610a125780601f106109e757610100808354040283529160200191610a12565b820191906000526020600020905b8154815290600101906020018083116109f557829003601f168201915b50505050509050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610aac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa390611b00565b60405180910390fd5b60018082604051610abd91906119b9565b908152602001604051809103902060030160006101000a81548160ff02191690831515021790555050565b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200160001515815250600182604051610b3291906119b9565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610ba39291906112e3565b506040820151816002019080519060200190610bc09291906112e3565b5060608201518160030160006101000a81548160ff0219169083151502179055509050505050565b600060011515600183604051610bfe91906119b9565b908152602001604051809103902060030160009054906101000a900460ff161515149050919050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610c8690611d33565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb290611d33565b8015610cff5780601f10610cd457610100808354040283529160200191610cff565b820191906000526020600020905b815481529060010190602001808311610ce257829003601f168201915b505050505090806002018054610d1490611d33565b80601f0160208091040260200160405190810160405280929190818152602001828054610d4090611d33565b8015610d8d5780601f10610d6257610100808354040283529160200191610d8d565b820191906000526020600020905b815481529060010190602001808311610d7057829003601f168201915b5050505050908060030160009054906101000a900460ff16905084565b6060806060806060806060600660076009600b6008600a600c86805480602002602001604051908101604052809291908181526020018280548015610e4457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610dfa575b5050505050965085805480602002602001604051908101604052809291908181526020018280548015610ecc57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e82575b5050505050955084805480602002602001604051908101604052809291908181526020016000905b82821015610fa0578382906000526020600020018054610f1390611d33565b80601f0160208091040260200160405190810160405280929190818152602001828054610f3f90611d33565b8015610f8c5780601f10610f6157610100808354040283529160200191610f8c565b820191906000526020600020905b815481529060010190602001808311610f6f57829003601f168201915b505050505081526020019060010190610ef4565b50505050945083805480602002602001604051908101604052809291908181526020016000905b82821015611073578382906000526020600020018054610fe690611d33565b80601f016020809104026020016040519081016040528092919081815260200182805461101290611d33565b801561105f5780601f106110345761010080835404028352916020019161105f565b820191906000526020600020905b81548152906001019060200180831161104257829003601f168201915b505050505081526020019060010190610fc7565b50505050935082805480602002602001604051908101604052809291908181526020016000905b828210156111465783829060005260206000200180546110b990611d33565b80601f01602080910402602001604051908101604052809291908181526020018280546110e590611d33565b80156111325780601f1061110757610100808354040283529160200191611132565b820191906000526020600020905b81548152906001019060200180831161111557829003601f168201915b50505050508152602001906001019061109a565b50505050925081805480602002602001604051908101604052809291908181526020016000905b8282101561121957838290600052602060002001805461118c90611d33565b80601f01602080910402602001604051908101604052809291908181526020018280546111b890611d33565b80156112055780601f106111da57610100808354040283529160200191611205565b820191906000526020600020905b8154815290600101906020018083116111e857829003601f168201915b50505050508152602001906001019061116d565b505050509150808054806020026020016040519081016040528092919081815260200182805480156112c557602002820191906000526020600020906000905b82829054906101000a900460ff16600281111561129f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200190600101906020826000010492830192600103820291508084116112595790505b50505050509050965096509650965096509650965090919293949596565b8280546112ef90611d33565b90600052602060002090601f0160209004810192826113115760008555611358565b82601f1061132a57805160ff1916838001178555611358565b82800160010185558215611358579182015b8281111561135757825182559160200191906001019061133c565b5b5090506113659190611369565b5090565b5b8082111561138257600081600090555060010161136a565b5090565b600061139961139484611b71565b611b40565b9050828152602081018484840111156113b157600080fd5b6113bc848285611cf1565b509392505050565b6000813590506113d381611e8f565b92915050565b6000813590506113e881611ea6565b92915050565b60008083601f84011261140057600080fd5b8235905067ffffffffffffffff81111561141957600080fd5b60208301915083600182028301111561143157600080fd5b9250929050565b600082601f83011261144957600080fd5b8135611459848260208601611386565b91505092915050565b60006020828403121561147457600080fd5b6000611482848285016113c4565b91505092915050565b6000806020838503121561149e57600080fd5b600083013567ffffffffffffffff8111156114b857600080fd5b6114c4858286016113ee565b92509250509250929050565b6000602082840312156114e257600080fd5b600082013567ffffffffffffffff8111156114fc57600080fd5b61150884828501611438565b91505092915050565b6000806040838503121561152457600080fd5b600083013567ffffffffffffffff81111561153e57600080fd5b61154a85828601611438565b925050602083013567ffffffffffffffff81111561156757600080fd5b61157385828601611438565b9150509250929050565b60008060006060848603121561159257600080fd5b600084013567ffffffffffffffff8111156115ac57600080fd5b6115b886828701611438565b935050602084013567ffffffffffffffff8111156115d557600080fd5b6115e186828701611438565b92505060406115f2868287016113d9565b9150509250925092565b6000806000806080858703121561161257600080fd5b600085013567ffffffffffffffff81111561162c57600080fd5b61163887828801611438565b945050602085013567ffffffffffffffff81111561165557600080fd5b61166187828801611438565b935050604085013567ffffffffffffffff81111561167e57600080fd5b61168a87828801611438565b925050606085013567ffffffffffffffff8111156116a757600080fd5b6116b387828801611438565b91505092959194509250565b60006116cb8383611703565b60208301905092915050565b60006116e38383611861565b60208301905092915050565b60006116fb8383611870565b905092915050565b61170c81611c84565b82525050565b61171b81611c84565b82525050565b600061172c82611bd1565b6117368185611c24565b935061174183611ba1565b8060005b8381101561177257815161175988826116bf565b975061176483611bfd565b925050600181019050611745565b5085935050505092915050565b600061178a82611bdc565b6117948185611c46565b935061179f83611bb1565b8060005b838110156117d05781516117b788826116d7565b97506117c283611c0a565b9250506001810190506117a3565b5085935050505092915050565b60006117e882611be7565b6117f28185611c35565b93508360208202850161180485611bc1565b8060005b85811015611840578484038952815161182185826116ef565b945061182c83611c17565b925060208a01995050600181019050611808565b50829750879550505050505092915050565b61185b81611c96565b82525050565b61186a81611cdf565b82525050565b600061187b82611bf2565b6118858185611c57565b9350611895818560208601611d00565b61189e81611e6a565b840191505092915050565b60006118b482611bf2565b6118be8185611c68565b93506118ce818560208601611d00565b6118d781611e6a565b840191505092915050565b60006118ed82611bf2565b6118f78185611c79565b9350611907818560208601611d00565b80840191505092915050565b6000611920602183611c68565b91507f4f6e6c792061646d696e2063616e20706572666f6d207468697320616374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611986601883611c68565b91507f726567697374657220666972737420746f2076657269667900000000000000006000830152602082019050919050565b60006119c582846118e2565b915081905092915050565b60006080820190506119e56000830187611712565b81810360208301526119f781866118a9565b90508181036040830152611a0b81856118a9565b9050611a1a6060830184611852565b95945050505050565b600060e0820190508181036000830152611a3d818a611721565b90508181036020830152611a518189611721565b90508181036040830152611a6581886117dd565b90508181036060830152611a7981876117dd565b90508181036080830152611a8d81866117dd565b905081810360a0830152611aa181856117dd565b905081810360c0830152611ab5818461177f565b905098975050505050505050565b6000602082019050611ad86000830184611852565b92915050565b60006020820190508181036000830152611af881846118a9565b905092915050565b60006020820190508181036000830152611b1981611913565b9050919050565b60006020820190508181036000830152611b3981611979565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715611b6757611b66611e3b565b5b8060405250919050565b600067ffffffffffffffff821115611b8c57611b8b611e3b565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611c8f82611cb5565b9050919050565b60008115159050919050565b6000819050611cb082611e7b565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611cea82611ca2565b9050919050565b82818337600083830152505050565b60005b83811015611d1e578082015181840152602081019050611d03565b83811115611d2d576000848401525b50505050565b60006002820490506001821680611d4b57607f821691505b60208210811415611d5f57611d5e611e0c565b5b50919050565b6000611d7082611cd5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611da357611da2611dae565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60038110611e8c57611e8b611ddd565b5b50565b611e9881611c84565b8114611ea357600080fd5b50565b60038110611eb357600080fd5b5056fea2646970667358221220f77bfe467a8d1c5474b616493594a04884cc391d9d7874bb9940f9cc629b90fc64736f6c63430008000033",
}

// VerifyABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyMetaData.ABI instead.
var VerifyABI = VerifyMetaData.ABI

// VerifyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifyMetaData.Bin instead.
var VerifyBin = VerifyMetaData.Bin

// DeployVerify deploys a new Ethereum contract, binding an instance of Verify to it.
func DeployVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verify, error) {
	parsed, err := VerifyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verify{VerifyCaller: VerifyCaller{contract: contract}, VerifyTransactor: VerifyTransactor{contract: contract}, VerifyFilterer: VerifyFilterer{contract: contract}}, nil
}

// Verify is an auto generated Go binding around an Ethereum contract.
type Verify struct {
	VerifyCaller     // Read-only binding to the contract
	VerifyTransactor // Write-only binding to the contract
	VerifyFilterer   // Log filterer for contract events
}

// VerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifySession struct {
	Contract     *Verify           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyCallerSession struct {
	Contract *VerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyTransactorSession struct {
	Contract     *VerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyRaw struct {
	Contract *Verify // Generic contract binding to access the raw methods on
}

// VerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyCallerRaw struct {
	Contract *VerifyCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyTransactorRaw struct {
	Contract *VerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerify creates a new instance of Verify, bound to a specific deployed contract.
func NewVerify(address common.Address, backend bind.ContractBackend) (*Verify, error) {
	contract, err := bindVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verify{VerifyCaller: VerifyCaller{contract: contract}, VerifyTransactor: VerifyTransactor{contract: contract}, VerifyFilterer: VerifyFilterer{contract: contract}}, nil
}

// NewVerifyCaller creates a new read-only instance of Verify, bound to a specific deployed contract.
func NewVerifyCaller(address common.Address, caller bind.ContractCaller) (*VerifyCaller, error) {
	contract, err := bindVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyCaller{contract: contract}, nil
}

// NewVerifyTransactor creates a new write-only instance of Verify, bound to a specific deployed contract.
func NewVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyTransactor, error) {
	contract, err := bindVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyTransactor{contract: contract}, nil
}

// NewVerifyFilterer creates a new log filterer instance of Verify, bound to a specific deployed contract.
func NewVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyFilterer, error) {
	contract, err := bindVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyFilterer{contract: contract}, nil
}

// bindVerify binds a generic wrapper to an already deployed contract.
func bindVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verify *VerifyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verify.Contract.VerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verify *VerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verify.Contract.VerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verify *VerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verify.Contract.VerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verify *VerifyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verify *VerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verify *VerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verify.Contract.contract.Transact(opts, method, params...)
}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] desc, uint8[] stats)
func (_Verify *VerifyCaller) GetDocuments(opts *bind.CallOpts) (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Desc      []string
	Stats     []uint8
}, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getDocuments")

	outstruct := new(struct {
		Requester []common.Address
		Verifer   []common.Address
		Institute []string
		Ipfs      []string
		Name      []string
		Desc      []string
		Stats     []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Verifer = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Institute = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Ipfs = *abi.ConvertType(out[3], new([]string)).(*[]string)
	outstruct.Name = *abi.ConvertType(out[4], new([]string)).(*[]string)
	outstruct.Desc = *abi.ConvertType(out[5], new([]string)).(*[]string)
	outstruct.Stats = *abi.ConvertType(out[6], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] desc, uint8[] stats)
func (_Verify *VerifySession) GetDocuments() (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Desc      []string
	Stats     []uint8
}, error) {
	return _Verify.Contract.GetDocuments(&_Verify.CallOpts)
}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] desc, uint8[] stats)
func (_Verify *VerifyCallerSession) GetDocuments() (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Desc      []string
	Stats     []uint8
}, error) {
	return _Verify.Contract.GetDocuments(&_Verify.CallOpts)
}

// GetInstituePublicKey is a free data retrieval call binding the contract method 0xba1f3dc7.
//
// Solidity: function getInstituePublicKey(string _name) view returns(string pubKey)
func (_Verify *VerifyCaller) GetInstituePublicKey(opts *bind.CallOpts, _name string) (string, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getInstituePublicKey", _name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetInstituePublicKey is a free data retrieval call binding the contract method 0xba1f3dc7.
//
// Solidity: function getInstituePublicKey(string _name) view returns(string pubKey)
func (_Verify *VerifySession) GetInstituePublicKey(_name string) (string, error) {
	return _Verify.Contract.GetInstituePublicKey(&_Verify.CallOpts, _name)
}

// GetInstituePublicKey is a free data retrieval call binding the contract method 0xba1f3dc7.
//
// Solidity: function getInstituePublicKey(string _name) view returns(string pubKey)
func (_Verify *VerifyCallerSession) GetInstituePublicKey(_name string) (string, error) {
	return _Verify.Contract.GetInstituePublicKey(&_Verify.CallOpts, _name)
}

// GetUserPublicKey is a free data retrieval call binding the contract method 0x11231fe0.
//
// Solidity: function getUserPublicKey(address userAddr) view returns(string)
func (_Verify *VerifyCaller) GetUserPublicKey(opts *bind.CallOpts, userAddr common.Address) (string, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getUserPublicKey", userAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserPublicKey is a free data retrieval call binding the contract method 0x11231fe0.
//
// Solidity: function getUserPublicKey(address userAddr) view returns(string)
func (_Verify *VerifySession) GetUserPublicKey(userAddr common.Address) (string, error) {
	return _Verify.Contract.GetUserPublicKey(&_Verify.CallOpts, userAddr)
}

// GetUserPublicKey is a free data retrieval call binding the contract method 0x11231fe0.
//
// Solidity: function getUserPublicKey(address userAddr) view returns(string)
func (_Verify *VerifyCallerSession) GetUserPublicKey(userAddr common.Address) (string, error) {
	return _Verify.Contract.GetUserPublicKey(&_Verify.CallOpts, userAddr)
}

// Institutions is a free data retrieval call binding the contract method 0xea6509f8.
//
// Solidity: function institutions(string ) view returns(address publicAddr, string publicKey, string name, bool approved)
func (_Verify *VerifyCaller) Institutions(opts *bind.CallOpts, arg0 string) (struct {
	PublicAddr common.Address
	PublicKey  string
	Name       string
	Approved   bool
}, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "institutions", arg0)

	outstruct := new(struct {
		PublicAddr common.Address
		PublicKey  string
		Name       string
		Approved   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PublicKey = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Approved = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Institutions is a free data retrieval call binding the contract method 0xea6509f8.
//
// Solidity: function institutions(string ) view returns(address publicAddr, string publicKey, string name, bool approved)
func (_Verify *VerifySession) Institutions(arg0 string) (struct {
	PublicAddr common.Address
	PublicKey  string
	Name       string
	Approved   bool
}, error) {
	return _Verify.Contract.Institutions(&_Verify.CallOpts, arg0)
}

// Institutions is a free data retrieval call binding the contract method 0xea6509f8.
//
// Solidity: function institutions(string ) view returns(address publicAddr, string publicKey, string name, bool approved)
func (_Verify *VerifyCallerSession) Institutions(arg0 string) (struct {
	PublicAddr common.Address
	PublicKey  string
	Name       string
	Approved   bool
}, error) {
	return _Verify.Contract.Institutions(&_Verify.CallOpts, arg0)
}

// IsApprovedInstitute is a free data retrieval call binding the contract method 0xe74c0342.
//
// Solidity: function isApprovedInstitute(string name) view returns(bool)
func (_Verify *VerifyCaller) IsApprovedInstitute(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "isApprovedInstitute", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedInstitute is a free data retrieval call binding the contract method 0xe74c0342.
//
// Solidity: function isApprovedInstitute(string name) view returns(bool)
func (_Verify *VerifySession) IsApprovedInstitute(name string) (bool, error) {
	return _Verify.Contract.IsApprovedInstitute(&_Verify.CallOpts, name)
}

// IsApprovedInstitute is a free data retrieval call binding the contract method 0xe74c0342.
//
// Solidity: function isApprovedInstitute(string name) view returns(bool)
func (_Verify *VerifyCallerSession) IsApprovedInstitute(name string) (bool, error) {
	return _Verify.Contract.IsApprovedInstitute(&_Verify.CallOpts, name)
}

// AddDocument is a paid mutator transaction binding the contract method 0x0e72ca7a.
//
// Solidity: function addDocument(string _EncryptedIPFSHash, string _institute, string _name, string _description) returns()
func (_Verify *VerifyTransactor) AddDocument(opts *bind.TransactOpts, _EncryptedIPFSHash string, _institute string, _name string, _description string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "addDocument", _EncryptedIPFSHash, _institute, _name, _description)
}

// AddDocument is a paid mutator transaction binding the contract method 0x0e72ca7a.
//
// Solidity: function addDocument(string _EncryptedIPFSHash, string _institute, string _name, string _description) returns()
func (_Verify *VerifySession) AddDocument(_EncryptedIPFSHash string, _institute string, _name string, _description string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, _EncryptedIPFSHash, _institute, _name, _description)
}

// AddDocument is a paid mutator transaction binding the contract method 0x0e72ca7a.
//
// Solidity: function addDocument(string _EncryptedIPFSHash, string _institute, string _name, string _description) returns()
func (_Verify *VerifyTransactorSession) AddDocument(_EncryptedIPFSHash string, _institute string, _name string, _description string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, _EncryptedIPFSHash, _institute, _name, _description)
}

// ApproveVerifier is a paid mutator transaction binding the contract method 0xba785725.
//
// Solidity: function approveVerifier(string _name) returns()
func (_Verify *VerifyTransactor) ApproveVerifier(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "approveVerifier", _name)
}

// ApproveVerifier is a paid mutator transaction binding the contract method 0xba785725.
//
// Solidity: function approveVerifier(string _name) returns()
func (_Verify *VerifySession) ApproveVerifier(_name string) (*types.Transaction, error) {
	return _Verify.Contract.ApproveVerifier(&_Verify.TransactOpts, _name)
}

// ApproveVerifier is a paid mutator transaction binding the contract method 0xba785725.
//
// Solidity: function approveVerifier(string _name) returns()
func (_Verify *VerifyTransactorSession) ApproveVerifier(_name string) (*types.Transaction, error) {
	return _Verify.Contract.ApproveVerifier(&_Verify.TransactOpts, _name)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x515ef106.
//
// Solidity: function registerAsUser(string _publicKey) returns()
func (_Verify *VerifyTransactor) RegisterAsUser(opts *bind.TransactOpts, _publicKey string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "registerAsUser", _publicKey)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x515ef106.
//
// Solidity: function registerAsUser(string _publicKey) returns()
func (_Verify *VerifySession) RegisterAsUser(_publicKey string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsUser(&_Verify.TransactOpts, _publicKey)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x515ef106.
//
// Solidity: function registerAsUser(string _publicKey) returns()
func (_Verify *VerifyTransactorSession) RegisterAsUser(_publicKey string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsUser(&_Verify.TransactOpts, _publicKey)
}

// RegisterInstitution is a paid mutator transaction binding the contract method 0xc6b5c5cc.
//
// Solidity: function registerInstitution(string _publicKey, string _name) returns()
func (_Verify *VerifyTransactor) RegisterInstitution(opts *bind.TransactOpts, _publicKey string, _name string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "registerInstitution", _publicKey, _name)
}

// RegisterInstitution is a paid mutator transaction binding the contract method 0xc6b5c5cc.
//
// Solidity: function registerInstitution(string _publicKey, string _name) returns()
func (_Verify *VerifySession) RegisterInstitution(_publicKey string, _name string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterInstitution(&_Verify.TransactOpts, _publicKey, _name)
}

// RegisterInstitution is a paid mutator transaction binding the contract method 0xc6b5c5cc.
//
// Solidity: function registerInstitution(string _publicKey, string _name) returns()
func (_Verify *VerifyTransactorSession) RegisterInstitution(_publicKey string, _name string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterInstitution(&_Verify.TransactOpts, _publicKey, _name)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x90752bb2.
//
// Solidity: function verifyDocument(string _institute, string _ipfs, uint8 _status) payable returns()
func (_Verify *VerifyTransactor) VerifyDocument(opts *bind.TransactOpts, _institute string, _ipfs string, _status uint8) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "verifyDocument", _institute, _ipfs, _status)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x90752bb2.
//
// Solidity: function verifyDocument(string _institute, string _ipfs, uint8 _status) payable returns()
func (_Verify *VerifySession) VerifyDocument(_institute string, _ipfs string, _status uint8) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocument(&_Verify.TransactOpts, _institute, _ipfs, _status)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x90752bb2.
//
// Solidity: function verifyDocument(string _institute, string _ipfs, uint8 _status) payable returns()
func (_Verify *VerifyTransactorSession) VerifyDocument(_institute string, _ipfs string, _status uint8) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocument(&_Verify.TransactOpts, _institute, _ipfs, _status)
}
