package nodeData

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

type IPFSManager struct{
	shell 		*shell.Shell
	hostPort 	string
}

func(i *IPFSManager)New(hostPort string){
	i.hostPort=fmt.Sprintf("localhost:%s",hostPort)
	i.shell=shell.NewShell(i.hostPort)
}

func(i *IPFSManager) Upload(filePath string)(string,error){
	file, err := os.Open(filePath);if err != nil {
		return "",err
	}
    defer file.Close()
	
	cid,err:=i.shell.Add(file); if err!=nil{
		return "",err
	}
	return cid,nil
}

func RetrieveLink(cid string)string{
	return 	fmt.Sprintf("https://ipfs.io/ipfs/%s",cid)
}