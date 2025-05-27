package main

import (
	"encoding/json"
	"net/http"

	"github.com/yanmoyy/go-http-server/internal/api"
	"github.com/yanmoyy/go-http-server/internal/auth"
)

func (cfg *apiConfig) handlePolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters api.PolkaWebhookParams

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil || apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "invalid api key", err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != api.EventUpgrade {
		respondNoContent(w)
		return
	}

	err = cfg.db.UpgradeUserChirpyRed(r.Context(), params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find user", err)
		return
	}

	respondNoContent(w)
}
