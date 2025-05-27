package api

import (
	"fmt"
	"net/http"
)

func (c *Client) Reset() error {
	resp, err := c.post(EndpointReset, "", nil)
	if err != nil {
		return fmt.Errorf("c.post: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not StatusOK(200): %d", resp.StatusCode)
	}
	return nil
}
