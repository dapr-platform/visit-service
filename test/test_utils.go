package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const baseURL = "http://182.92.117.41:37080/swagger/visit-service" // 替换为实际的测试服务器地址

type TestClient struct {
	BaseURL string
	Client  *http.Client
}

func NewTestClient() *TestClient {
	return &TestClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (c *TestClient) DoRequest(t *testing.T, method, path string, body interface{}) (*http.Response, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			t.Fatalf("Failed to marshal request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return c.Client.Do(req)
}
