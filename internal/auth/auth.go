package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yanmoyy/go-http-server/internal/api"
	"golang.org/x/crypto/bcrypt"
)

type TokenType string

const (
	// TokenTypeAccess -
	TokenTypeAccess TokenType = "chirpy-access"
)

var ErrNoAuthHeaderIncluded = errors.New("no auth header include in request")

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    string(TokenTypeAccess),
		IssuedAt:  jwt.NewNumericDate(now.UTC()),
		ExpiresAt: jwt.NewNumericDate(now.UTC().Add(expiresIn)),
		Subject:   userID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	var claims jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(*jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("jwt.ParseWIthClaims: %w", err)
	}
	subject, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("token.Claims.GetSubject: %w", err)
	}
	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}
	if issuer != string(TokenTypeAccess) {
		return uuid.Nil, errors.New("invalid issuer")
	}

	id, err := uuid.Parse(subject)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID: %w", err)
	}
	return id, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get(api.HeaderAuthorization)
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	// Should split and check the header is correct format!
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) != 2 || splitAuth[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}
	return splitAuth[1], nil
}

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get(api.HeaderAuthorization)
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) != 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}
	return splitAuth[1], nil
}
