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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_docAddressOnIPFS\",\"type\":\"string\"}],\"name\":\"addDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"approveVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocumentList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"requester\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"verifer\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"name\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"desc\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"ipfsAddress\",\"type\":\"string[]\"},{\"internalType\":\"enumVerification.DocStatus[]\",\"name\":\"stats\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"userDocId\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"}],\"name\":\"registerAsUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_aadharNum\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"}],\"name\":\"registerAsVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"AadharNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"institute\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isApprovedVerifier\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_docAddressOnIPFS\",\"type\":\"string\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"verifyDocuments\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600b5534801561001557600080fd5b5033600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550612802806100666000396000f3fe60806040526004361061007b5760003560e01c8063cfc197b01161004e578063cfc197b01461013e578063d6307acd14610167578063f036d2e814610190578063fb63030c146101c15761007b565b80631e9b13121461008057806320857a9f146100ab57806336a8915b146100d45780636c824487146100fd575b600080fd5b34801561008c57600080fd5b506100956101dd565b6040516100a29190612293565b60405180910390f35b3480156100b757600080fd5b506100d260048036038101906100cd9190611c4f565b610203565b005b3480156100e057600080fd5b506100fb60048036038101906100f69190611cc4565b61037a565b005b34801561010957600080fd5b50610124600480360381019061011f9190611c26565b61096b565b60405161013595949392919061234e565b60405180910390f35b34801561014a57600080fd5b5061016560048036038101906101609190611c26565b610bce565b005b34801561017357600080fd5b5061018e60048036038101906101899190611dbc565b610cbb565b005b34801561019c57600080fd5b506101a5610dc0565b6040516101b897969594939291906122ae565b60405180910390f35b6101db60048036038101906101d69190611d68565b61176f565b005b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604051806040016040528085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815250600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190805190602001906102fb929190611a21565b506020820151816001019080519060200190610318929190611a21565b509050506001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690506001151581151514610412576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610409906123dd565b60405180910390fd5b60006040518060e001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200189898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152602001600b548152602001600280811115610577577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152509050806002858560405161058f929190612263565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019080519060200190610647929190611a21565b506060820151816003019080519060200190610664929190611a21565b506080820151816004019080519060200190610681929190611a21565b5060a0820151816005015560c08201518160060160006101000a81548160ff021916908360028111156106dd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550905050600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600b5490806001815401808255809150506001900390600052602060002001600090919091909150556005339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506007888890918060018154018082558091505060019003906000526020600020016000909192909192909192909192509190610851929190611aa7565b50600886869091806001815401808255809150506001900390600052602060002001600090919290919290919290919250919061088f929190611aa7565b5060098484909180600181540180825580915050600190039060005260206000200160009091929091929091929091925091906108cd929190611aa7565b50600a600290806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff02191690836002811115610944577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550600b600081548092919061095c9061267b565b91905055505050505050505050565b600060205280600052604060002060009150905080600001805461098e90612649565b80601f01602080910402602001604051908101604052809291908181526020018280546109ba90612649565b8015610a075780601f106109dc57610100808354040283529160200191610a07565b820191906000526020600020905b8154815290600101906020018083116109ea57829003601f168201915b505050505090806001018054610a1c90612649565b80601f0160208091040260200160405190810160405280929190818152602001828054610a4890612649565b8015610a955780601f10610a6a57610100808354040283529160200191610a95565b820191906000526020600020905b815481529060010190602001808311610a7857829003601f168201915b505050505090806002018054610aaa90612649565b80601f0160208091040260200160405190810160405280929190818152602001828054610ad690612649565b8015610b235780601f10610af857610100808354040283529160200191610b23565b820191906000526020600020905b815481529060010190602001808311610b0657829003601f168201915b505050505090806003018054610b3890612649565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6490612649565b8015610bb15780601f10610b8657610100808354040283529160200191610bb1565b820191906000526020600020905b815481529060010190602001808311610b9457829003601f168201915b5050505050908060040160009054906101000a900460ff16905085565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c55906123bd565b60405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff02191690831515021790555050565b6040518060a00160405280858152602001848152602001838152602001828152602001600015158152506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190610d3f929190611a21565b506020820151816001019080519060200190610d5c929190611a21565b506040820151816002019080519060200190610d79929190611a21565b506060820151816003019080519060200190610d96929190611a21565b5060808201518160040160006101000a81548160ff02191690831515021790555090505050505050565b60608060608060608060606000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff16156112e45760056006600760086009600a600067ffffffffffffffff811115610e6c577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015610e9a5781602001602082028036833780820191505090505b5086805480602002602001604051908101604052809291908181526020018280548015610f1c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610ed2575b5050505050965085805480602002602001604051908101604052809291908181526020018280548015610fa457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610f5a575b5050505050955084805480602002602001604051908101604052809291908181526020016000905b82821015611078578382906000526020600020018054610feb90612649565b80601f016020809104026020016040519081016040528092919081815260200182805461101790612649565b80156110645780601f1061103957610100808354040283529160200191611064565b820191906000526020600020905b81548152906001019060200180831161104757829003601f168201915b505050505081526020019060010190610fcc565b50505050945083805480602002602001604051908101604052809291908181526020016000905b8282101561114b5783829060005260206000200180546110be90612649565b80601f01602080910402602001604051908101604052809291908181526020018280546110ea90612649565b80156111375780601f1061110c57610100808354040283529160200191611137565b820191906000526020600020905b81548152906001019060200180831161111a57829003601f168201915b50505050508152602001906001019061109f565b50505050935082805480602002602001604051908101604052809291908181526020016000905b8282101561121e57838290600052602060002001805461119190612649565b80601f01602080910402602001604051908101604052809291908181526020018280546111bd90612649565b801561120a5780601f106111df5761010080835404028352916020019161120a565b820191906000526020600020905b8154815290600101906020018083116111ed57829003601f168201915b505050505081526020019060010190611172565b505050509250818054806020026020016040519081016040528092919081815260200182805480156112ca57602002820191906000526020600020906000905b82829054906101000a900460ff1660028111156112a4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001906001019060208260000104928301926001038202915080841161125e5790505b505050505091509650965096509650965096509650611766565b6005600660076008600067ffffffffffffffff81111561132d577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561136057816020015b606081526020019060019003908161134b5790505b50600a600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208680548060200260200160405190810160405280929190818152602001828054801561142357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116113d9575b50505050509650858054806020026020016040519081016040528092919081815260200182805480156114ab57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611461575b5050505050955084805480602002602001604051908101604052809291908181526020016000905b8282101561157f5783829060005260206000200180546114f290612649565b80601f016020809104026020016040519081016040528092919081815260200182805461151e90612649565b801561156b5780601f106115405761010080835404028352916020019161156b565b820191906000526020600020905b81548152906001019060200180831161154e57829003601f168201915b5050505050815260200190600101906114d3565b50505050945083805480602002602001604051908101604052809291908181526020016000905b828210156116525783829060005260206000200180546115c590612649565b80601f01602080910402602001604051908101604052809291908181526020018280546115f190612649565b801561163e5780601f106116135761010080835404028352916020019161163e565b820191906000526020600020905b81548152906001019060200180831161162157829003601f168201915b5050505050815260200190600101906115a6565b505050509350818054806020026020016040519081016040528092919081815260200182805480156116fe57602002820191906000526020600020906000905b82829054906101000a900460ff1660028111156116d8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200190600101906020826000010492830192600103820291508084116116925790505b505050505091508080548060200260200160405190810160405280929190818152602001828054801561175057602002820191906000526020600020905b81548152602001906001019080831161173c575b5050505050905096509650965096509650965096505b90919293949596565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff1690506001151581151514611809576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611800906123fd565b60405180910390fd5b8160028460405161181a919061227c565b908152602001604051809103902060060160006101000a81548160ff02191690836002811115611873577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555033600284604051611889919061227c565b908152602001604051809103902060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006002846040516118ea919061227c565b908152602001604051809103902060050154905082600a8281548110611939577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090602091828204019190066101000a81548160ff02191690836002811115611993577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555033600682815481106119d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b828054611a2d90612649565b90600052602060002090601f016020900481019282611a4f5760008555611a96565b82601f10611a6857805160ff1916838001178555611a96565b82800160010185558215611a96579182015b82811115611a95578251825591602001919060010190611a7a565b5b509050611aa39190611b2d565b5090565b828054611ab390612649565b90600052602060002090601f016020900481019282611ad55760008555611b1c565b82601f10611aee57803560ff1916838001178555611b1c565b82800160010185558215611b1c579182015b82811115611b1b578235825591602001919060010190611b00565b5b509050611b299190611b2d565b5090565b5b80821115611b46576000816000905550600101611b2e565b5090565b6000611b5d611b588461244e565b61241d565b905082815260208101848484011115611b7557600080fd5b611b80848285612607565b509392505050565b600081359050611b97816127a5565b92915050565b600081359050611bac816127bc565b92915050565b60008083601f840112611bc457600080fd5b8235905067ffffffffffffffff811115611bdd57600080fd5b602083019150836001820283011115611bf557600080fd5b9250929050565b600082601f830112611c0d57600080fd5b8135611c1d848260208601611b4a565b91505092915050565b600060208284031215611c3857600080fd5b6000611c4684828501611b88565b91505092915050565b60008060008060408587031215611c6557600080fd5b600085013567ffffffffffffffff811115611c7f57600080fd5b611c8b87828801611bb2565b9450945050602085013567ffffffffffffffff811115611caa57600080fd5b611cb687828801611bb2565b925092505092959194509250565b60008060008060008060608789031215611cdd57600080fd5b600087013567ffffffffffffffff811115611cf757600080fd5b611d0389828a01611bb2565b9650965050602087013567ffffffffffffffff811115611d2257600080fd5b611d2e89828a01611bb2565b9450945050604087013567ffffffffffffffff811115611d4d57600080fd5b611d5989828a01611bb2565b92509250509295509295509295565b60008060408385031215611d7b57600080fd5b600083013567ffffffffffffffff811115611d9557600080fd5b611da185828601611bfc565b9250506020611db285828601611b9d565b9150509250929050565b60008060008060808587031215611dd257600080fd5b600085013567ffffffffffffffff811115611dec57600080fd5b611df887828801611bfc565b945050602085013567ffffffffffffffff811115611e1557600080fd5b611e2187828801611bfc565b935050604085013567ffffffffffffffff811115611e3e57600080fd5b611e4a87828801611bfc565b925050606085013567ffffffffffffffff811115611e6757600080fd5b611e7387828801611bfc565b91505092959194509250565b6000611e8b8383611edb565b60208301905092915050565b6000611ea38383612097565b60208301905092915050565b6000611ebb83836120cb565b905092915050565b6000611ecf8383612254565b60208301905092915050565b611ee48161259a565b82525050565b611ef38161259a565b82525050565b6000611f04826124be565b611f0e8185612529565b9350611f198361247e565b8060005b83811015611f4a578151611f318882611e7f565b9750611f3c836124f5565b925050600181019050611f1d565b5085935050505092915050565b6000611f62826124c9565b611f6c818561255c565b9350611f778361248e565b8060005b83811015611fa8578151611f8f8882611e97565b9750611f9a83612502565b925050600181019050611f7b565b5085935050505092915050565b6000611fc0826124d4565b611fca818561253a565b935083602082028501611fdc8561249e565b8060005b858110156120185784840389528151611ff98582611eaf565b94506120048361250f565b925060208a01995050600181019050611fe0565b50829750879550505050505092915050565b6000612035826124df565b61203f818561254b565b935061204a836124ae565b8060005b8381101561207b5781516120628882611ec3565b975061206d8361251c565b92505060018101905061204e565b5085935050505092915050565b612091816125ac565b82525050565b6120a0816125f5565b82525050565b60006120b2838561258f565b93506120bf838584612607565b82840190509392505050565b60006120d6826124ea565b6120e0818561256d565b93506120f0818560208601612616565b6120f981612780565b840191505092915050565b600061210f826124ea565b612119818561257e565b9350612129818560208601612616565b61213281612780565b840191505092915050565b6000612148826124ea565b612152818561258f565b9350612162818560208601612616565b80840191505092915050565b600061217b60218361257e565b91507f4f6e6c792061646d696e2063616e20706572666f6d207468697320616374696f60008301527f6e000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006121e160188361257e565b91507f726567697374657220666972737420746f2076657269667900000000000000006000830152602082019050919050565b600061222160138361257e565b91507f596f75277265206e6f74207665726966696572000000000000000000000000006000830152602082019050919050565b61225d816125eb565b82525050565b60006122708284866120a6565b91508190509392505050565b6000612288828461213d565b915081905092915050565b60006020820190506122a86000830184611eea565b92915050565b600060e08201905081810360008301526122c8818a611ef9565b905081810360208301526122dc8189611ef9565b905081810360408301526122f08188611fb5565b905081810360608301526123048187611fb5565b905081810360808301526123188186611fb5565b905081810360a083015261232c8185611f57565b905081810360c0830152612340818461202a565b905098975050505050505050565b600060a08201905081810360008301526123688188612104565b9050818103602083015261237c8187612104565b905081810360408301526123908186612104565b905081810360608301526123a48185612104565b90506123b36080830184612088565b9695505050505050565b600060208201905081810360008301526123d68161216e565b9050919050565b600060208201905081810360008301526123f6816121d4565b9050919050565b6000602082019050818103600083015261241681612214565b9050919050565b6000604051905081810181811067ffffffffffffffff8211171561244457612443612751565b5b8060405250919050565b600067ffffffffffffffff82111561246957612468612751565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006125a5826125cb565b9050919050565b60008115159050919050565b60008190506125c682612791565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000612600826125b8565b9050919050565b82818337600083830152505050565b60005b83811015612634578082015181840152602081019050612619565b83811115612643576000848401525b50505050565b6000600282049050600182168061266157607f821691505b6020821081141561267557612674612722565b5b50919050565b6000612686826125eb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156126b9576126b86126c4565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b600381106127a2576127a16126f3565b5b50565b6127ae8161259a565b81146127b957600080fd5b50565b600381106127c957600080fd5b5056fea2646970667358221220e307d66c2270b16759ce713b368f074eb354481c4e7b96faf17b685e47cd62d364736f6c63430008000033",
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

