package tests

import (
	"testing"
)

func TestCreateChirp(t *testing.T) {
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
}
