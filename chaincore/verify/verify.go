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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_requestor\",\"type\":\"address\"}],\"name\":\"addCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shaHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"}],\"name\":\"addDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"approveVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocIndexCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocuments\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"requester\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"verifer\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"institute\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"hash\",\"type\":\"string[]\"},{\"internalType\":\"enumVerification.DocStatus[]\",\"name\":\"stats\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getInstituePublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getUserPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"institutions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"publicAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"isApprovedInstitute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"registerAsUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"registerInstitution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shaHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_proofHash\",\"type\":\"string\"}],\"name\":\"verifyDocument\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600b5534801561001557600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506121fa806100656000396000f3fe6080604052600436106100a75760003560e01c8063ba1f3dc711610064578063ba1f3dc7146101ab578063ba785725146101e8578063c6b5c5cc14610211578063e74c03421461023a578063ea6509f814610277578063ef2d8700146102b7576100a7565b806311231fe0146100ac578063188c71d9146100e9578063362846dc146101055780634d2b19781461012e578063515ef10614610157578063b8b4da4c14610180575b600080fd5b3480156100b857600080fd5b506100d360048036038101906100ce9190611728565b6102e6565b6040516100e09190611db1565b60405180910390f35b61010360048036038101906100fe91906118c2565b6103ba565b005b34801561011157600080fd5b5061012c60048036038101906101279190611843565b610674565b005b34801561013a57600080fd5b50610155600480360381019061015091906117d7565b61091e565b005b34801561016357600080fd5b5061017e60048036038101906101799190611751565b610c1c565b005b34801561018c57600080fd5b50610195610d2a565b6040516101a29190611e33565b60405180910390f35b3480156101b757600080fd5b506101d260048036038101906101cd9190611796565b610d34565b6040516101df9190611db1565b60405180910390f35b3480156101f457600080fd5b5061020f600480360381019061020a9190611796565b610de7565b005b34801561021d57600080fd5b50610238600480360381019061023391906117d7565b610eb1565b005b34801561024657600080fd5b50610261600480360381019061025c9190611796565b611061565b60405161026e9190611d96565b60405180910390f35b34801561028357600080fd5b5061029e60048036038101906102999190611796565b6110a0565b6040516102ae9493929190611ccd565b60405180910390f35b3480156102c357600080fd5b506102cc611223565b6040516102dd959493929190611d20565b60405180910390f35b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001805461033590612041565b80601f016020809104026020016040519081016040528092919081815260200182805461036190612041565b80156103ae5780601f10610383576101008083540402835291602001916103ae565b820191906000526020600020905b81548152906001019060200180831161039157829003601f168201915b50505050509050919050565b600115156001846040516103ce9190611cb6565b908152602001604051809103902060030160009054906101000a900460ff16151514801561046857503373ffffffffffffffffffffffffffffffffffffffff1660018460405161041e9190611cb6565b908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b61047157600080fd5b60006003856040516104839190611cb6565b908152602001604051809103902054905082600a82815481106104cf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090602091828204019190066101000a81548160ff02191690836002811115610529577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060038560405161053e9190611cb6565b90815260200160405180910390205460038360405161055d9190611cb6565b90815260200160405180910390208190555081600782815481106105aa577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200190805190602001906105c79291906115a9565b506003856040516105d89190611cb6565b9081526020016040518091039020600090553360068281548110610625577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b6001826040516106849190611cb6565b908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610723576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071a90611dd3565b60405180910390fd5b600b546003846040516107369190611cb6565b9081526020016040518091039020819055506005819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506008829080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906108499291906115a9565b50600a600090806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff021916908360028111156108c0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055506007839080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906109009291906115a9565b50600b600081548092919061091490612073565b9190505550505050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16806109e257506001816040516109819190611cb6565b908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610a21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1890611dd3565b60405180910390fd5b600b54600383604051610a349190611cb6565b9081526020016040518091039020819055506005339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600881908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610b489291906115a9565b50600a600290806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff02191690836002811115610bbf577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550600782908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610bff9291906115a9565b50600b6000815480929190610c1390612073565b91905055505050565b604051806020016040528083838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815250600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190610cca9291906115a9565b509050506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000600b54905090565b6060600182604051610d469190611cb6565b90815260200160405180910390206001018054610d6290612041565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8e90612041565b8015610ddb5780601f10610db057610100808354040283529160200191610ddb565b820191906000526020600020905b815481529060010190602001808311610dbe57829003601f168201915b50505050509050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6c90611e13565b60405180910390fd5b60018082604051610e869190611cb6565b908152602001604051809103902060030160006101000a81548160ff02191690831515021790555050565b600073ffffffffffffffffffffffffffffffffffffffff16600182604051610ed99190611cb6565b908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610f61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5890611df3565b60405180910390fd5b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200160001515815250600182604051610fab9190611cb6565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908051906020019061101c9291906115a9565b5060408201518160020190805190602001906110399291906115a9565b5060608201518160030160006101000a81548160ff0219169083151502179055509050505050565b6000600115156001836040516110779190611cb6565b908152602001604051809103902060030160009054906101000a900460ff161515149050919050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010180546110ff90612041565b80601f016020809104026020016040519081016040528092919081815260200182805461112b90612041565b80156111785780601f1061114d57610100808354040283529160200191611178565b820191906000526020600020905b81548152906001019060200180831161115b57829003601f168201915b50505050509080600201805461118d90612041565b80601f01602080910402602001604051908101604052809291908181526020018280546111b990612041565b80156112065780601f106111db57610100808354040283529160200191611206565b820191906000526020600020905b8154815290600101906020018083116111e957829003601f168201915b5050505050908060030160009054906101000a900460ff16905084565b60608060608060606005600660086007600a848054806020026020016040519081016040528092919081815260200182805480156112b657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161126c575b505050505094508380548060200260200160405190810160405280929190818152602001828054801561133e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116112f4575b5050505050935082805480602002602001604051908101604052809291908181526020016000905b8282101561141257838290600052602060002001805461138590612041565b80601f01602080910402602001604051908101604052809291908181526020018280546113b190612041565b80156113fe5780601f106113d3576101008083540402835291602001916113fe565b820191906000526020600020905b8154815290600101906020018083116113e157829003601f168201915b505050505081526020019060010190611366565b50505050925081805480602002602001604051908101604052809291908181526020016000905b828210156114e557838290600052602060002001805461145890612041565b80601f016020809104026020016040519081016040528092919081815260200182805461148490612041565b80156114d15780601f106114a6576101008083540402835291602001916114d1565b820191906000526020600020905b8154815290600101906020018083116114b457829003601f168201915b505050505081526020019060010190611439565b5050505091508080548060200260200160405190810160405280929190818152602001828054801561159157602002820191906000526020600020906000905b82829054906101000a900460ff16600281111561156b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200190600101906020826000010492830192600103820291508084116115255790505b50505050509050945094509450945094509091929394565b8280546115b590612041565b90600052602060002090601f0160209004810192826115d7576000855561161e565b82601f106115f057805160ff191683800117855561161e565b8280016001018555821561161e579182015b8281111561161d578251825591602001919060010190611602565b5b50905061162b919061162f565b5090565b5b80821115611648576000816000905550600101611630565b5090565b600061165f61165a84611e7f565b611e4e565b90508281526020810184848401111561167757600080fd5b611682848285611fff565b509392505050565b6000813590506116998161219d565b92915050565b6000813590506116ae816121b4565b92915050565b60008083601f8401126116c657600080fd5b8235905067ffffffffffffffff8111156116df57600080fd5b6020830191508360018202830111156116f757600080fd5b9250929050565b600082601f83011261170f57600080fd5b813561171f84826020860161164c565b91505092915050565b60006020828403121561173a57600080fd5b60006117488482850161168a565b91505092915050565b6000806020838503121561176457600080fd5b600083013567ffffffffffffffff81111561177e57600080fd5b61178a858286016116b4565b92509250509250929050565b6000602082840312156117a857600080fd5b600082013567ffffffffffffffff8111156117c257600080fd5b6117ce848285016116fe565b91505092915050565b600080604083850312156117ea57600080fd5b600083013567ffffffffffffffff81111561180457600080fd5b611810858286016116fe565b925050602083013567ffffffffffffffff81111561182d57600080fd5b611839858286016116fe565b9150509250929050565b60008060006060848603121561185857600080fd5b600084013567ffffffffffffffff81111561187257600080fd5b61187e868287016116fe565b935050602084013567ffffffffffffffff81111561189b57600080fd5b6118a7868287016116fe565b92505060406118b88682870161168a565b9150509250925092565b600080600080608085870312156118d857600080fd5b600085013567ffffffffffffffff8111156118f257600080fd5b6118fe878288016116fe565b945050602085013567ffffffffffffffff81111561191b57600080fd5b611927878288016116fe565b93505060406119388782880161169f565b925050606085013567ffffffffffffffff81111561195557600080fd5b611961878288016116fe565b91505092959194509250565b600061197983836119b1565b60208301905092915050565b60006119918383611b0f565b60208301905092915050565b60006119a98383611b1e565b905092915050565b6119ba81611f92565b82525050565b6119c981611f92565b82525050565b60006119da82611edf565b6119e48185611f32565b93506119ef83611eaf565b8060005b83811015611a20578151611a07888261196d565b9750611a1283611f0b565b9250506001810190506119f3565b5085935050505092915050565b6000611a3882611eea565b611a428185611f54565b9350611a4d83611ebf565b8060005b83811015611a7e578151611a658882611985565b9750611a7083611f18565b925050600181019050611a51565b5085935050505092915050565b6000611a9682611ef5565b611aa08185611f43565b935083602082028501611ab285611ecf565b8060005b85811015611aee5784840389528151611acf858261199d565b9450611ada83611f25565b925060208a01995050600181019050611ab6565b50829750879550505050505092915050565b611b0981611fa4565b82525050565b611b1881611fed565b82525050565b6000611b2982611f00565b611b338185611f65565b9350611b4381856020860161200e565b611b4c81612178565b840191505092915050565b6000611b6282611f00565b611b6c8185611f76565b9350611b7c81856020860161200e565b611b8581612178565b840191505092915050565b6000611b9b82611f00565b611ba58185611f87565b9350611bb581856020860161200e565b80840191505092915050565b6000611bce601883611f76565b91507f726567697374657220666972737420746f2075706c6f616400000000000000006000830152602082019050919050565b6000611c0e601e83611f76565b91507f496e737469747574696f6e20616c7265616479207265676973746572656400006000830152602082019050919050565b6000611c4e602183611f76565b91507f4f6e6c792061646d696e2063616e20706572666f6d207468697320616374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b611cb081611fe3565b82525050565b6000611cc28284611b90565b915081905092915050565b6000608082019050611ce260008301876119c0565b8181036020830152611cf48186611b57565b90508181036040830152611d088185611b57565b9050611d176060830184611b00565b95945050505050565b600060a0820190508181036000830152611d3a81886119cf565b90508181036020830152611d4e81876119cf565b90508181036040830152611d628186611a8b565b90508181036060830152611d768185611a8b565b90508181036080830152611d8a8184611a2d565b90509695505050505050565b6000602082019050611dab6000830184611b00565b92915050565b60006020820190508181036000830152611dcb8184611b57565b905092915050565b60006020820190508181036000830152611dec81611bc1565b9050919050565b60006020820190508181036000830152611e0c81611c01565b9050919050565b60006020820190508181036000830152611e2c81611c41565b9050919050565b6000602082019050611e486000830184611ca7565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611e7557611e74612149565b5b8060405250919050565b600067ffffffffffffffff821115611e9a57611e99612149565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611f9d82611fc3565b9050919050565b60008115159050919050565b6000819050611fbe82612189565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611ff882611fb0565b9050919050565b82818337600083830152505050565b60005b8381101561202c578082015181840152602081019050612011565b8381111561203b576000848401525b50505050565b6000600282049050600182168061205957607f821691505b6020821081141561206d5761206c61211a565b5b50919050565b600061207e82611fe3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156120b1576120b06120bc565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6003811061219a576121996120eb565b5b50565b6121a681611f92565b81146121b157600080fd5b50565b600381106121c157600080fd5b5056fea2646970667358221220983b96416bfeb5fc8b1bda2ae57f92682f2bdec72136eb83bd89bc362702c2e164736f6c63430008000033",
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

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] hash, uint8[] stats)
func (_Verify *VerifyCaller) GetDocuments(opts *bind.CallOpts) (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Hash      []string
	Stats     []uint8
}, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getDocuments")

	outstruct := new(struct {
		Requester []common.Address
		Verifer   []common.Address
		Institute []string
		Hash      []string
		Stats     []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Verifer = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Institute = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Hash = *abi.ConvertType(out[3], new([]string)).(*[]string)
	outstruct.Stats = *abi.ConvertType(out[4], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] hash, uint8[] stats)
func (_Verify *VerifySession) GetDocuments() (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
	Hash      []string
	Stats     []uint8
}, error) {
	return _Verify.Contract.GetDocuments(&_Verify.CallOpts)
}

