package main

import (
	"testing"

	blockchain "github.com/Suy56/ProofChain/chaincore/core"
)

func TestApprove(t *testing.T) {
	pk:="0x552d079fe5a76c34f075ddd694784a0f9a3c19aeadc62632fde406492073b48d"
	contract:="0xBBf03Cb4A6074844B40020acd6baAf4F436E9DCd"
	c:=	&blockchain.ClientConnection{}
	i:= &blockchain.ContractVerifyOperations{}
	host:="http://localhost:7545"
	if err:=blockchain.Init(
		c,
		i,
		pk[2:],
		contract,
		host,
	);err!=nil{
		t.Fatal(err)
	}

	if err:=i.ApproveVerifier(c.TxOpts,"ins");err!=nil{
		t.Fatal(err)
	}

}