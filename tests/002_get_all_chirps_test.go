package tests

import (
	"os"
	"testing"
	"time"

	"github.com/yanmoyy/go-http-server/internal/api"
)

func TestGetAllChirps(t *testing.T) {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		t.Fatal("BASE_URL not set in .env or environment")
	}

	client := api.NewClient(baseURL, 5*time.Second)
	// Reset the system
	if err := client.Reset(); err != nil {
		t.Fatalf("Failed to reset: %v", err)
	}
	t.Log("System reset successfully")

	// Create a user
	user, err := client.CreateUser(api.CreateUserParams{Email: "saul@bettercall.com"})
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	t.Logf("Created user: %+v", user)

	// Create a chirp #1
	params := api.CreateChirpParams{
		Body:   "If you're committed enough, you can make any story work.",
		UserID: user.ID,
	}
	chirp, err := client.CreateChirp(params)
	if err != nil {
		t.Fatalf("Failed to create chirp: %v", err)
	}
	t.Logf("Created chirp: %+v", chirp)

	// Create a chirp #2
	params = api.CreateChirpParams{
		Body:   "I once told a woman I was Kevin Costner, and it worked because I believed it.",
		UserID: user.ID,
	}
	chirp, err = client.CreateChirp(params)
	if err != nil {
		t.Fatalf("Failed to create chirp: %v", err)
	}
	t.Logf("Created chirp: %+v", chirp)
	list, err := client.GetChirpList()
	if err != nil {
		t.Fatalf("Failed to get chirp list: %v", err)
	}
	t.Log("List of chirp: ")
	for i, c := range list {
		t.Logf("Chirp #%d\n", i)
		t.Log("id: ", c.ID)
		t.Log("created_at: ", c.CreatedAt)
		t.Log("updated_at: ", c.UpdatedAt)
		t.Log("body: ", c.Body)
		t.Log("user_id: ", c.UserID)
	}
}
