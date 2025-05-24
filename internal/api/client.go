package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Client represents the API client
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new API client with the specified base URL and timeout
func NewClient(baseURL string, timeout time.Duration) Client {
	return Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) doRequest(method, endpoint string, body []byte) (*http.Response, error) {
	url := c.baseURL + endpoint

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}
	if resp.StatusCode >= 300 {
		bodyByte, err := io.ReadAll(resp.Body)
		defer func() { _ = resp.Body.Close() }()
		if err != nil {
			return nil, fmt.Errorf("io.ReadAll(resp.Body): %w", err)
		}
		log.Printf("Status Wrong\nResponse: %s\n", bodyByte)
		resp.Body = io.NopCloser(bytes.NewReader(bodyByte))
	}

	return resp, nil
}

func (c *Client) decodeResponse(resp *http.Response, v any) error {
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return err
	}
	return nil
}

func (c *Client) get(endpoint string) (*http.Response, error) {
	return c.doRequest(http.MethodGet, endpoint, nil)
}

func (c *Client) post(endpoint string, body any) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	return c.doRequest(http.MethodPost, endpoint, jsonData)
}

func (c *Client) put(endpoint string, body any) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	return c.doRequest(http.MethodPut, endpoint, jsonData)
}

func (c *Client) delete(endpoint string) (*http.Response, error) {
	return c.doRequest(http.MethodDelete, endpoint, nil)
}
