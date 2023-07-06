package main

import (
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

	t.Run("GET Pepper Score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := resp.Body.String()
		want := "20"
		// assertHTTPMethod(t, http.MethodGet, req.Method)
		assertStatus(t, resp.Code, http.StatusOK)
		assertScore(t, got, want)

	})

	t.Run("GET Flow's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Flow", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := resp.Body.String()
		want := "10"

		assertStatus(t, resp.Code, http.StatusOK)
		assertScore(t, got, want)
	})

	t.Run("Test for 4040 error", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/John", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := resp.Code
		want := http.StatusNotFound

		assertStatus(t, got, want)

	})

}
func TestStoreWins(t *testing.T) {
	store := StubStore{
		scores: map[string]int{},
	}
	server := &PlayerServer{&store}

	t.Run("Accepted on Post Req", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/players/John", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusAccepted)
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

type StubStore struct {
	scores map[string]int
}

func (s *StubStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score

}
