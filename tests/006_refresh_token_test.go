package tests

import (
	"testing"

	"github.com/yanmoyy/go-http-server/internal/api"
)

func TestRefreshToken(t *testing.T) {
	c := getClient(t)
	runReset(t, c)
	if t.Failed() {
		return
	}
	runCreateUserDefault(t, c)
	if t.Failed() {
		return
	}
	resp := runLoginUserDefault(t, c)
	// create with refresh token (err)
	_, err := c.CreateChirp(
		api.CreateChirpParams{
			Body: "Let’s just say I know a guy... who knows a guy... who knows another guy.",
		}, resp.RefreshToken,
	)
	if err == nil {
		t.Errorf("CreateChirp with Refresh Token: should be error")
	}
	runCreateChirp(t, c, resp.Token, "Let’s just say I know a guy... who knows a guy... who knows another guy.")
	if t.Failed() {
		return
	}
	tokenResp, err := c.RefreshToken(resp.RefreshToken)
	if err != nil {
		t.Errorf("Failed refresh token: %v", err)
	}
	runCreateChirp(t, c, tokenResp.Token, "I'm the guy who's gonna win you this case.")

	err = c.RevokeToken(resp.RefreshToken)
	if err != nil {
		t.Errorf("Failed revoke token: %v", err)
	}
	_, err = c.RefreshToken(resp.RefreshToken)
	if err == nil {
		t.Errorf("Refresh Token after revoke: should be error")
	}
}
