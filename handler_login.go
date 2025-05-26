package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yanmoyy/go-http-server/internal/api"
	"github.com/yanmoyy/go-http-server/internal/auth"
)

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type parameters api.LoginParams
	type response api.LoginResponse

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	expirationTime := time.Hour
	if params.ExpiresInSeconds > 0 && params.ExpiresInSeconds < 3600 {
		expirationTime = time.Duration(params.ExpiresInSeconds) * time.Second
	}
	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, expirationTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create access JWT", err)
	}
	respondWithJSON(w, http.StatusOK, response{
		User: api.User{
			ID:        user.ID,
			CreateAt:  user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
		Token: accessToken,
	})
}
