package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/go-http-server/internal/api"
)

func TestPolkaWebhook(t *testing.T) {
	key := os.Getenv("POLKA_KEY")
	if key == "" {
		t.Errorf("POLKA_KEY not found")
	}

	c := getClient(t)
	runReset(t, c)
	if t.Failed() {
		return
	}

	user := runCreateUserDefault(t, c)
	if t.Failed() {
		return
	}
	assert.Equal(t, defaultEmail, user.Email)
	assert.Equal(t, false, user.IsChirpyRed)
	code, err := c.PolkaWebhookPost(user.ID, "user.payment_failed", key)
	if err != nil || code != http.StatusNoContent {
		t.Errorf("Failed to PolkaWebhookPost: %v, code=%v", err, code)
	}
	code, err = c.PolkaWebhookPost(user.ID, api.EventUpgrade, key)
	if err != nil || code != http.StatusNoContent {
		t.Errorf("Failed to PolkaWebhookPost: %v, code=%v", err, code)
	}
	resp := runLoginUserDefault(t, c)
	if t.Failed() {
		return
	}
	assert.Equal(t, true, resp.IsChirpyRed)
}