// GetDocuments is a free data retrieval call binding the contract method 0xef2d8700.
//
// Solidity: function getDocuments() view returns(address[] requester, address[] verifer, string[] institute, string[] hash, uint8[] stats)
func (_Verify *VerifyCallerSession) GetDocuments() (struct {
	Requester []common.Address
	Verifer   []common.Address
	Institute []string
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

// AddCertificate is a paid mutator transaction binding the contract method 0x362846dc.
//
// Solidity: function addCertificate(string _hash, string _institute, address _requestor) returns()
func (_Verify *VerifyTransactor) AddCertificate(opts *bind.TransactOpts, _hash string, _institute string, _requestor common.Address) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "addCertificate", _hash, _institute, _requestor)
}

// AddCertificate is a paid mutator transaction binding the contract method 0x362846dc.
//
// Solidity: function addCertificate(string _hash, string _institute, address _requestor) returns()
func (_Verify *VerifySession) AddCertificate(_hash string, _institute string, _requestor common.Address) (*types.Transaction, error) {
	return _Verify.Contract.AddCertificate(&_Verify.TransactOpts, _hash, _institute, _requestor)
}

// AddCertificate is a paid mutator transaction binding the contract method 0x362846dc.
//
// Solidity: function addCertificate(string _hash, string _institute, address _requestor) returns()
func (_Verify *VerifyTransactorSession) AddCertificate(_hash string, _institute string, _requestor common.Address) (*types.Transaction, error) {
	return _Verify.Contract.AddCertificate(&_Verify.TransactOpts, _hash, _institute, _requestor)
}

