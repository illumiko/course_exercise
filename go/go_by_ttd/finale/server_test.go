package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubStore{
		scores: map[string]int{
			"Pepper": 20,
			"Flow":   10,
		},
	}
	server := NewPlayerServer(&store)
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
	server := NewPlayerServer(&store)

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

func TestLeague(t *testing.T) {
	store := StubStore{}
	server := NewPlayerServer(&store)
	t.Run("Get status 200 for /league", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		var got League
		err := json.NewDecoder(resp.Body).Decode(&got)

		if err != nil {
			t.Fatal("Unable to parse Json:", err)
		}

		assertStatus(t, resp.Code, http.StatusOK)

	})

	t.Run("Return League table as Json", func(t *testing.T) {
		expectedLeague := League{
			{"John", 20},
		}
		store := StubStore{league: expectedLeague}
		server := NewPlayerServer(&store)

		req := newLeagueRequest()
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := getLeagueFromResponse(t, resp.Body)

		assertStatus(t, resp.Code, http.StatusOK)
		assertLeague(t, got, expectedLeague)
		assertContentType(t, resp, jsonContentType)

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

func assertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v\ngot:%v", want, got)
	}

}
func assertContentType(t testing.TB, got *httptest.ResponseRecorder, want string) {
	t.Helper()
	response := got.Result().Header.Get("content-type")
	if response != want {
		t.Errorf("Unexpected content-type\nexpected:%v\ngot:%v", want, response)

	}

}

func getLeagueFromResponse(t testing.TB, body io.Reader) League {
	t.Helper()

	var got League
	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatal("json error:", err)
	}
	return got

}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
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
	league   League
}

func (s *StubStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score

}

func (s *StubStore) GetLeague() League {
	return s.league

}

func (s *StubStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
