package storageclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Suy56/ProofChain/storage/models"
)

type Client struct {
	BaseURL string
	Client  *http.Client
}

// New returns a new instance of the storage service client
func New(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

// UploadDocument uploads a document to the storage service
func (c *Client) UploadDocument(doc models.Document) error {
	return c.DoRequest("POST", "/add", doc, nil)
}

func (c *Client) RetrieveDocument(sha string) (models.Document, error) {
	payload := map[string]string{"shahash": sha}
	var result models.Document

	err := c.DoRequest("POST", "/retrieve", payload, &result)
	if err != nil {
		return models.Document{}, err
	}
	return result, nil
}


func (c *Client) GetApprovedInstitution() ([]string, error) {
	var result struct {
		Institution []string `bson:"institution" json:"institution"`
	}

	err := c.DoRequest("GET", "/retrieve/institution", nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Institution, nil
}	

func (c *Client) DoRequest(method, path string, body interface{}, out interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Centralized HTTP status code handling
	if err := handleHTTPError(resp); err != nil {
		return err
	}

	// If no response expected, return early
	if out == nil {
		return nil
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if err := json.Unmarshal(respData, out); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// handleHTTPError maps HTTP response codes to human-friendly errors
func handleHTTPError(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("bad request: please check your data")
	case http.StatusUnauthorized:
		return fmt.Errorf("unauthorized: invalid or missing credentials")
	case http.StatusForbidden:
		return fmt.Errorf("forbidden: you don’t have permission for this resource")
	case http.StatusNotFound:
		return fmt.Errorf("resource not found")
	case http.StatusInternalServerError:
		return fmt.Errorf("server error: please try again later")
	case http.StatusServiceUnavailable:
		return fmt.Errorf("service unavailable: storage service may be down")
	default:
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected response (%d): %s", resp.StatusCode, string(body))
	}
}