// AddDocument is a paid mutator transaction binding the contract method 0x4d2b1978.
//
// Solidity: function addDocument(string shaHash, string _institute) returns()
func (_Verify *VerifyTransactor) AddDocument(opts *bind.TransactOpts, shaHash string, _institute string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "addDocument", shaHash, _institute)
}

// AddDocument is a paid mutator transaction binding the contract method 0x4d2b1978.
//
// Solidity: function addDocument(string shaHash, string _institute) returns()
func (_Verify *VerifySession) AddDocument(shaHash string, _institute string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, shaHash, _institute)
}

// AddDocument is a paid mutator transaction binding the contract method 0x4d2b1978.
//
// Solidity: function addDocument(string shaHash, string _institute) returns()
func (_Verify *VerifyTransactorSession) AddDocument(shaHash string, _institute string) (*types.Transaction, error) {
	return _Verify.Contract.AddDocument(&_Verify.TransactOpts, shaHash, _institute)
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

// VerifyDocument is a paid mutator transaction binding the contract method 0x188c71d9.
//
// Solidity: function verifyDocument(string shaHash, string _institute, uint8 _status, string _proofHash) payable returns()
func (_Verify *VerifyTransactor) VerifyDocument(opts *bind.TransactOpts, shaHash string, _institute string, _status uint8, _proofHash string) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "verifyDocument", shaHash, _institute, _status, _proofHash)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x188c71d9.
//
// Solidity: function verifyDocument(string shaHash, string _institute, uint8 _status, string _proofHash) payable returns()
func (_Verify *VerifySession) VerifyDocument(shaHash string, _institute string, _status uint8, _proofHash string) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocument(&_Verify.TransactOpts, shaHash, _institute, _status, _proofHash)
}

// VerifyDocument is a paid mutator transaction binding the contract method 0x188c71d9.
//
// Solidity: function verifyDocument(string shaHash, string _institute, uint8 _status, string _proofHash) payable returns()
func (_Verify *VerifyTransactorSession) VerifyDocument(shaHash string, _institute string, _status uint8, _proofHash string) (*types.Transaction, error) {
	return _Verify.Contract.VerifyDocument(&_Verify.TransactOpts, shaHash, _institute, _status, _proofHash)
}
