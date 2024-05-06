package blockchain

import (
	// "github.com/Suy56/ProofChain/verify"
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