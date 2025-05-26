package api

import (
	"fmt"
	"net/http"
)

type LoginParams struct {
	Password         string `json:"password"`
	Email            string `json:"email"`
	ExpiresInSeconds int64  `json:"expires_in_seconds"`
}

type LoginResponse struct {
	User
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *Client) Login(params LoginParams) (LoginResponse, error) {
	resp, err := c.post(EndpointLogin, params)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("c.post: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return LoginResponse{}, fmt.Errorf("status code is not OK(200): %d", resp.StatusCode)
	}
	var loginResp LoginResponse
	err = c.decodeResponse(resp, &loginResp)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("c.decodeResponse: %w", err)
	}
	return loginResp, nil
}
