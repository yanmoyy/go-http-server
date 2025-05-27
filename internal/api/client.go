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

type Queries map[string]string

func (c *Client) doRequest(method, endpoint, auth string, queries Queries, body []byte) (*http.Response, error) {
	url := c.baseURL + endpoint

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	query := req.URL.Query()
	for k, v := range queries {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set(HeaderContentType, HeaderApplicationJson)
	}
	if auth != "" {
		req.Header.Add(HeaderAuthorization, auth)
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
	return c.doRequest(http.MethodGet, endpoint, "", nil, nil)
}

func (c *Client) post(endpoint string, auth string, body any) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	return c.doRequest(http.MethodPost, endpoint, auth, nil, jsonData)
}

func (c *Client) deleteWithToken(endpoint, token string) (*http.Response, error) {
	return c.doRequest(http.MethodDelete, endpoint, BearerPrefix+token, nil, nil)
}

func (c *Client) putWithToken(endpoint, token string, body any) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	return c.doRequest(http.MethodPut, endpoint, BearerPrefix+token, nil, jsonData)
}

func (c *Client) postWithToken(endpoint, token string, body any) (*http.Response, error) {
	return c.post(endpoint, BearerPrefix+token, body)
}

func (c *Client) postWithAPIKey(endpoint, key string, body any) (*http.Response, error) {
	return c.post(endpoint, ApiKeyPrefix+key, body)
}

func (c *Client) getWithQuery(endpoint string, queries Queries) (*http.Response, error) {
	return c.doRequest(http.MethodGet, endpoint, "", queries, nil)
}
