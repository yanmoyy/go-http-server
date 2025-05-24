package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	os.Exit(m.Run())
}

func TestCreateChirp(t *testing.T) {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		t.Fatal("BASE_URL not set in .env or environment")
	}
	makePostRequest := func(url, contentType string, body []byte) (*http.Response, error) {
		resp, err := http.Post(url, contentType, bytes.NewReader(body))
		if err != nil {
			return nil, fmt.Errorf("failed to send POST request to %s: %v", url, err)
		}
		if resp.StatusCode > 299 {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to read response body: %v", err)
			}
			return nil, fmt.Errorf("Response Error (%d): %s", resp.StatusCode, string(bodyBytes))
		}
		return resp, nil
	}
	decodeResponse := func(resp *http.Response, v any) error {
		defer func() {
			_ = resp.Body.Close()
		}()
		return json.NewDecoder(resp.Body).Decode(v)
	}
	// Reset the system
	if _, err := makePostRequest(baseURL+"/admin/reset", "application/json", nil); err != nil {
		t.Fatalf("Failed to reset: %v", err)
	}

	// Create user
	userReq := struct {
		Email string `json:"email"`
	}{
		Email: "saul@bettercall.com",
	}

	jsonData, err := json.Marshal(userReq)
	if err != nil {
		t.Fatalf("Failed to marshal user JSON: %v", err)
	}
	resp, err := makePostRequest(baseURL+"/api/users", "application/json", jsonData)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	var user User
	if err := decodeResponse(resp, &user); err != nil {
		t.Fatalf("Failed to decode user response: %v", err)
	}
	chirpReq := struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}{
		Body:   "If you're committed enough, you can make any story work.",
		UserID: user.ID,
	}
	jsonData, err = json.Marshal(chirpReq)
	if err != nil {
		t.Fatalf("Failed to marshal chirp JSON: %v", err)
	}
	resp, err = makePostRequest(baseURL+"/api/chirps", "application/json", jsonData)
	if err != nil {
		t.Fatalf("Failed to create chirp: %v", err)
	}
	var chirp Chirp
	if err := decodeResponse(resp, &chirp); err != nil {
		t.Fatalf("Failed to decode chirp response: %v", err)
	}
	t.Logf("Created Chirp: %+v", chirp)
}
