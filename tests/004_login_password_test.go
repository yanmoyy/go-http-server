package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/go-http-server/internal/api"
)

func TestLoginPassword(t *testing.T) {
	client := getClient(t)
	runReset(t, client)
	if t.Failed() {
		return
	}
	email := "saul@bettercall.com"
	password := "123456"
	user := runCreateUser(t, client, email, password)
	if t.Failed() {
		return
	}
	assert.Equal(t, user.Email, email)
	_ = runLoginUser(t, client, email, password)
	if t.Failed() {
		return
	}
	// login with different user
	_, err := client.Login(api.LoginParams{
		Email: email, Password: "000011112222",
	})
	if err != nil {
		t.Logf("Login failed (success): %v", err)
	} else {
		t.Errorf("Login Should be failed")
	}
}
