package api

import (
	"fmt"

	"github.com/google/uuid"
)

type Event string

// Webhooks Event
const EventUpgrade = Event("user.upgraded")

type PolkaWebhookParams struct {
	Event Event `json:"event"`
	Data  struct {
		UserID uuid.UUID `json:"user_id"`
	} `json:"data"`
}

func (c *Client) PolkaWebhookPost(userID uuid.UUID, event Event, key string) (statusCode int, err error) {
	params := PolkaWebhookParams{
		Event: event,
		Data: struct {
			UserID uuid.UUID `json:"user_id"`
		}{
			UserID: userID,
		},
	}
	resp, err := c.postWithAPIKey(EndpointPolkaWebhooks, key, params)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("c.post: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return resp.StatusCode, fmt.Errorf("status code is not 200~299: %d", resp.StatusCode)
	}
	return resp.StatusCode, nil
}
