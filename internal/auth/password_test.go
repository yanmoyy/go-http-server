package auth

import "testing"

func TestPassword(t *testing.T) {
	password := "example_password"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Errorf("Failed to hash password: %v", err)
	}
	err = CheckPasswordHash(hashed, password)
	if err != nil {
		t.Errorf("Failed to check password: %v", err)
	}

	password2 := "example_password2"

	err = CheckPasswordHash(hashed, password2)
	if err != nil {
		t.Logf("CheckPasswordHash: %v", err)
	} else {
		t.Errorf("Failed: Password2 must be raise error!")
	}
}
