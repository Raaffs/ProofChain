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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shaHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"}],\"name\":\"addDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"approveVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocIndexCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDocuments\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"requester\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"verifer\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"institute\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"hash\",\"type\":\"string[]\"},{\"internalType\":\"enumVerification.DocStatus[]\",\"name\":\"stats\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getInstituePublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"getUserPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"institutions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"publicAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"isApprovedInstitute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"registerAsUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"registerInstitution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"shaHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_institute\",\"type\":\"string\"},{\"internalType\":\"enumVerification.DocStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_proofHash\",\"type\":\"string\"}],\"name\":\"verifyDocument\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f600b553480156012575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611f058061005f5f395ff3fe60806040526004361061009b575f3560e01c8063ba1f3dc711610063578063ba1f3dc714610171578063ba785725146101ad578063c6b5c5cc146101d5578063e74c0342146101fd578063ea6509f814610239578063ef2d8700146102785761009b565b806311231fe01461009f578063188c71d9146100db5780634d2b1978146100f7578063515ef1061461011f578063b8b4da4c14610147575b5f5ffd5b3480156100aa575f5ffd5b506100c560048036038101906100c0919061115a565b6102a6565b6040516100d291906111f5565b60405180910390f35b6100f560048036038101906100f09190611364565b610375565b005b348015610102575f5ffd5b5061011d6004803603810190610118919061141c565b610583565b005b34801561012a575f5ffd5b50610145600480360381019061014091906114ef565b6107cf565b005b348015610152575f5ffd5b5061015b6108ce565b6040516101689190611552565b60405180910390f35b34801561017c575f5ffd5b506101976004803603810190610192919061156b565b6108d7565b6040516101a491906111f5565b60405180910390f35b3480156101b8575f5ffd5b506101d360048036038101906101ce919061156b565b610988565b005b3480156101e0575f5ffd5b506101fb60048036038101906101f6919061141c565b610a51565b005b348015610208575f5ffd5b50610223600480360381019061021e919061156b565b610beb565b60405161023091906115cc565b60405180910390f35b348015610244575f5ffd5b5061025f600480360381019061025a919061156b565b610c28565b60405161026f94939291906115f4565b60405180910390f35b348015610283575f5ffd5b5061028c610da3565b60405161029d95949392919061191a565b60405180910390f35b606060025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0180546102f2906119bb565b80601f016020809104026020016040519081016040528092919081815260200182805461031e906119bb565b80156103695780601f1061034057610100808354040283529160200191610369565b820191905f5260205f20905b81548152906001019060200180831161034c57829003601f168201915b50505050509050919050565b600115156001846040516103899190611a25565b90815260200160405180910390206003015f9054906101000a900460ff16151514801561042057503373ffffffffffffffffffffffffffffffffffffffff166001846040516103d89190611a25565b90815260200160405180910390205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610428575f5ffd5b5f6003856040516104399190611a25565b908152602001604051809103902054905082600a828154811061045f5761045e611a3b565b5b905f5260205f2090602091828204019190066101000a81548160ff0219169083600281111561049157610490611828565b5b02179055506003856040516104a69190611a25565b9081526020016040518091039020546003836040516104c59190611a25565b90815260200160405180910390208190555081600782815481106104ec576104eb611a3b565b5b905f5260205f200190816105009190611c08565b506003856040516105119190611a25565b90815260200160405180910390205f9055336006828154811061053757610536611a3b565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1690506001151581151514610617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060e90611d21565b60405180910390fd5b600b5460038460405161062a9190611a25565b908152602001604051809103902081905550600533908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060065f908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600882908060018154018082558091505060019003905f5260205f20015f90919091909150908161072d9190611c08565b50600a6002908060018154018082558091505060019003905f5260205f2090602091828204019190069091909190916101000a81548160ff0219169083600281111561077c5761077b611828565b5b0217905550600783908060018154018082558091505060019003905f5260205f20015f9091909190915090816107b29190611c08565b50600b5f8154809291906107c590611d6c565b9190505550505050565b604051806020016040528083838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081525060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f0190816108719190611c08565b50905050600160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505050565b5f600b54905090565b60606001826040516108e99190611a25565b90815260200160405180910390206001018054610905906119bb565b80601f0160208091040260200160405190810160405280929190818152602001828054610931906119bb565b801561097c5780601f106109535761010080835404028352916020019161097c565b820191905f5260205f20905b81548152906001019060200180831161095f57829003601f168201915b50505050509050919050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0d90611e23565b60405180910390fd5b60018082604051610a279190611a25565b90815260200160405180910390206003015f6101000a81548160ff02191690831515021790555050565b5f73ffffffffffffffffffffffffffffffffffffffff16600182604051610a789190611a25565b90815260200160405180910390205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610afe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af590611e8b565b60405180910390fd5b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020015f1515815250600182604051610b479190611a25565b90815260200160405180910390205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610bae9190611c08565b506040820151816002019081610bc49190611c08565b506060820151816003015f6101000a81548160ff0219169083151502179055509050505050565b5f60011515600183604051610c009190611a25565b90815260200160405180910390206003015f9054906101000a900460ff161515149050919050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610c84906119bb565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb0906119bb565b8015610cfb5780601f10610cd257610100808354040283529160200191610cfb565b820191905f5260205f20905b815481529060010190602001808311610cde57829003601f168201915b505050505090806002018054610d10906119bb565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3c906119bb565b8015610d875780601f10610d5e57610100808354040283529160200191610d87565b820191905f5260205f20905b815481529060010190602001808311610d6a57829003601f168201915b505050505090806003015f9054906101000a900460ff16905084565b60608060608060606005600660086007600a84805480602002602001604051908101604052809291908181526020018280548015610e3357602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610dea575b5050505050945083805480602002602001604051908101604052809291908181526020018280548015610eb857602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e6f575b5050505050935082805480602002602001604051908101604052809291908181526020015f905b82821015610f87578382905f5260205f20018054610efc906119bb565b80601f0160208091040260200160405190810160405280929190818152602001828054610f28906119bb565b8015610f735780601f10610f4a57610100808354040283529160200191610f73565b820191905f5260205f20905b815481529060010190602001808311610f5657829003601f168201915b505050505081526020019060010190610edf565b50505050925081805480602002602001604051908101604052809291908181526020015f905b82821015611055578382905f5260205f20018054610fca906119bb565b80601f0160208091040260200160405190810160405280929190818152602001828054610ff6906119bb565b80156110415780601f1061101857610100808354040283529160200191611041565b820191905f5260205f20905b81548152906001019060200180831161102457829003601f168201915b505050505081526020019060010190610fad565b505050509150808054806020026020016040519081016040528092919081815260200182805480156110d757602002820191905f5260205f20905f905b82829054906101000a900460ff1660028111156110b2576110b1611828565b5b815260200190600101906020825f010492830192600103820291508084116110925790505b50505050509050945094509450945094509091929394565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61112982611100565b9050919050565b6111398161111f565b8114611143575f5ffd5b50565b5f8135905061115481611130565b92915050565b5f6020828403121561116f5761116e6110f8565b5b5f61117c84828501611146565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6111c782611185565b6111d1818561118f565b93506111e181856020860161119f565b6111ea816111ad565b840191505092915050565b5f6020820190508181035f83015261120d81846111bd565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611253826111ad565b810181811067ffffffffffffffff821117156112725761127161121d565b5b80604052505050565b5f6112846110ef565b9050611290828261124a565b919050565b5f67ffffffffffffffff8211156112af576112ae61121d565b5b6112b8826111ad565b9050602081019050919050565b828183375f83830152505050565b5f6112e56112e084611295565b61127b565b90508281526020810184848401111561130157611300611219565b5b61130c8482856112c5565b509392505050565b5f82601f83011261132857611327611215565b5b81356113388482602086016112d3565b91505092915050565b6003811061134d575f5ffd5b50565b5f8135905061135e81611341565b92915050565b5f5f5f5f6080858703121561137c5761137b6110f8565b5b5f85013567ffffffffffffffff811115611399576113986110fc565b5b6113a587828801611314565b945050602085013567ffffffffffffffff8111156113c6576113c56110fc565b5b6113d287828801611314565b93505060406113e387828801611350565b925050606085013567ffffffffffffffff811115611404576114036110fc565b5b61141087828801611314565b91505092959194509250565b5f5f60408385031215611432576114316110f8565b5b5f83013567ffffffffffffffff81111561144f5761144e6110fc565b5b61145b85828601611314565b925050602083013567ffffffffffffffff81111561147c5761147b6110fc565b5b61148885828601611314565b9150509250929050565b5f5ffd5b5f5ffd5b5f5f83601f8401126114af576114ae611215565b5b8235905067ffffffffffffffff8111156114cc576114cb611492565b5b6020830191508360018202830111156114e8576114e7611496565b5b9250929050565b5f5f60208385031215611505576115046110f8565b5b5f83013567ffffffffffffffff811115611522576115216110fc565b5b61152e8582860161149a565b92509250509250929050565b5f819050919050565b61154c8161153a565b82525050565b5f6020820190506115655f830184611543565b92915050565b5f602082840312156115805761157f6110f8565b5b5f82013567ffffffffffffffff81111561159d5761159c6110fc565b5b6115a984828501611314565b91505092915050565b5f8115159050919050565b6115c6816115b2565b82525050565b5f6020820190506115df5f8301846115bd565b92915050565b6115ee8161111f565b82525050565b5f6080820190506116075f8301876115e5565b818103602083015261161981866111bd565b9050818103604083015261162d81856111bd565b905061163c60608301846115bd565b95945050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6116778161111f565b82525050565b5f611688838361166e565b60208301905092915050565b5f602082019050919050565b5f6116aa82611645565b6116b4818561164f565b93506116bf8361165f565b805f5b838110156116ef5781516116d6888261167d565b97506116e183611694565b9250506001810190506116c2565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f61173f82611185565b6117498185611725565b935061175981856020860161119f565b611762816111ad565b840191505092915050565b5f6117788383611735565b905092915050565b5f602082019050919050565b5f611796826116fc565b6117a08185611706565b9350836020820285016117b285611716565b805f5b858110156117ed57848403895281516117ce858261176d565b94506117d983611780565b925060208a019950506001810190506117b5565b50829750879550505050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6003811061186657611865611828565b5b50565b5f81905061187682611855565b919050565b5f61188582611869565b9050919050565b6118958161187b565b82525050565b5f6118a6838361188c565b60208301905092915050565b5f602082019050919050565b5f6118c8826117ff565b6118d28185611809565b93506118dd83611819565b805f5b8381101561190d5781516118f4888261189b565b97506118ff836118b2565b9250506001810190506118e0565b5085935050505092915050565b5f60a0820190508181035f83015261193281886116a0565b9050818103602083015261194681876116a0565b9050818103604083015261195a818661178c565b9050818103606083015261196e818561178c565b9050818103608083015261198281846118be565b90509695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806119d257607f821691505b6020821081036119e5576119e461198e565b5b50919050565b5f81905092915050565b5f6119ff82611185565b611a0981856119eb565b9350611a1981856020860161119f565b80840191505092915050565b5f611a3082846119f5565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611ac47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611a89565b611ace8683611a89565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611b09611b04611aff8461153a565b611ae6565b61153a565b9050919050565b5f819050919050565b611b2283611aef565b611b36611b2e82611b10565b848454611a95565b825550505050565b5f5f905090565b611b4d611b3e565b611b58818484611b19565b505050565b5b81811015611b7b57611b705f82611b45565b600181019050611b5e565b5050565b601f821115611bc057611b9181611a68565b611b9a84611a7a565b81016020851015611ba9578190505b611bbd611bb585611a7a565b830182611b5d565b50505b505050565b5f82821c905092915050565b5f611be05f1984600802611bc5565b1980831691505092915050565b5f611bf88383611bd1565b9150826002028217905092915050565b611c1182611185565b67ffffffffffffffff811115611c2a57611c2961121d565b5b611c3482546119bb565b611c3f828285611b7f565b5f60209050601f831160018114611c70575f8415611c5e578287015190505b611c688582611bed565b865550611ccf565b601f198416611c7e86611a68565b5f5b82811015611ca557848901518255600182019150602085019450602081019050611c80565b86831015611cc25784890151611cbe601f891682611bd1565b8355505b6001600288020188555050505b505050505050565b7f726567697374657220666972737420746f2076657269667900000000000000005f82015250565b5f611d0b60188361118f565b9150611d1682611cd7565b602082019050919050565b5f6020820190508181035f830152611d3881611cff565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611d768261153a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611da857611da7611d3f565b5b600182019050919050565b7f4f6e6c792061646d696e2063616e20706572666f6d207468697320616374696f5f8201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f611e0d60218361118f565b9150611e1882611db3565b604082019050919050565b5f6020820190508181035f830152611e3a81611e01565b9050919050565b7f496e737469747574696f6e20616c7265616479207265676973746572656400005f82015250565b5f611e75601e8361118f565b9150611e8082611e41565b602082019050919050565b5f6020820190508181035f830152611ea281611e69565b905091905056fea2646970667358221220476ec5360fe105261c62a9e2929ee3ee6f1aa9f8844529e2de3fb2ddf731302864736f6c637828302e382e33302d646576656c6f702e323032352e372e32352b636f6d6d69742e37333731326130310059",
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
