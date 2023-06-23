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
	val := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "value: "+val)
}

/*
   Here, im getting the value of "q", displaying the value at the route "/"
*/
