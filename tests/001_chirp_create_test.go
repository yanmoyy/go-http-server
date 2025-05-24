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

	user := runCreateUserDefault(t, client)
	if t.Failed() {
		return
	}

	runCreateChirp(t, client, user, "I'm gonna be a damn good developer, and people are gonna know about it.")
}
