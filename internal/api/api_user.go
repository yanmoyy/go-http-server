package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type CreateUserParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (c *Client) CreateUser(params CreateUserParams) (User, error) {
	resp, err := c.post(UsersEndpoint, params)
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
