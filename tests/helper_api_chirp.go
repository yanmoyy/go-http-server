package tests

import (
	"testing"

	"github.com/yanmoyy/go-http-server/internal/api"
)

func runCreateChirp(t *testing.T, client *api.Client, token string, body string) api.Chirp {
	var chirp api.Chirp
	t.Run("CreateChirp", func(t *testing.T) {
		params := api.CreateChirpParams{
			Body: body,
		}
		var err error
		chirp, err = client.CreateChirp(params, token)
		if err != nil {
			t.Errorf("Failed to create chirp: %v", err)
		}
		logJSON(t, "Created chirp", chirp)
	})
	return chirp
}
