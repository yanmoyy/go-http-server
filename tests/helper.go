package tests

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/yanmoyy/go-http-server/internal/api"
)

func getClient(t *testing.T) *api.Client {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		t.Fatal("BASE_URL not set in .env or environment")
	}
	client := api.NewClient(baseURL, 5*time.Second)
	return client
}

// logJSON pretty-prints a struct as JSON with indentation
func logJSON(t *testing.T, prefix string, v any) {
	t.Helper()
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Errorf("Failed to marshal %s to JSON: %v", prefix, err)
		return
	}
	t.Logf("%s:\n%s", prefix, string(data))
}

func runReset(t *testing.T, client *api.Client) {
	t.Run("Reset", func(t *testing.T) {
		if err := client.Reset(); err != nil {
			t.Fatalf("Failed to reset: %v", err)
		}
		t.Log("System reset successfully")
	})
}

func runCreateUser(t *testing.T, client *api.Client, email, password string) api.User {
	var user api.User
	t.Run("CreateUser", func(t *testing.T) {
		var err error
		user, err = client.CreateUser(api.CreateUserParams{
			Email: email, Password: password,
		})
		if err != nil {
			t.Fatalf("Failed to create user: %v", err)
		}
		logJSON(t, "Created user", user)
	})
	return user
}
func runCreateUserDefault(t *testing.T, c *api.Client) api.User {
	return runCreateUser(t, c, "saul@bettercall.com", "123456")
}

func runCreateChirp(t *testing.T, client *api.Client, user api.User, body string) api.Chirp {
	var chirp api.Chirp
	t.Run("CreateChirp", func(t *testing.T) {
		params := api.CreateChirpParams{
			Body:   body,
			UserID: user.ID,
		}
		var err error
		chirp, err = client.CreateChirp(params)
		if err != nil {
			t.Fatalf("Failed to create chirp: %v", err)
		}
		logJSON(t, "Created chirp", chirp)
	})
	return chirp
}
