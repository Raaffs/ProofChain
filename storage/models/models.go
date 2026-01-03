package models

type Document struct {
	Shahash           string `bson:"shahash" json:"shahash"`
	EncryptedDocument []byte `bson:"encryptedDocument" json:"encryptedDocument"`
	PublicAddress     string `bson:"publicAddress" json:"publicAddress"`
}

type CertificateData struct {
    CertificateName string     			`json:"certificateName" bson:"certificateName"`
    PublicAddress   string     			`json:"publicAddress" bson:"publicAddress"`
    Name            string     			`json:"name" bson:"name"`
    Address         string     			`json:"address" bson:"address"`
    Age             string 				`json:"age" bson:"age"` // number or string
    BirthDate       string     			`json:"birthDate" bson:"birthDate"`
    UniqueID        string     			`json:"uniqueId" bson:"uniqueId"`
    Extra           map[string]string	`json:"-" bson:",inline"` // inline extra props
}
