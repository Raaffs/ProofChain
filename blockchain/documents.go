package blockchain

import "github.com/ethereum/go-ethereum/common"


//needs to be tested
type  Address interface{
	~string|common.Address
}

func FilterDocument[T Address](docs []VerificationDocument, condition func(VerificationDocument,T)bool, requester T)[]VerificationDocument{
	var userDocs []VerificationDocument
	for _,doc :=range docs{
			if(condition(doc,requester)){
					userDocs=append(userDocs, doc)
			}
	}
	return userDocs
}

