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
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, newGetScoreRequest(player))

	assertStatus(t, resp.Code, http.StatusOK)
	assertScore(t, resp.Body.String(), "3")

}
