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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shaHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_EncryptedIPFSHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"approveVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocIndexCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shaHash\",\"type\":\"string\"}],\"name\":\"getDocumentIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocuments\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"requester\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"verifer\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"institute\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"ipfs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"name\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"hash\",\"type\":\"string[]\"},{\"internalType\":\"enumVerification.DocStatus[]\",\"name\":\"stats\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getInstituePublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getUserPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"institutions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"publicAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"isApprovedInstitute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"registerAsUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"registerInstitution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shaHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"verifyDocument\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600d5534801561001557600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611fc6806100656000396000f3fe6080604052600436106100a75760003560e01c8063ba1f3dc711610064578063ba1f3dc7146101bf578063ba785725146101fc578063c6b5c5cc14610225578063e74c03421461024e578063ea6509f81461028b578063ef2d8700146102cb576100a7565b80630e72ca7a146100ac57806311231fe0146100d5578063515ef106146101125780638e7616641461013b57806390752bb214610178578063b8b4da4c14610194575b600080fd5b3480156100b857600080fd5b506100d360048036038101906100ce91906116ac565b6102fc565b005b3480156100e157600080fd5b506100fc60048036038101906100f79190611512565b61060a565b6040516101099190611b9d565b60405180910390f35b34801561011e57600080fd5b506101396004803603810190610134919061153b565b6106de565b005b34801561014757600080fd5b50610162600480360381019061015d9190611580565b6107ec565b60405161016f9190611bff565b60405180910390f35b610192600480360381019061018d919061162d565b610814565b005b3480156101a057600080fd5b506101a9610a11565b6040516101b69190611bff565b60405180910390f35b3480156101cb57600080fd5b506101e660048036038101906101e19190611580565b610a1b565b6040516101f39190611b9d565b60405180910390f35b34801561020857600080fd5b50610223600480360381019061021e9190611580565b610ace565b005b34801561023157600080fd5b5061024c600480360381019061024791906115c1565b610b98565b005b34801561025a57600080fd5b5061027560048036038101906102709190611580565b610c98565b6040516102829190611b82565b60405180910390f35b34801561029757600080fd5b506102b260048036038101906102ad9190611580565b610cd7565b6040516102c29493929190611a8f565b60405180910390f35b3480156102d757600080fd5b506102e0610e5a565b6040516102f39796959493929190611ae2565b60405180910390f35b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690506001151581151514610394576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038b90611bdf565b60405180910390fd5b600d546003866040516103a79190611a78565b9081526020016040518091039020819055506005339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506009839080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906104bb929190611393565b506007829080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906104f7929190611393565b50600c600290806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff0219169083600281111561056e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550600b849080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906105ae929190611393565b506008859080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906105ea929190611393565b50600d60008154809291906105fe90611e3f565b91905055505050505050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001805461065990611e0d565b80601f016020809104026020016040519081016040528092919081815260200182805461068590611e0d565b80156106d25780601f106106a7576101008083540402835291602001916106d2565b820191906000526020600020905b8154815290600101906020018083116106b557829003601f168201915b50505050509050919050565b604051806020016040528083838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815250600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001908051906020019061078c929190611393565b509050506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60006003826040516107fe9190611a78565b9081526020016040518091039020549050919050565b600115156001836040516108289190611a78565b908152602001604051809103902060030160009054906101000a900460ff1615151480156108c257503373ffffffffffffffffffffffffffffffffffffffff166001836040516108789190611a78565b908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b6108cb57600080fd5b60006003846040516108dd9190611a78565b908152602001604051809103902054905081600c8281548110610929577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090602091828204019190066101000a81548160ff02191690836002811115610983577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555033600682815481106109c3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6000600d54905090565b6060600182604051610a2d9190611a78565b90815260200160405180910390206001018054610a4990611e0d565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7590611e0d565b8015610ac25780601f10610a9757610100808354040283529160200191610ac2565b820191906000526020600020905b815481529060010190602001808311610aa557829003601f168201915b50505050509050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5390611bbf565b60405180910390fd5b60018082604051610b6d9190611a78565b908152602001604051809103902060030160006101000a81548160ff02191690831515021790555050565b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200160001515815250600182604051610be29190611a78565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610c53929190611393565b506040820151816002019080519060200190610c70929190611393565b5060608201518160030160006101000a81548160ff0219169083151502179055509050505050565b600060011515600183604051610cae9190611a78565b908152602001604051809103902060030160009054906101000a900460ff161515149050919050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610d3690611e0d565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6290611e0d565b8015610daf5780601f10610d8457610100808354040283529160200191610daf565b820191906000526020600020905b815481529060010190602001808311610d9257829003601f168201915b505050505090806002018054610dc490611e0d565b80601f0160208091040260200160405190810160405280929190818152602001828054610df090611e0d565b8015610e3d5780601f10610e1257610100808354040283529160200191610e3d565b820191906000526020600020905b815481529060010190602001808311610e2057829003601f168201915b5050505050908060030160009054906101000a900460ff16905084565b6060806060806060806060600560066009600b60076008600c86805480602002602001604051908101604052809291908181526020018280548015610ef457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610eaa575b5050505050965085805480602002602001604051908101604052809291908181526020018280548015610f7c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610f32575b5050505050955084805480602002602001604051908101604052809291908181526020016000905b82821015611050578382906000526020600020018054610fc390611e0d565b80601f0160208091040260200160405190810160405280929190818152602001828054610fef90611e0d565b801561103c5780601f106110115761010080835404028352916020019161103c565b820191906000526020600020905b81548152906001019060200180831161101f57829003601f168201915b505050505081526020019060010190610fa4565b50505050945083805480602002602001604051908101604052809291908181526020016000905b8282101561112357838290600052602060002001805461109690611e0d565b80601f01602080910402602001604051908101604052809291908181526020018280546110c290611e0d565b801561110f5780601f106110e45761010080835404028352916020019161110f565b820191906000526020600020905b8154815290600101906020018083116110f257829003601f168201915b505050505081526020019060010190611077565b50505050935082805480602002602001604051908101604052809291908181526020016000905b828210156111f657838290600052602060002001805461116990611e0d565b80601f016020809104026020016040519081016040528092919081815260200182805461119590611e0d565b80156111e25780601f106111b7576101008083540402835291602001916111e2565b820191906000526020600020905b8154815290600101906020018083116111c557829003601f168201915b50505050508152602001906001019061114a565b50505050925081805480602002602001604051908101604052809291908181526020016000905b828210156112c957838290600052602060002001805461123c90611e0d565b80601f016020809104026020016040519081016040528092919081815260200182805461126890611e0d565b80156112b55780601f1061128a576101008083540402835291602001916112b5565b820191906000526020600020905b81548152906001019060200180831161129857829003601f168201915b50505050508152602001906001019061121d565b5050505091508080548060200260200160405190810160405280929190818152602001828054801561137557602002820191906000526020600020906000905b82829054906101000a900460ff16600281111561134f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200190600101906020826000010492830192600103820291508084116113095790505b50505050509050965096509650965096509650965090919293949596565b82805461139f90611e0d565b90600052602060002090601f0160209004810192826113c15760008555611408565b82601f106113da57805160ff1916838001178555611408565b82800160010185558215611408579182015b828111156114075782518255916020019190600101906113ec565b5b5090506114159190611419565b5090565b5b8082111561143257600081600090555060010161141a565b5090565b600061144961144484611c4b565b611c1a565b90508281526020810184848401111561146157600080fd5b61146c848285611dcb565b509392505050565b60008135905061148381611f69565b92915050565b60008135905061149881611f80565b92915050565b60008083601f8401126114b057600080fd5b8235905067ffffffffffffffff8111156114c957600080fd5b6020830191508360018202830111156114e157600080fd5b9250929050565b600082601f8301126114f957600080fd5b8135611509848260208601611436565b91505092915050565b60006020828403121561152457600080fd5b600061153284828501611474565b91505092915050565b6000806020838503121561154e57600080fd5b600083013567ffffffffffffffff81111561156857600080fd5b6115748582860161149e565b92509250509250929050565b60006020828403121561159257600080fd5b600082013567ffffffffffffffff8111156115ac57600080fd5b6115b8848285016114e8565b91505092915050565b600080604083850312156115d457600080fd5b600083013567ffffffffffffffff8111156115ee57600080fd5b6115fa858286016114e8565b925050602083013567ffffffffffffffff81111561161757600080fd5b611623858286016114e8565b9150509250929050565b60008060006060848603121561164257600080fd5b600084013567ffffffffffffffff81111561165c57600080fd5b611668868287016114e8565b935050602084013567ffffffffffffffff81111561168557600080fd5b611691868287016114e8565b92505060406116a286828701611489565b9150509250925092565b600080600080608085870312156116c257600080fd5b600085013567ffffffffffffffff8111156116dc57600080fd5b6116e8878288016114e8565b945050602085013567ffffffffffffffff81111561170557600080fd5b611711878288016114e8565b935050604085013567ffffffffffffffff81111561172e57600080fd5b61173a878288016114e8565b925050606085013567ffffffffffffffff81111561175757600080fd5b611763878288016114e8565b91505092959194509250565b600061177b83836117b3565b60208301905092915050565b60006117938383611911565b60208301905092915050565b60006117ab8383611920565b905092915050565b6117bc81611d5e565b82525050565b6117cb81611d5e565b82525050565b60006117dc82611cab565b6117e68185611cfe565b93506117f183611c7b565b8060005b83811015611822578151611809888261176f565b975061181483611cd7565b9250506001810190506117f5565b5085935050505092915050565b600061183a82611cb6565b6118448185611d20565b935061184f83611c8b565b8060005b838110156118805781516118678882611787565b975061187283611ce4565b925050600181019050611853565b5085935050505092915050565b600061189882611cc1565b6118a28185611d0f565b9350836020820285016118b485611c9b565b8060005b858110156118f057848403895281516118d1858261179f565b94506118dc83611cf1565b925060208a019950506001810190506118b8565b50829750879550505050505092915050565b61190b81611d70565b82525050565b61191a81611db9565b82525050565b600061192b82611ccc565b6119358185611d31565b9350611945818560208601611dda565b61194e81611f44565b840191505092915050565b600061196482611ccc565b61196e8185611d42565b935061197e818560208601611dda565b61198781611f44565b840191505092915050565b600061199d82611ccc565b6119a78185611d53565b93506119b7818560208601611dda565b80840191505092915050565b60006119d0602183611d42565b91507f4f6e6c792061646d696e2063616e20706572666f6d207468697320616374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611a36601883611d42565b91507f726567697374657220666972737420746f2076657269667900000000000000006000830152602082019050919050565b611a7281611daf565b82525050565b6000611a848284611992565b915081905092915050565b6000608082019050611aa460008301876117c2565b8181036020830152611ab68186611959565b90508181036040830152611aca8185611959565b9050611ad96060830184611902565b95945050505050565b600060e0820190508181036000830152611afc818a6117d1565b90508181036020830152611b1081896117d1565b90508181036040830152611b24818861188d565b90508181036060830152611b38818761188d565b90508181036080830152611b4c818661188d565b905081810360a0830152611b60818561188d565b905081810360c0830152611b74818461182f565b905098975050505050505050565b6000602082019050611b976000830184611902565b92915050565b60006020820190508181036000830152611bb78184611959565b905092915050565b60006020820190508181036000830152611bd8816119c3565b9050919050565b60006020820190508181036000830152611bf881611a29565b9050919050565b6000602082019050611c146000830184611a69565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611c4157611c40611f15565b5b8060405250919050565b600067ffffffffffffffff821115611c6657611c65611f15565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611d6982611d8f565b9050919050565b60008115159050919050565b6000819050611d8a82611f55565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611dc482611d7c565b9050919050565b82818337600083830152505050565b60005b83811015611df8578082015181840152602081019050611ddd565b83811115611e07576000848401525b50505050565b60006002820490506001821680611e2557607f821691505b60208210811415611e3957611e38611ee6565b5b50919050565b6000611e4a82611daf565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e7d57611e7c611e88565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60038110611f6657611f65611eb7565b5b50565b611f7281611d5e565b8114611f7d57600080fd5b50565b60038110611f8d57600080fd5b5056fea26469706673582212202922fe893dc7237b676db33d0b284e9f27f670956420e67a732df7de5164c50d64736f6c63430008000033",
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

