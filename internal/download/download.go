package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/Suy56/ProofChain/internal/utils"
)

type hashedField struct {
	Hash  string `json:"hash"`
	Key   string `json:"key"`
	Salt  string `json:"salt"`
	Value string `json:"value"`
}

type DownloadProof struct {
	Address         hashedField            `json:"Address"`
	Age             hashedField            `json:"Age"`
	BirthDate       hashedField            `json:"BirthDate"`
	CertificateName hashedField            `json:"CertificateName"`
	Name            hashedField            `json:"Name"`
	PublicAddress   hashedField            `json:"PublicAddress"`
	UniqueID        hashedField            `json:"UniqueID"`
	Extra           map[string]hashedField `json:"Extra"`
}

func Download(document []byte) error {
	var doc DownloadProof
	path, err := getDownloadDir()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(document, &doc); err != nil {
		return err
	}
	for k, v := range utils.Walk(doc) {
		proof_k := extractProofValues(doc, k, v)
		dir:=filepath.Join(path,"ProofChain",doc.CertificateName.Value)
		if err:=store(k,dir,proof_k);err!=nil{
			log.Println(err)
		}
	}
	return nil
}

func store(key string, dir string, proof DownloadProof)error{
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(proof, "", "  ")
	if err != nil {
		return err
	}

	filename := filepath.Join(dir, key+".json")
	return os.WriteFile(filename, data, 0644)
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
	for k, val := range v.Extra {
		result.Extra[k] = slim(val)
	}

	// 3. RESTORE the full value for the active key
	// We use reflection to find the field by string name
	rv := reflect.ValueOf(&result).Elem()
	field := rv.FieldByName(activeKey)

	// First: assert fullValue is actually a hashedField
	hf, ok := fullValue.(hashedField)
	if !ok {
		// optional: log / return / silently ignore
		return result
	}

	if field.IsValid() && field.CanSet() {
		// Make sure the types actually match before Set
		if field.Type() == reflect.TypeOf(hf) {
			field.Set(reflect.ValueOf(hf))
		}
	} else {
		// Falls back to Extra map
		result.Extra[activeKey] = hf
	}
	return result
}


func getDownloadDir() (string, error) {
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
			return "", fmt.Errorf("error getting download dir %w", err)
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
