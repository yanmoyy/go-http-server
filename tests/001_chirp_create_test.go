package tests

import (
	"testing"
)

func TestCreateChirp(t *testing.T) {
	client := getClient(t)
	runReset(t, client)
	user := runCreateUser(t, client)
	runCreateChirp(t, client, user, "I'm gonna be a damn good developer, and people are gonna know about it.")
}
