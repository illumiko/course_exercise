package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
        <form method="GET"> 
            <span>Name: </span>
            <input type="text" name="name">
            <input type="submit">
        </form
        <br>
    `+val)
}
