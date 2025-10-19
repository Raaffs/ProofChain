package blockchain

func FilterDocument(docs []VerificationDocument, condition func(VerificationDocument)bool)[]VerificationDocument{
	var userDocs []VerificationDocument
	for _,doc :=range docs{
		if(condition(doc)){	
			userDocs=append(userDocs,doc)
		}
	}
	return userDocs
}
