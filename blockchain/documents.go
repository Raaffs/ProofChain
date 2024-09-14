package blockchain

type DisplayDocument struct{
	ID			int
	Requester   string
	Verifier    string
	Institute	string
	Name        string
	Desc        string
	//ipfsAddress is stored as []byte on ethereum for proper mapping, so need to covnert
	//it to a string
	IpfsAddress string
	Stats       uint8
}
func FilterDocument(docs []VerificationDocument, condition func(VerificationDocument)bool)[]VerificationDocument{
	var userDocs []VerificationDocument
	for _,doc :=range docs{
		if(condition(doc)){	
			userDocs=append(userDocs,doc)
		}
	}
	return userDocs
}
