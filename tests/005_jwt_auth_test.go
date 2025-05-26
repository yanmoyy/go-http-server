package tests

import (
	"testing"
)

func TestAuthWithJWT(t *testing.T) {
	client := getClient(t)
	runReset(t, client)
	if t.Failed() {
		return
	}
	saulEmail := "saul@bettercall.com"
	saulPassword := "123456"
	runCreateUser(t, client, saulEmail, saulPassword)
	if t.Failed() {
		return
	}
	mikeEmail := "mike@bettercall.com"
	mikePassword := "98765"
	runCreateUser(t, client, mikeEmail, mikePassword)
	if t.Failed() {
		return
	}
	respSaul := runLoginUser(t, client, saulEmail, saulPassword)
	if t.Failed() {
		return
	}
	respMike := runLoginUser(t, client, mikeEmail, mikePassword)
	if t.Failed() {
		return
	}

	runCreateChirp(t, client, respSaul.Token, "Clearly his taste in women is the same as his taste in lawyers: only the very best... with just a right amount of dirty!")
	runCreateChirp(t, client, respMike.Token, "No more half-measures.")
}
