package api

import (
	"fmt"
	"net/http"
)

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

func (c *Client) RefreshToken(refreshToken string) (RefreshTokenResponse, error) {
	resp, err := c.postWithToken(EndpointRefresh, refreshToken, nil)
	if err != nil {
		return RefreshTokenResponse{}, fmt.Errorf("c.postWithToken: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return RefreshTokenResponse{}, fmt.Errorf("status code is not OK(200): %d", resp.StatusCode)
	}
	var refreshResp RefreshTokenResponse
	err = c.decodeResponse(resp, &refreshResp)
	if err != nil {
		return RefreshTokenResponse{}, fmt.Errorf("c.decodeResponse: %w", err)
	}
	return refreshResp, nil
}

func (c *Client) RevokeToken(refreshToken string) error {
	resp, err := c.postWithToken(EndpointRevoke, refreshToken, nil)
	if err != nil {
		return fmt.Errorf("c.postWithToken: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("status code is not NoContent(204): %d", resp.StatusCode)
	}
	return nil
}
