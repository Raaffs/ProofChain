package blockchain

import "github.com/ethereum/go-ethereum/ethclient"

const (
	Approved=iota
	Rejected
	Pending
)

type Connect interface{
	New(string)error
	SetClient(*ethclient.Client)
}

//GetDocument method in Verify.go returns an anonymous struct with corresponding
//fields.
type VerificationDocument struct{
	//id field is required in frontend for data row element
	ID			int
	Requester   string
	Verifier    string
	Institute	string
	ShaHash		string
	Stats       uint8
}