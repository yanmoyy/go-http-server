package tests

import (
	"testing"
)

func TestGetChirp(t *testing.T) {
	client := getClient(t)
	runReset(t, client)
	if t.Failed() {
		return
	}

	runCreateUserDefault(t, client)
	if t.Failed() {
		return
	}
	resp := runLoginUserDefault(t, client)
	if t.Failed() {
		return
	}

	chirp := runCreateChirp(t, client, resp.Token, "I'm gonna be a damn good developer, and people are gonna know about it.")
	if t.Failed() {
		return
	}

	t.Run("GetChirpByID", func(t *testing.T) {
		fetchedChirp, err := client.GetChirpByID(chirp.ID)
		if err != nil {
			t.Fatalf("Failed to fetch chirp: %v", err)
		}
		logJSON(t, "Fetched chirp", fetchedChirp)
		// Verify the fetched chirp
		if fetchedChirp.ID != chirp.ID {
			t.Errorf("Expected chirp ID %v, got %v", chirp.ID, fetchedChirp.ID)
		}
		if fetchedChirp.Body != chirp.Body {
			t.Errorf("Expected chirp body %q, got %q", chirp.Body, fetchedChirp.Body)
		}
	})
}
