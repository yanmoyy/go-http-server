package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	CreateAt    time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Email       string    `json:"email"`
	IsChirpyRed bool      `json:"is_chirpy_red"`
}

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Client) CreateUser(params CreateUserParams) (User, error) {
	resp, err := c.post(EndpointUsers, params)
	if err != nil {
		return User{}, fmt.Errorf("c.post: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusCreated {
		return User{}, fmt.Errorf("status code is not StatusCreated(201): %d", resp.StatusCode)
	}
	var user User
	err = c.decodeResponse(resp, &user)
	if err != nil {
		return User{}, fmt.Errorf("c.decodeResponse: %w", err)
	}
	return user, nil
}

type UpdateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Client) UpdateUser(params UpdateUserParams, token string) (User, error) {
	resp, err := c.putWithToken(EndpointUsers, token, params)
	if err != nil {
		return User{}, fmt.Errorf("c.put: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("status code is not StatusOK(200): %d", resp.StatusCode)
	}
	var user User
	err = c.decodeResponse(resp, &user)
	if err != nil {
		return User{}, fmt.Errorf("c.decodeResponse: %w", err)
	}
	return user, nil
}
