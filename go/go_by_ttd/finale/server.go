package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.handleLeague))
	router.Handle("/players/", http.HandlerFunc(p.handlePlayers))

	p.Handler = router

	return p
}

func (p *PlayerServer) handlePlayers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.getPlayer(w, r)
	}
}

func (p *PlayerServer) handleLeague(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")

	p.store.RecordWin(playerName)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) getPlayer(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(playerName)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)

}
