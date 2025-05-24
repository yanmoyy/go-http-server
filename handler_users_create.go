package main

import (
	"encoding/json"
	"net/http"

	"github.com/yanmoyy/go-http-server/internal/api"
)

type User api.User

func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters api.CreateUserParams

	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldin't decode parameters", err)
		return
	}
	user, err := cfg.db.CreateUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        user.ID,
			CreateAt:  user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
	})
}
