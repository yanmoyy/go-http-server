package api

import "fmt"

func (c *Client) Reset() error {
	resp, err := c.post(ResetEndpoint, nil)
	defer func() { _ = resp.Body.Close() }()
	if err != nil {
		return fmt.Errorf("c.post: %w", err)
	}
	return nil
}