// GetDocIndexCounter is a free data retrieval call binding the contract method 0xb8b4da4c.
//
// Solidity: function getDocIndexCounter() view returns(uint256)
func (_Verify *VerifyCaller) GetDocIndexCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getDocIndexCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDocIndexCounter is a free data retrieval call binding the contract method 0xb8b4da4c.
//
// Solidity: function getDocIndexCounter() view returns(uint256)
func (_Verify *VerifySession) GetDocIndexCounter() (*big.Int, error) {
	return _Verify.Contract.GetDocIndexCounter(&_Verify.CallOpts)
}

// GetDocIndexCounter is a free data retrieval call binding the contract method 0xb8b4da4c.
//
// Solidity: function getDocIndexCounter() view returns(uint256)
func (_Verify *VerifyCallerSession) GetDocIndexCounter() (*big.Int, error) {
	return _Verify.Contract.GetDocIndexCounter(&_Verify.CallOpts)
}

// GetDocumentIndex is a free data retrieval call binding the contract method 0x8e761664.
//
// Solidity: function getDocumentIndex(string shaHash) view returns(uint256)
func (_Verify *VerifyCaller) GetDocumentIndex(opts *bind.CallOpts, shaHash string) (*big.Int, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getDocumentIndex", shaHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDocumentIndex is a free data retrieval call binding the contract method 0x8e761664.
//
// Solidity: function getDocumentIndex(string shaHash) view returns(uint256)
func (_Verify *VerifySession) GetDocumentIndex(shaHash string) (*big.Int, error) {
	return _Verify.Contract.GetDocumentIndex(&_Verify.CallOpts, shaHash)
}

// GetDocumentIndex is a free data retrieval call binding the contract method 0x8e761664.
//
// Solidity: function getDocumentIndex(string shaHash) view returns(uint256)
func (_Verify *VerifyCallerSession) GetDocumentIndex(shaHash string) (*big.Int, error) {
	return _Verify.Contract.GetDocumentIndex(&_Verify.CallOpts, shaHash)
}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] hash, uint8[] stats)
func (_Verify *VerifyCaller) GetDocuments(opts *bind.CallOpts) (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Hash      []string
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
		Hash      []string
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
	outstruct.Hash = *abi.ConvertType(out[5], new([]string)).(*[]string)
	outstruct.Stats = *abi.ConvertType(out[6], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] hash, uint8[] stats)
