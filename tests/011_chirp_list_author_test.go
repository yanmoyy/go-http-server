package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChirpListAuthorTest(t *testing.T) {
	c := getClient(t)
	if t.Failed() {
		return
	}

	runReset(t, c)
	if t.Failed() {
		return
	}

	saul := runCreateUserDefault(t, c)
	if t.Failed() {
		return
	}

	resp := runLoginUserDefault(t, c)
	if t.Failed() {
		return
	}
	saulAccessToken := resp.Token

	saulBody1 := "I'm the one who knocks!"
	saulBody2 := "Gale!"
	runCreateChirp(t, c, saulAccessToken, saulBody1)
	if t.Failed() {
		return
	}
	runCreateChirp(t, c, saulAccessToken, saulBody2)
	if t.Failed() {
		return
	}

	email := "skyler@breakingbad.com"
	password := "000111"
	skyler := runCreateUser(t, c, email, password)
	if t.Failed() {
		return
	}
	resp = runLoginUser(t, c, email, password)
	if t.Failed() {
		return
	}
	skylerAccessToken := resp.Token
	skylerBody1 := "Mr President...."
	runCreateChirp(t, c, skylerAccessToken, skylerBody1)

	saulList, err := c.GetChirpList(saul.ID.String(), "")
	if err != nil {
		t.Errorf("Failed to get Chirp list: %v", err)
	}
	skyelerList, err := c.GetChirpList(skyler.ID.String(), "")
	if err != nil {
		t.Errorf("Failed to get Chirp list: %v", err)
	}

	assert.Equal(t, 2, len(saulList))
	assert.Equal(t, 1, len(skyelerList))

	assert.Equal(t, saulBody1, saulList[0].Body)
	assert.Equal(t, saulBody2, saulList[1].Body)
	assert.Equal(t, skylerBody1, skyelerList[0].Body)
}
