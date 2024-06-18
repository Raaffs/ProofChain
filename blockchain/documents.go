package blockchain

func FilterDocument[T comparable](docs []VerificationDocument, condition func(VerificationDocument,T)bool, requester T)[]VerificationDocument{
	var userDocs []VerificationDocument
	for _,doc :=range docs{
			if(condition(doc,requester)){
					userDocs=append(userDocs, doc)
			}
	}
	return userDocs
}


	