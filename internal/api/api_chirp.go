package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

type CreateChirpParams struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

func (c *Client) CreateChirp(params CreateChirpParams) (Chirp, error) {
	resp, err := c.post(ChirpsEndpoint, params)
	if err != nil {
		return Chirp{}, fmt.Errorf("c.post: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusCreated {
		return Chirp{}, fmt.Errorf("status code is not StatusCreated(201): %d", resp.StatusCode)
	}
	var user Chirp
	err = c.decodeResponse(resp, &user)
	if err != nil {
		return Chirp{}, fmt.Errorf("c.decodeResponse: %w", err)
	}
	return user, nil
}

func (c *Client) GetChirpList() ([]Chirp, error) {
	resp, err := c.get(ChirpsEndpoint)
	if err != nil {
		return nil, fmt.Errorf("c.get: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is not StatusOK(200): %d", resp.StatusCode)
	}

	var list []Chirp
	err = c.decodeResponse(resp, &list)
	if err != nil {
		return nil, fmt.Errorf("c.decodeRespons: %w", err)
	}
	return list, nil
}

func (c *Client) GetChirpByID(id uuid.UUID) (Chirp, error) {
	resp, err := c.get(ChirpsEndpoint + "/" + id.String())
	if err != nil {
		return Chirp{}, fmt.Errorf("c.get: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return Chirp{}, fmt.Errorf("status code is not StatusOK(200): %d", resp.StatusCode)
	}
	chirp := Chirp{}
	err = c.decodeResponse(resp, &chirp)
	if err != nil {
		return Chirp{}, fmt.Errorf("c.decodeRespons: %w", err)
	}
	return chirp, nil
}
