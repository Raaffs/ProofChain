package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

type hashedField struct {
	Hash  string `json:"hash"`
	Key   string `json:"key"`
	Salt  string `json:"salt"`
	Value string `json:"value"`
}

type DownloadProof struct {
	Address         hashedField            			`json:"Address"`
	Age             hashedField            			`json:"Age"`
	BirthDate       hashedField            			`json:"BirthDate"`
	CertificateName hashedField            			`json:"CertificateName"`
	Name            hashedField            			`json:"Name"`
	PublicAddress   hashedField            			`json:"PublicAddress"`
	UniqueID        hashedField            			`json:"UniqueID"`
	Extra           map[string]hashedField 			`json:"Extra"`
}

func Download(document []byte) error {
	var doc DownloadProof
	path, err := getDownloadDir()
	if err != nil {
		return err
	}
	fmt.Println(path)
	// Use a helper to ensure Extra map is populated from JSON
	if err := json.Unmarshal(document, &doc); err != nil {
		return err
	}
	fmt.Println("doc: ",doc)
	// Manual step if your JSON has extra fields and you aren't using a custom unmarshaler:

	for k, v := range Walk(doc) {
		// Pass 'v' (the full field) into the extractor
		proof_k := extractProofValues(doc, k, v)
		
		if  k=="Name" || k=="BirthDate" || k=="Address" || k=="Extra"{
			fmt.Printf("\n--- Generated Proof for Key: %s ---\n", k)
			fmt.Printf("%v\n%v\n%v\n%v\n", proof_k.Name,proof_k.BirthDate,proof_k.Address, proof_k.Extra,)
		}
	}
	return nil
}

func extractProofValues(v DownloadProof, activeKey string, fullValue any) DownloadProof {
	slim := func(f hashedField) hashedField {
		return hashedField{Hash: f.Hash, Key: f.Key}
	}

	// 1. Create the template where EVERYTHING is slimmed
	result := DownloadProof{
		Address:         slim(v.Address),
		Age:             slim(v.Age),
		BirthDate:       slim(v.BirthDate),
		CertificateName: slim(v.CertificateName),
		Name:            slim(v.Name),
		PublicAddress:   slim(v.PublicAddress),
		UniqueID:        slim(v.UniqueID),
		Extra:           make(map[string]hashedField),
	}

	// 2. Slim down the Extra map fields
	for k, val:=range v.Extra{
		result.Extra[k]=slim(val)
	}
	// 3. RESTORE the full value for the active key
	// We use reflection to find the field by string name
	rv := reflect.ValueOf(&result).Elem()
	field := rv.FieldByName(activeKey)

	if field.IsValid() && field.CanSet() {
		// The key matches a struct field (e.g., "BirthDate")
		field.Set(reflect.ValueOf(fullValue))
	} else {
		// The key belongs in the Extra map
		result.Extra[activeKey] = fullValue.(hashedField)
	}

	return result
}

// Helper to grab fields not defined in the struct from the raw JSON
// func extractExtraFields(data []byte) map[string]hashedField {
// 	var raw map[string]hashedField
// 	json.Unmarshal(data, &raw)
// 	fixed := map[string]bool{
// 		"Address": true, "Age": true, "BirthDate": true, 
// 		"CertificateName": true, "Name": true, "PublicAddress": true, "UniqueID": true,
// 	}
// 	extra := make(map[string]hashedField)
// 	for k, v := range raw {
// 		if !fixed[k] {
// 			fmt.Println("k,v",k,v)
// 			// If it looks like a hashedField map, convert it
// 				if _, hasHash := m["Hash"]; hasHash {
// 					extra[k] = hashedField{
// 						Hash:  fmt.Sprint(v["Hash"]),
// 						Key:   fmt.Sprint(v["Key"]),
// 						Salt:  fmt.Sprint(v["Salt"]),
// 						Value: fmt.Sprint(v["Value"]),
// 					}
// 					continue
// 				}
			
// 			extra[k] = v
// 		}
// 	}
// 	return extra
// }

func getDownloadDir()(string,error){
	var downloadDir string

	// 1. Try to get the localized/configured path via xdg-user-dir (Linux standard)
	cmd := exec.Command("xdg-user-dir", "DOWNLOAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err == nil {
		downloadDir = strings.TrimSpace(out.String())
	}

	// 2. Fallback for macOS or Linux systems without xdg-user-dir
	if downloadDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting download dir %w",err)
		}
		downloadDir = filepath.Join(home, "Downloads")
	}

	// 3. Create your app-specific subfolder
	finalPath := filepath.Join(downloadDir, "ProofChain")
	
	// Perm 0755: Owner can Read/Write/Execute, others can Read/Execute
	if err := os.MkdirAll(finalPath, 0755); err != nil {
		return "", err
	}

	return finalPath, nil
}