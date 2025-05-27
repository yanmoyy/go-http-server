package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/go-http-server/internal/api"
)

func TestGetChirpListSortTest(t *testing.T) {
	c := getClient(t)
	if t.Failed() {
		return
	}

	runReset(t, c)
	if t.Failed() {
		return
	}

	saul := runCreateUserDefault(t, c)
	if t.Failed() {
		return
	}

	resp := runLoginUserDefault(t, c)
	if t.Failed() {
		return
	}
	saulAccessToken := resp.Token

	bodies := []string{
		"I'm the one who knocks!",
		"Gale!",
		"Cmon Pinkman",
		"Darn that fly, I just wanna cook",
	}
	for _, body := range bodies {
		runCreateChirp(t, c, saulAccessToken, body)
		if t.Failed() {
			return
		}
	}
	descList, err := c.GetChirpList(saul.ID.String(), api.SortDesc)
	if err != nil {
		t.Errorf("Failed to get Chirp list: %v", err)
	}
	for i, chirp := range descList {
		assert.Equal(t, bodies[len(bodies)-1-i], chirp.Body)
	}

	ascList, err := c.GetChirpList(saul.ID.String(), api.SortAsc)
	if err != nil {
		t.Errorf("Failed to get Chirp list: %v", err)
	}
	for i, chirp := range ascList {
		assert.Equal(t, bodies[i], chirp.Body)
	}
}
