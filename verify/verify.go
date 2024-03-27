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

// VerificationDocument is an auto generated low-level Go binding around an user-defined struct.
type VerificationDocument struct {
	Requester        common.Address
	VerifiedBy       common.Address
	Name             string
	Description      string
	DocAddressOnIPFS string
	DocIndex         *big.Int
	Status           uint8
}

// VerifyMetaData contains all meta data concerning the Verify contract.
var VerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_docAddressOnIPFS\",\"type\":\"string\"}],\"name\":\"addDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocumentList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifiedBy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"docAddressOnIPFS\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"docIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structVerification.Document[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"}],\"name\":\"registerAsUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_aadharNum\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"}],\"name\":\"registerAsVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"AadharNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"institute\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVerifier\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_docAddressOnIPFS\",\"type\":\"string\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"verifyDocuments\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600655600060075534801561001a57600080fd5b506121a48061002a6000396000f3fe6080604052600436106100555760003560e01c806320857a9f1461005a57806336a8915b146100835780636c824487146100ac578063d6307acd146100ed578063f036d2e814610116578063fb63030c14610141575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c9190611882565b61015d565b005b34801561008f57600080fd5b506100aa60048036038101906100a591906118f7565b6102d4565b005b3480156100b857600080fd5b506100d360048036038101906100ce9190611859565b610984565b6040516100e4959493929190611dbb565b60405180910390f35b3480156100f957600080fd5b50610114600480360381019061010f91906119ef565b610be7565b005b34801561012257600080fd5b5061012b610cec565b6040516101389190611d99565b60405180910390f35b61015b6004803603810190610156919061199b565b611422565b005b604051806040016040528085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815250600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906102559291906116da565b5060208201518160010190805190602001906102729291906116da565b509050506001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050600115158115151461036c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036390611e2a565b60405180910390fd5b60006040518060e001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200189898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200160065481526020016002808111156104d1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250905080600285856040516104e9929190611d69565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906105a19291906116da565b5060608201518160030190805190602001906105be9291906116da565b5060808201518160040190805190602001906105db9291906116da565b5060a0820151816005015560c08201518160060160006101000a81548160ff02191690836002811115610637577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550905050600581908060018154018082558091505060019003906000526020600020906007020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906107159291906116da565b5060608201518160030190805190602001906107329291906116da565b50608082015181600401908051906020019061074f9291906116da565b5060a0820151816005015560c08201518160060160006101000a81548160ff021916908360028111156107ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055505050600660008154809291906107c59061201d565b9190505550600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906007020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906108dd9291906116da565b5060608201518160030190805190602001906108fa9291906116da565b5060808201518160040190805190602001906109179291906116da565b5060a0820151816005015560c08201518160060160006101000a81548160ff02191690836002811115610973577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555050505050505050505050565b60006020528060005260406000206000915090508060000180546109a790611feb565b80601f01602080910402602001604051908101604052809291908181526020018280546109d390611feb565b8015610a205780601f106109f557610100808354040283529160200191610a20565b820191906000526020600020905b815481529060010190602001808311610a0357829003601f168201915b505050505090806001018054610a3590611feb565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6190611feb565b8015610aae5780601f10610a8357610100808354040283529160200191610aae565b820191906000526020600020905b815481529060010190602001808311610a9157829003601f168201915b505050505090806002018054610ac390611feb565b80601f0160208091040260200160405190810160405280929190818152602001828054610aef90611feb565b8015610b3c5780601f10610b1157610100808354040283529160200191610b3c565b820191906000526020600020905b815481529060010190602001808311610b1f57829003601f168201915b505050505090806003018054610b5190611feb565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7d90611feb565b8015610bca5780601f10610b9f57610100808354040283529160200191610bca565b820191906000526020600020905b815481529060010190602001808311610bad57829003601f168201915b5050505050908060040160009054906101000a900460ff16905085565b6040518060a00160405280858152602001848152602001838152602001828152602001600115158152506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190610c6b9291906116da565b506020820151816001019080519060200190610c889291906116da565b506040820151816002019080519060200190610ca59291906116da565b506060820151816003019080519060200190610cc29291906116da565b5060808201518160040160006101000a81548160ff02191690831515021790555090505050505050565b60606000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff1615611094576005805480602002602001604051908101604052809291908181526020016000905b8282101561108957838290600052602060002090600702016040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054610e4390611feb565b80601f0160208091040260200160405190810160405280929190818152602001828054610e6f90611feb565b8015610ebc5780601f10610e9157610100808354040283529160200191610ebc565b820191906000526020600020905b815481529060010190602001808311610e9f57829003601f168201915b50505050508152602001600382018054610ed590611feb565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0190611feb565b8015610f4e5780601f10610f2357610100808354040283529160200191610f4e565b820191906000526020600020905b815481529060010190602001808311610f3157829003601f168201915b50505050508152602001600482018054610f6790611feb565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9390611feb565b8015610fe05780601f10610fb557610100808354040283529160200191610fe0565b820191906000526020600020905b815481529060010190602001808311610fc357829003601f168201915b50505050508152602001600582015481526020016006820160009054906101000a900460ff16600281111561103e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115611076577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152505081526020019060010190610d64565b50505050905061141f565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561141857838290600052602060002090600702016040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546111d290611feb565b80601f01602080910402602001604051908101604052809291908181526020018280546111fe90611feb565b801561124b5780601f106112205761010080835404028352916020019161124b565b820191906000526020600020905b81548152906001019060200180831161122e57829003601f168201915b5050505050815260200160038201805461126490611feb565b80601f016020809104026020016040519081016040528092919081815260200182805461129090611feb565b80156112dd5780601f106112b2576101008083540402835291602001916112dd565b820191906000526020600020905b8154815290600101906020018083116112c057829003601f168201915b505050505081526020016004820180546112f690611feb565b80601f016020809104026020016040519081016040528092919081815260200182805461132290611feb565b801561136f5780601f106113445761010080835404028352916020019161136f565b820191906000526020600020905b81548152906001019060200180831161135257829003601f168201915b50505050508152602001600582015481526020016006820160009054906101000a900460ff1660028111156113cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115611405577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525050815260200190600101906110f3565b5050505090505b90565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff16905060011515811515146114bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b390611e4a565b60405180910390fd5b816002846040516114cd9190611d82565b908152602001604051809103902060060160006101000a81548160ff02191690836002811115611526577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055503360028460405161153c9190611d82565b908152602001604051809103902060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600060028460405161159d9190611d82565b908152602001604051809103902060050154905082600582815481106115ec577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906007020160060160006101000a81548160ff02191690836002811115611645577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055503360058281548110611685577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906007020160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b8280546116e690611feb565b90600052602060002090601f016020900481019282611708576000855561174f565b82601f1061172157805160ff191683800117855561174f565b8280016001018555821561174f579182015b8281111561174e578251825591602001919060010190611733565b5b50905061175c9190611760565b5090565b5b80821115611779576000816000905550600101611761565b5090565b600061179061178b84611e9b565b611e6a565b9050828152602081018484840111156117a857600080fd5b6117b3848285611fa9565b509392505050565b6000813590506117ca81612147565b92915050565b6000813590506117df8161215e565b92915050565b60008083601f8401126117f757600080fd5b8235905067ffffffffffffffff81111561181057600080fd5b60208301915083600182028301111561182857600080fd5b9250929050565b600082601f83011261184057600080fd5b813561185084826020860161177d565b91505092915050565b60006020828403121561186b57600080fd5b6000611879848285016117bb565b91505092915050565b6000806000806040858703121561189857600080fd5b600085013567ffffffffffffffff8111156118b257600080fd5b6118be878288016117e5565b9450945050602085013567ffffffffffffffff8111156118dd57600080fd5b6118e9878288016117e5565b925092505092959194509250565b6000806000806000806060878903121561191057600080fd5b600087013567ffffffffffffffff81111561192a57600080fd5b61193689828a016117e5565b9650965050602087013567ffffffffffffffff81111561195557600080fd5b61196189828a016117e5565b9450945050604087013567ffffffffffffffff81111561198057600080fd5b61198c89828a016117e5565b92509250509295509295509295565b600080604083850312156119ae57600080fd5b600083013567ffffffffffffffff8111156119c857600080fd5b6119d48582860161182f565b92505060206119e5858286016117d0565b9150509250929050565b60008060008060808587031215611a0557600080fd5b600085013567ffffffffffffffff811115611a1f57600080fd5b611a2b8782880161182f565b945050602085013567ffffffffffffffff811115611a4857600080fd5b611a548782880161182f565b935050604085013567ffffffffffffffff811115611a7157600080fd5b611a7d8782880161182f565b925050606085013567ffffffffffffffff811115611a9a57600080fd5b611aa68782880161182f565b91505092959194509250565b6000611abe8383611cb0565b905092915050565b611acf81611f3c565b82525050565b6000611ae082611edb565b611aea8185611efe565b935083602082028501611afc85611ecb565b8060005b85811015611b385784840389528151611b198582611ab2565b9450611b2483611ef1565b925060208a01995050600181019050611b00565b50829750879550505050505092915050565b611b5381611f4e565b82525050565b611b6281611f97565b82525050565b6000611b748385611f31565b9350611b81838584611fa9565b82840190509392505050565b6000611b9882611ee6565b611ba28185611f0f565b9350611bb2818560208601611fb8565b611bbb81612122565b840191505092915050565b6000611bd182611ee6565b611bdb8185611f20565b9350611beb818560208601611fb8565b611bf481612122565b840191505092915050565b6000611c0a82611ee6565b611c148185611f31565b9350611c24818560208601611fb8565b80840191505092915050565b6000611c3d601883611f20565b91507f726567697374657220666972737420746f2076657269667900000000000000006000830152602082019050919050565b6000611c7d601383611f20565b91507f596f75277265206e6f74207665726966696572000000000000000000000000006000830152602082019050919050565b600060e083016000830151611cc86000860182611ac6565b506020830151611cdb6020860182611ac6565b5060408301518482036040860152611cf38282611b8d565b91505060608301518482036060860152611d0d8282611b8d565b91505060808301518482036080860152611d278282611b8d565b91505060a0830151611d3c60a0860182611d5a565b5060c0830151611d4f60c0860182611b59565b508091505092915050565b611d6381611f8d565b82525050565b6000611d76828486611b68565b91508190509392505050565b6000611d8e8284611bff565b915081905092915050565b60006020820190508181036000830152611db38184611ad5565b905092915050565b600060a0820190508181036000830152611dd58188611bc6565b90508181036020830152611de98187611bc6565b90508181036040830152611dfd8186611bc6565b90508181036060830152611e118185611bc6565b9050611e206080830184611b4a565b9695505050505050565b60006020820190508181036000830152611e4381611c30565b9050919050565b60006020820190508181036000830152611e6381611c70565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715611e9157611e906120f3565b5b8060405250919050565b600067ffffffffffffffff821115611eb657611eb56120f3565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611f4782611f6d565b9050919050565b60008115159050919050565b6000819050611f6882612133565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611fa282611f5a565b9050919050565b82818337600083830152505050565b60005b83811015611fd6578082015181840152602081019050611fbb565b83811115611fe5576000848401525b50505050565b6000600282049050600182168061200357607f821691505b60208210811415612017576120166120c4565b5b50919050565b600061202882611f8d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561205b5761205a612066565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6003811061214457612143612095565b5b50565b61215081611f3c565b811461215b57600080fd5b50565b6003811061216b57600080fd5b5056fea2646970667358221220b3a78b29ba500d41d2a5e334f92d72e47729304b4c3f578b58b0a48dd686db7264736f6c63430008000033",
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

