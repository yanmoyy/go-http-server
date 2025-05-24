package api

import (
	"fmt"
	"net/http"
)

type LoginParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (c *Client) Login(params LoginParams) (User, error) {
	resp, err := c.post(LoginEndpoint, params)
	if err != nil {
		return User{}, fmt.Errorf("c.post: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("status code is not OK(200): %d", resp.StatusCode)
	}
	var user User
	err = c.decodeResponse(resp, &user)
	if err != nil {
		return User{}, fmt.Errorf("c.decodeResponse: %w", err)
	}
	return user, nil
}
