package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	os.Exit(m.Run())
}