// GetDocumentList is a free data retrieval call binding the contract method 0xf036d2e8.
//
// Solidity: function getDocumentList() view returns((address,address,string,string,string,uint256,uint8)[])
func (_Verify *VerifyCaller) GetDocumentList(opts *bind.CallOpts) ([]VerificationDocument, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getDocumentList")

	if err != nil {
		return *new([]VerificationDocument), err
	}

	out0 := *abi.ConvertType(out[0], new([]VerificationDocument)).(*[]VerificationDocument)

	return out0, err

}

// GetDocumentList is a free data retrieval call binding the contract method 0xf036d2e8.
//
// Solidity: function getDocumentList() view returns((address,address,string,string,string,uint256,uint8)[])
func (_Verify *VerifySession) GetDocumentList() ([]VerificationDocument, error) {
	return _Verify.Contract.GetDocumentList(&_Verify.CallOpts)
}

// GetDocumentList is a free data retrieval call binding the contract method 0xf036d2e8.
//
// Solidity: function getDocumentList() view returns((address,address,string,string,string,uint256,uint8)[])
func (_Verify *VerifyCallerSession) GetDocumentList() ([]VerificationDocument, error) {
	return _Verify.Contract.GetDocumentList(&_Verify.CallOpts)
}

// Verifiers is a free data retrieval call binding the contract method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(string name, string email, string AadharNumber, string institute, bool isVerifier)
func (_Verify *VerifyCaller) Verifiers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name         string
	Email        string
	AadharNumber string
	Institute    string
	IsVerifier   bool
}, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "verifiers", arg0)

	outstruct := new(struct {
		Name         string
		Email        string
		AadharNumber string
		Institute    string
		IsVerifier   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AadharNumber = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Institute = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsVerifier = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Verifiers is a free data retrieval call binding the contract method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(string name, string email, string AadharNumber, string institute, bool isVerifier)
func (_Verify *VerifySession) Verifiers(arg0 common.Address) (struct {
	Name         string
	Email        string
	AadharNumber string
	Institute    string
	IsVerifier   bool
}, error) {
	return _Verify.Contract.Verifiers(&_Verify.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(string name, string email, string AadharNumber, string institute, bool isVerifier)
func (_Verify *VerifyCallerSession) Verifiers(arg0 common.Address) (struct {
	Name         string
	Email        string
	AadharNumber string
	Institute    string
	IsVerifier   bool
}, error) {
	return _Verify.Contract.Verifiers(&_Verify.CallOpts, arg0)
}

// AddDocument is a paid mutator transaction binding the contract method 0x36a8915b.
//
// Solidity: function addDocument(string _name, string _description, string _docAddressOnIPFS) returns()
func (_Verify *VerifyTransactor) AddDocument(opts *bind.TransactOpts, _name string, _description string, _docAddressOnIPFS string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "addDocument", _name, _description, _docAddressOnIPFS)
}

// AddDocument is a paid mutator transaction binding the contract method 0x36a8915b.
//
// Solidity: function addDocument(string _name, string _description, string _docAddressOnIPFS) returns()
func (_Verify *VerifySession) AddDocument(_name string, _description string, _docAddressOnIPFS string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, _name, _description, _docAddressOnIPFS)
}

// AddDocument is a paid mutator transaction binding the contract method 0x36a8915b.
//
// Solidity: function addDocument(string _name, string _description, string _docAddressOnIPFS) returns()
func (_Verify *VerifyTransactorSession) AddDocument(_name string, _description string, _docAddressOnIPFS string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, _name, _description, _docAddressOnIPFS)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x20857a9f.
//
// Solidity: function registerAsUser(string _name, string _email) returns()
func (_Verify *VerifyTransactor) RegisterAsUser(opts *bind.TransactOpts, _name string, _email string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "registerAsUser", _name, _email)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x20857a9f.
//
// Solidity: function registerAsUser(string _name, string _email) returns()
func (_Verify *VerifySession) RegisterAsUser(_name string, _email string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsUser(&_Verify.TransactOpts, _name, _email)
}

// RegisterAsUser is a paid mutator transaction binding the contract method 0x20857a9f.
//
// Solidity: function registerAsUser(string _name, string _email) returns()
func (_Verify *VerifyTransactorSession) RegisterAsUser(_name string, _email string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsUser(&_Verify.TransactOpts, _name, _email)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0xd6307acd.
//
// Solidity: function registerAsVerifier(string _name, string _email, string _aadharNum, string _institute) returns()
func (_Verify *VerifyTransactor) RegisterAsVerifier(opts *bind.TransactOpts, _name string, _email string, _aadharNum string, _institute string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "registerAsVerifier", _name, _email, _aadharNum, _institute)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0xd6307acd.
//
// Solidity: function registerAsVerifier(string _name, string _email, string _aadharNum, string _institute) returns()
func (_Verify *VerifySession) RegisterAsVerifier(_name string, _email string, _aadharNum string, _institute string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsVerifier(&_Verify.TransactOpts, _name, _email, _aadharNum, _institute)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0xd6307acd.
//
// Solidity: function registerAsVerifier(string _name, string _email, string _aadharNum, string _institute) returns()
func (_Verify *VerifyTransactorSession) RegisterAsVerifier(_name string, _email string, _aadharNum string, _institute string) (*types.Transaction, error) {
	return _Verify.Contract.RegisterAsVerifier(&_Verify.TransactOpts, _name, _email, _aadharNum, _institute)
}

// VerifyDocuments is a paid mutator transaction binding the contract method 0xfb63030c.
//
// Solidity: function verifyDocuments(string _docAddressOnIPFS, uint8 _status) payable returns()
func (_Verify *VerifyTransactor) VerifyDocuments(opts *bind.TransactOpts, _docAddressOnIPFS string, _status uint8) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "verifyDocuments", _docAddressOnIPFS, _status)
}

// VerifyDocuments is a paid mutator transaction binding the contract method 0xfb63030c.
//
// Solidity: function verifyDocuments(string _docAddressOnIPFS, uint8 _status) payable returns()
func (_Verify *VerifySession) VerifyDocuments(_docAddressOnIPFS string, _status uint8) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocuments(&_Verify.TransactOpts, _docAddressOnIPFS, _status)
}

// VerifyDocuments is a paid mutator transaction binding the contract method 0xfb63030c.
//
// Solidity: function verifyDocuments(string _docAddressOnIPFS, uint8 _status) payable returns()
func (_Verify *VerifyTransactorSession) VerifyDocuments(_docAddressOnIPFS string, _status uint8) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocuments(&_Verify.TransactOpts, _docAddressOnIPFS, _status)
}