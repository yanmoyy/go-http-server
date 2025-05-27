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
	runGetChirpByID(t, client, chirp.ID)
	if t.Failed() {
		return
	}
}
