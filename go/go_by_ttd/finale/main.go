package main

import (
	"log"
	"net/http"
)

type InMemoryStore struct{}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return 123

}

func main() {
	server := &PlayerServer{&InMemoryStore{}}
	log.Fatal(http.ListenAndServe(":8080", server))
}
