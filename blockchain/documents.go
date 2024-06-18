package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
)
func GetDocuments(documents []VerificationDocument,condition func(VerificationDocument,common.Address)bool, requester common.Address)[]VerificationDocument{
	var filteredDocs []VerificationDocument
	for _,document :=range documents{
		if condition(document,requester){
			filteredDocs=append(filteredDocs, document)
		}
	}
	return filteredDocs
}

func FilterDocument[T comparable](docs []VerificationDocument, condition func(VerificationDocument,T)bool, requester T)[]VerificationDocument{
	var userDocs []VerificationDocument
	for _,doc :=range docs{
			if(condition(doc,requester)){
					userDocs=append(userDocs, doc)
			}
	}
	return userDocs
}

	