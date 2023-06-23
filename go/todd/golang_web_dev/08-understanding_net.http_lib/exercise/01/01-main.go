package main

import (
	"fmt"
	"net/http"
)

func handleMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello im miko")
}
func handleDog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello im a dog")
}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main() {
	http.HandleFunc("/me/", handleMe)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}
