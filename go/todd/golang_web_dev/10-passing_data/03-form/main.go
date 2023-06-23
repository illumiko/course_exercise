package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type person struct {
	Fname string
	Lname string
}

func init() {
	tpl = template.Must(template.ParseGlob("./template/*.gohtml"))
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	fn := r.FormValue("fn")
	ln := r.FormValue("ln")
	user := person{fn, ln}
	err := tpl.ExecuteTemplate(w, "index.gohtml", user)
	if err != nil {
		log.Fatal(err)
	}

}
