package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	// http.Handle("/public", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":8080", nil)
}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}
