package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)

}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
