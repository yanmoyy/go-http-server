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
