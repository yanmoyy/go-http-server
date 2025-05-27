package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-http-server/internal/api"
)

func (cfg *apiConfig) handleGetChirpList(w http.ResponseWriter, r *http.Request) {
	resp := []Chirp{}

	list, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirp list", err)
		return
	}
	for _, chirp := range list {
		resp = append(resp, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, resp)
}

func (cfg *apiConfig) handleGetChirpByID(w http.ResponseWriter, r *http.Request) {

	type response struct {
		Chirp
	}
	idString := r.PathValue(api.ChirpIDParam)
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp id", err)
		return
	}
	chirp, err := cfg.db.GetChirpById(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "couldn't fetch chirp", err)
		return
	}
	respondWithJSON(w, http.StatusOK, response{
		Chirp: Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		},
	})
}