func (_Verify *VerifySession) GetDocuments() (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Hash      []string
	Stats     []uint8
}, error) {
	return _Verify.Contract.GetDocuments(&_Verify.CallOpts)
}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] ipfs, string[] name, string[] hash, uint8[] stats)
func (_Verify *VerifyCallerSession) GetDocuments() (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Ipfs      []string
	Name      []string
	Hash      []string
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
// Solidity: function addDocument(string shaHash, string _EncryptedIPFSHash, string _institute, string _name) returns()
func (_Verify *VerifyTransactor) AddDocument(opts *bind.TransactOpts, shaHash string, _EncryptedIPFSHash string, _institute string, _name string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "addDocument", shaHash, _EncryptedIPFSHash, _institute, _name)
}

// AddDocument is a paid mutator transaction binding the contract method 0x0e72ca7a.
//
// Solidity: function addDocument(string shaHash, string _EncryptedIPFSHash, string _institute, string _name) returns()
func (_Verify *VerifySession) AddDocument(shaHash string, _EncryptedIPFSHash string, _institute string, _name string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, shaHash, _EncryptedIPFSHash, _institute, _name)
}

// AddDocument is a paid mutator transaction binding the contract method 0x0e72ca7a.
//
// Solidity: function addDocument(string shaHash, string _EncryptedIPFSHash, string _institute, string _name) returns()
func (_Verify *VerifyTransactorSession) AddDocument(shaHash string, _EncryptedIPFSHash string, _institute string, _name string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, shaHash, _EncryptedIPFSHash, _institute, _name)
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
// Solidity: function verifyDocument(string shaHash, string _institute, uint8 _status) payable returns()
func (_Verify *VerifyTransactor) VerifyDocument(opts *bind.TransactOpts, shaHash string, _institute string, _status uint8) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "verifyDocument", shaHash, _institute, _status)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x90752bb2.
//
// Solidity: function verifyDocument(string shaHash, string _institute, uint8 _status) payable returns()
func (_Verify *VerifySession) VerifyDocument(shaHash string, _institute string, _status uint8) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocument(&_Verify.TransactOpts, shaHash, _institute, _status)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x90752bb2.
//
// Solidity: function verifyDocument(string shaHash, string _institute, uint8 _status) payable returns()
func (_Verify *VerifyTransactorSession) VerifyDocument(shaHash string, _institute string, _status uint8) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocument(&_Verify.TransactOpts, shaHash, _institute, _status)
}
