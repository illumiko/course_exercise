package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(newInMemoryStore())
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
