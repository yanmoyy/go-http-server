package main

import (
	"net/http"
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
