package tests

import (
	"testing"

	"github.com/google/uuid"
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

func runGetChirpByID(t *testing.T, client *api.Client, chirpID uuid.UUID) api.Chirp {
	var chirp api.Chirp
	t.Run("GetChirpByID", func(t *testing.T) {
		fetchedChirp, err := client.GetChirpByID(chirpID)
		if err != nil {
			t.Fatalf("Failed to fetch chirp: %v", err)
		}
		logJSON(t, "Fetched chirp", fetchedChirp)
		// Verify the fetched chirp
		if fetchedChirp.ID != chirpID {
			t.Errorf("Expected chirp ID %v, got %v", chirpID, fetchedChirp.ID)
		}
		chirp = fetchedChirp
	})
	return chirp
}
