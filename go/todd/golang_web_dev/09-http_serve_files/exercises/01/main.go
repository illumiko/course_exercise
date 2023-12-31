// # ListenAndServe on port 8080 of localhost
//
// For the default route "/"
// Have a func called "foo"
// which writes to the response "foo ran"
//
// For the route "/dog/"
// Have a func called "dog"
// which parses a template called "dog.gohtml"
// and writes to the response "<h1>This is from dog</h1>"
// and also shows a picture of a dog when the template is executed.
//
// Use "http.ServeFile"
// to serve the file "dog.jpeg"
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/toby.jpg", handleDoggo)
	http.ListenAndServe(":8080", nil)
}
func handleDoggo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./toby.jpg")
	// io.WriteString(w, `<h1>Hello World</h1>`)
}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<h1>Hello World</h1>`)
}
