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
func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) doRequest(method, endpoint, token string, body []byte) (*http.Response, error) {
	url := c.baseURL + endpoint

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set(HeaderContentType, HeaderApplicationJson)
	}

	if token != "" {
		req.Header.Add(HeaderAuthorization, BearerPrefix+token)
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
		log.Printf("\n    Response(%d): %s", resp.StatusCode, bodyByte)
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
	return c.getWithToken(endpoint, "")
}

func (c *Client) post(endpoint string, body any) (*http.Response, error) {
	return c.postWithToken(endpoint, "", body)
}

func (c *Client) deleteWithToken(endpoint, token string) (*http.Response, error) {
	return c.doRequest(http.MethodDelete, endpoint, token, nil)
}

func (c *Client) putWithToken(endpoint, token string, body any) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	return c.doRequest(http.MethodPut, endpoint, token, jsonData)
}

func (c *Client) getWithToken(endpoint, token string) (*http.Response, error) {
	return c.doRequest(http.MethodGet, endpoint, token, nil)
}

func (c *Client) postWithToken(endpoint, token string, body any) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	return c.doRequest(http.MethodPost, endpoint, token, jsonData)
}
