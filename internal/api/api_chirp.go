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
	Body string `json:"body"`
}

func (c *Client) CreateChirp(params CreateChirpParams, token string) (Chirp, error) {
	resp, err := c.postWithToken(EndpointChirps, token, params)
	if err != nil {
		return Chirp{}, fmt.Errorf("c.post: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusCreated {
		return Chirp{}, fmt.Errorf("status code is not StatusCreated(201): %d", resp.StatusCode)
	}
	var chirp Chirp
	err = c.decodeResponse(resp, &chirp)
	if err != nil {
		return Chirp{}, fmt.Errorf("c.decodeResponse: %w", err)
	}
	return chirp, nil
}

func (c *Client) GetChirpList() ([]Chirp, error) {
	resp, err := c.get(EndpointChirps)
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
	resp, err := c.get(EndpointChirps + "/" + id.String())
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

func (c *Client) DeleteChirpByID(chirpID uuid.UUID, token string) (statusCode int, err error) {
	resp, err := c.deleteWithToken(EndpointChirps+"/"+chirpID.String(), token)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("c.deleteWithToken: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusNoContent {
		return resp.StatusCode, fmt.Errorf("status code is not StatusNoContent(204): %d", resp.StatusCode)
	}
	return resp.StatusCode, nil
}
