package nodeData

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func Upload(filePath string)(string,error){
	sh:=shell.NewShell("localhost:5001")
	file, err := os.Open(filePath);if err != nil {
		return "",err
	}
    defer file.Close()
	cid,err:=sh.Add(file); if err!=nil{
		return "",err
	}
	return cid,nil
}

func RetrieveLink(cid string)string{
	return 	fmt.Sprintf("https://ipfs.io/ipfs/%s",cid)
}