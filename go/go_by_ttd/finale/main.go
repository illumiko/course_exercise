package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(&InMemoryStore{})
	log.Fatal(http.ListenAndServe(":8080", server))
}
