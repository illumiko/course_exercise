package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubStore{
		scores: map[string]int{
			"Pepper": 20,
			"Flow":   10,
		},
	}
	server := &PlayerServer{&store}
	tests := []struct {
		name               string
		player             string
		expectedHTTPStatus int
		expectedScore      string
	}{
		{"GET Pepper Score", "Pepper", http.StatusOK, "20"},
		{"GET Flow Score", "Flow", http.StatusOK, "10"},
		{"GET 404 error", "John", http.StatusNotFound, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := newGetScoreRequest(tt.player)
			resp := httptest.NewRecorder()

			server.ServeHTTP(resp, req)

			assertStatus(t, resp.Code, tt.expectedHTTPStatus)
			assertScore(t, resp.Body.String(), tt.expectedScore)

		})

	}

}
func TestStoreWins(t *testing.T) {
	store := StubStore{
		scores:   map[string]int{},
		winCalls: nil,
	}
	server := &PlayerServer{&store}

	t.Run("Accepted on Post Req", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/players/John", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("Player score was not added:%v", store.winCalls)
		}
	})

}

func assertScore(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("\nwant: %v,\ngot:%v", want, got)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("\nwant: %v,\ngot:%v", want, got)
	}

}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req

}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req

}

type StubStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score

}

func (s *StubStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
