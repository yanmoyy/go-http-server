package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yanmoyy/go-http-server/internal/api"
)

func (cfg *apiConfig) handlePolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters api.PolkaWebhookParams

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	fmt.Printf("%+v\n", params)

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
