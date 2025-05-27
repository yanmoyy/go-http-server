package main

import (
	"net/http"

	"sort"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-http-server/internal/api"
	"github.com/yanmoyy/go-http-server/internal/database"
)

func (cfg *apiConfig) handleGetChirpList(w http.ResponseWriter, r *http.Request) {
	resp := []Chirp{}
	query := r.URL.Query()
	authorIDString := query.Get(api.QueryAuthorID)
	sortString := query.Get(api.QuerySort)

	var list []database.Chirp
	var err error

	if authorIDString == "" {
		list, err = cfg.db.GetAllChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get chirp list", err)
			return
		}
	} else {
		authorID, err := uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "malformed author ID", err)
			return
		}
		list, err = cfg.db.GetChirpsByUser(r.Context(), authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get author's chirp list", err)
			return
		}
	}

	if sortString == api.SortDesc {
		sort.Slice(list, func(i, j int) bool { return list[i].CreatedAt.After(list[j].CreatedAt) })
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
	idString := r.PathValue(api.ParamChirpID)
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid chirp id", err)
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
