package tests

import (
	"testing"

	"github.com/yanmoyy/go-http-server/internal/api"
)

func runReset(t *testing.T, client *api.Client) {
	t.Run("Reset", func(t *testing.T) {
		if err := client.Reset(); err != nil {
			t.Fatalf("Failed to reset: %v", err)
		}
		t.Log("System reset successfully")
	})
}

func runCreateUser(t *testing.T, client *api.Client, email, password string) api.User {
	var user api.User
	t.Run("CreateUser", func(t *testing.T) {
		var err error
		user, err = client.CreateUser(api.CreateUserParams{
			Email: email, Password: password,
		})
		if err != nil {
			t.Errorf("Failed to create user: %v", err)
		}
		logJSON(t, "Created user", user)
	})
	return user
}

func runCreateUserDefault(t *testing.T, c *api.Client) api.User {
	return runCreateUser(t, c, "saul@bettercall.com", "123456")
}

func runLoginUser(t *testing.T, client *api.Client, email, password string) api.LoginResponse {
	var resp api.LoginResponse
	t.Run("LoginUser", func(t *testing.T) {
		var err error
		resp, err = client.Login(api.LoginParams{
			Email: email, Password: password,
		})
		if err != nil {
			t.Errorf("Failed to login user: %v", err)
		}
		logJSON(t, "Login Response", resp)
	})
	return resp
}

func runLoginUserDefault(t *testing.T, c *api.Client) api.LoginResponse {
	return runLoginUser(t, c, "saul@bettercall.com", "123456")
}

func runUpdateUser(t *testing.T, client *api.Client, token, email, password string) api.User {
	var user api.User
	t.Run("UpdateUser", func(t *testing.T) {
		var err error
		user, err = client.UpdateUser(api.UpdateUserParams{
			Email:    email,
			Password: password,
		}, token)
		if err != nil {
			t.Errorf("Failed to update user: %v", err)
		}
		logJSON(t, "Updated User", user)
	})
	return user
}
