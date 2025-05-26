package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/go-http-server/internal/api"
)

func TestUpdateUser(t *testing.T) {
	c := getClient(t)
	runReset(t, c)
	if t.Failed() {
		return
	}

	email := "walt@breakingbad.com"
	password := "123456"

	_ = runCreateUser(t, c, email, password)
	if t.Failed() {
		return
	}
	resp := runLoginUser(t, c, email, password)
	jwtToken := resp.Token

	newPassword := "losPollosHermanos"
	updatedUser := runUpdateUser(t, c, jwtToken, email, newPassword)
	if t.Failed() {
		return
	}
	assert.Equal(t, email, updatedUser.Email)

	newPassword2 := "j3ssePinkM@nCantCook"
	badToken := "badToken"

	// With no Token
	_, err := c.UpdateUser(api.UpdateUserParams{
		Email:    email,
		Password: newPassword2,
	}, "")
	if err == nil {
		t.Errorf("Should Error (no token): %v", err)
	}

	// With bad Token
	_, err = c.UpdateUser(api.UpdateUserParams{
		Email:    email,
		Password: newPassword2,
	}, badToken)
	if err == nil {
		t.Errorf("Should Error (bad token): %v", err)
	}
}
