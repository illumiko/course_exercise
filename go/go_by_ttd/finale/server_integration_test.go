package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	store := newInMemoryStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	for i := 0; i < 3; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("Get Score", func(t *testing.T) {
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, newGetScoreRequest(player))

		assertStatus(t, resp.Code, http.StatusOK)
		assertScore(t, resp.Body.String(), "3")
	})

	t.Run("Get League", func(t *testing.T) {
		resp := httptest.NewRecorder()
		req := newLeagueRequest()

		server.ServeHTTP(resp, req)

		want := []Player{
			{"John", 2},
			{"Pepper", 3},
		}
		got := getLeagueFromResponse(t, resp.Body)
		assertStatus(t, resp.Code, http.StatusOK)

		assertLeague(t, got, want)

	})

}