// AccountAddress is a free data retrieval call binding the contract method 0x1e9b1312.
//
// Solidity: function accountAddress() view returns(address)
func (_Verify *VerifyCaller) AccountAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "accountAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAddress is a free data retrieval call binding the contract method 0x1e9b1312.
//
// Solidity: function accountAddress() view returns(address)
func (_Verify *VerifySession) AccountAddress() (common.Address, error) {
	return _Verify.Contract.AccountAddress(&_Verify.CallOpts)
}

// AccountAddress is a free data retrieval call binding the contract method 0x1e9b1312.
//
// Solidity: function accountAddress() view returns(address)
func (_Verify *VerifyCallerSession) AccountAddress() (common.Address, error) {
	return _Verify.Contract.AccountAddress(&_Verify.CallOpts)
}

// GetDocumentList is a free data retrieval call binding the contract method 0xf036d2e8.
//
// Solidity: function getDocumentList() view returns(address[] requester, address[] verifer, string[] name, string[] desc, string[] ipfsAddress, uint8[] stats, uint256[] userDocId)
func (_Verify *VerifyCaller) GetDocumentList(opts *bind.CallOpts) (struct {
	Requester   []common.Address
	Verifer     []common.Address
	Name        []string
	Desc        []string
	IpfsAddress []string
	Stats       []uint8
	UserDocId   []*big.Int
}, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "getDocumentList")

	outstruct := new(struct {
		Requester   []common.Address
		Verifer     []common.Address
		Name        []string
		Desc        []string
		IpfsAddress []string
		Stats       []uint8
		UserDocId   []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Verifer = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Name = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Desc = *abi.ConvertType(out[3], new([]string)).(*[]string)
	outstruct.IpfsAddress = *abi.ConvertType(out[4], new([]string)).(*[]string)
	outstruct.Stats = *abi.ConvertType(out[5], new([]uint8)).(*[]uint8)
	outstruct.UserDocId = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetDocumentList is a free data retrieval call binding the contract method 0xf036d2e8.
//
// Solidity: function getDocumentList() view returns(address[] requester, address[] verifer, string[] name, string[] desc, string[] ipfsAddress, uint8[] stats, uint256[] userDocId)
func (_Verify *VerifySession) GetDocumentList() (struct {
	Requester   []common.Address
	Verifer     []common.Address
	Name        []string
	Desc        []string
	IpfsAddress []string
	Stats       []uint8
	UserDocId   []*big.Int
}, error) {
	return _Verify.Contract.GetDocumentList(&_Verify.CallOpts)
}

// GetDocumentList is a free data retrieval call binding the contract method 0xf036d2e8.
//
// Solidity: function getDocumentList() view returns(address[] requester, address[] verifer, string[] name, string[] desc, string[] ipfsAddress, uint8[] stats, uint256[] userDocId)
func (_Verify *VerifyCallerSession) GetDocumentList() (struct {
	Requester   []common.Address
	Verifer     []common.Address
	Name        []string
	Desc        []string
	IpfsAddress []string
	Stats       []uint8
	UserDocId   []*big.Int
}, error) {
	return _Verify.Contract.GetDocumentList(&_Verify.CallOpts)
}

// Verifiers is a free data retrieval call binding the contract method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(string name, string email, string AadharNumber, string institute, bool isApprovedVerifier)
func (_Verify *VerifyCaller) Verifiers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name               string
	Email              string
	AadharNumber       string
	Institute          string
	IsApprovedVerifier bool
}, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "verifiers", arg0)

	outstruct := new(struct {
		Name               string
		Email              string
		AadharNumber       string
		Institute          string
		IsApprovedVerifier bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AadharNumber = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Institute = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsApprovedVerifier = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Verifiers is a free data retrieval call binding the contract method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(string name, string email, string AadharNumber, string institute, bool isApprovedVerifier)
func (_Verify *VerifySession) Verifiers(arg0 common.Address) (struct {
	Name               string
	Email              string
	AadharNumber       string
	Institute          string
	IsApprovedVerifier bool
}, error) {
	return _Verify.Contract.Verifiers(&_Verify.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(string name, string email, string AadharNumber, string institute, bool isApprovedVerifier)
func (_Verify *VerifyCallerSession) Verifiers(arg0 common.Address) (struct {
	Name               string
	Email              string
	AadharNumber       string
	Institute          string
	IsApprovedVerifier bool
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

// ApproveVerifier is a paid mutator transaction binding the contract method 0xcfc197b0.
//
// Solidity: function approveVerifier(address verifier) returns()
func (_Verify *VerifyTransactor) ApproveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _Verify.contract.Transact(opts, "approveVerifier", verifier)
}

// ApproveVerifier is a paid mutator transaction binding the contract method 0xcfc197b0.
//
// Solidity: function approveVerifier(address verifier) returns()
func (_Verify *VerifySession) ApproveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _Verify.Contract.ApproveVerifier(&_Verify.TransactOpts, verifier)
}

// ApproveVerifier is a paid mutator transaction binding the contract method 0xcfc197b0.
//
// Solidity: function approveVerifier(address verifier) returns()
func (_Verify *VerifyTransactorSession) ApproveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _Verify.Contract.ApproveVerifier(&_Verify.TransactOpts, verifier)
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
