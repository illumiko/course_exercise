package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public/"))
	http.HandleFunc("/", handleIndex)
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))
	http.ListenAndServe(":8080", nil)
}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}
