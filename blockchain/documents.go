package blockchain

import (
	"github.com/Suy56/ProofChain/verify"
	"github.com/ethereum/go-ethereum/common"
)
func GetDocuments(documents []verify.VerificationDocument,condition func(verify.VerificationDocument,common.Address)bool, requester common.Address)[]verify.VerificationDocument{
	var filteredDocs []verify.VerificationDocument
	for _,document :=range documents{
		if condition(document,requester){
			filteredDocs=append(filteredDocs, document)
		}
	}
	return filteredDocs
}