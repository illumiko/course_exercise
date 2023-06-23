package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello", r.URL)
}
func main() {
	newMux := http.NewServeMux()
	newMux.Handle("/", http.HandlerFunc(handleIndex))
	http.ListenAndServe(":8080", newMux)
}
