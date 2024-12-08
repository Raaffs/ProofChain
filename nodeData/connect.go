package nodeData

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

)

// Connects to MongoDB and sets a Stable API version
type DocumentStore struct {
	Shahash           string `bson:"shahash" json:"shahash"`
	EncryptedDocument []byte `bson:"encryptedDocument" json:"encryptedDocument"`
	PublicAddress     string `bson:"publicAddress" json:"publicAddress"`
}

func UploadDocument(doc DocumentStore) error {
	jsonData, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to marshal document: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/add", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("bad request: please check the data")
	case http.StatusUnauthorized:
		return fmt.Errorf("unauthorized access")
	case http.StatusInternalServerError:
		return fmt.Errorf("server error: please try again later")
	default:
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func RetrieveDocument(sha string) (DocumentStore, error) {
	shaData := struct {
		Sha string `json:"shahash"`
	}{Sha: sha}
	log.Println("sha of document to retrieve : ",shaData.Sha)
	jsonData, err := json.Marshal(shaData)
	if err != nil {
		return DocumentStore{}, fmt.Errorf("failed to marshal sha data: %v", err)
	}
	var data DocumentStore
	resp, err := http.Post("http://localhost:8080/retrieve", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return DocumentStore{}, fmt.Errorf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return DocumentStore{}, fmt.Errorf("failed to retrieve document, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DocumentStore{}, fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return DocumentStore{}, fmt.Errorf("failed to decode response into Document: %v", err)
	}
	return data, nil
}
