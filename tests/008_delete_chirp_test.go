package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteChirp(t *testing.T) {
	c := getClient(t)
	runReset(t, c)
	if t.Failed() {
		return
	}

	email := "walt@breakingbad.com"
	password := "123456"
	user := runCreateUser(t, c, email, password)
	if t.Failed() {
		return
	}
	assert.Equal(t, email, user.Email)
	resp := runLoginUser(t, c, email, password)
	if t.Failed() {
		return
	}

	waltToken := resp.Token

	chirp := runCreateChirp(t, c, waltToken, "I did it for me. I liked it. I was good at it. And I was really... I was alive.")
	if t.Failed() {
		return
	}

	chirpId := chirp.ID
	chirp = runGetChirpByID(t, c, chirpId)
	if t.Failed() {
		return
	}
	// with no token
	statusCode, err := c.DeleteChirpByID(chirpId, "")
	if err == nil || statusCode != http.StatusUnauthorized {
		t.Errorf("Should be failed and Unauthorized(401): %v", statusCode)
	}

	runCreateUserDefault(t, c)
	if t.Failed() {
		return
	}
	resp = runLoginUserDefault(t, c)
	if t.Failed() {
		return
	}
	saulToken := resp.Token

	statusCode, err = c.DeleteChirpByID(chirpId, saulToken)
	if err == nil || statusCode != http.StatusForbidden {
		t.Errorf("Should be failed and Forbidden(403): %v", statusCode)
	}
	statusCode, err = c.DeleteChirpByID(chirpId, waltToken)
	if err != nil || statusCode != http.StatusNoContent {
		t.Errorf("Failed to Delete Chirp: error= %v code(204)= %v", err, statusCode)
	}
	_, err = c.GetChirpByID(chirpId)
	if err == nil {
		t.Errorf("Should be failed and 404")
	}
}
