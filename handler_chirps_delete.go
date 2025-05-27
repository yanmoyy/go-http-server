package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-http-server/internal/api"
	"github.com/yanmoyy/go-http-server/internal/auth"
)

func (cfg *apiConfig) handleDeleteChirpByID(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue(api.ChirpIDParam)
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chirp ID", err)
		return
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "No Token", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid Token", err)
		return
	}
	chirp, err := cfg.db.GetChirpById(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp Not Found", err)
		return
	}
	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "Not author of chirp", err)
		return
	}
	err = cfg.db.DeleteChirpById(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete chirp", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
