package tests

import (
	"fmt"
	"testing"
)

func TestGetAllChirps(t *testing.T) {
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

	runCreateChirp(t, client, resp.Token, "I'm gonna be a damn good developer, and people are gonna know about it.")
	if t.Failed() {
		return
	}

	runCreateChirp(t, client, resp.Token, "I once told a woman I was Kevin Costner, and it worked because I believed it.")
	if t.Failed() {
		return
	}

	list, err := client.GetChirpList("")
	if err != nil {
		t.Fatalf("Failed to get chirp list: %v", err)
	}
	t.Log("List of chirp: ")
	for i, c := range list {
		logJSON(t, fmt.Sprintf("Chirp #%d", i), c)
	}
}
