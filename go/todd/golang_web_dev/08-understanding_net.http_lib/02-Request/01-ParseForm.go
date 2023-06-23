package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type xox int

func (x xox) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Body)
	fmt.Println("i ran")
	tpl.ExecuteTemplate(w, "01-index.gohtml", r.Form)

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./01-index.gohtml"))
}

func main() {
	var idk xox
	http.ListenAndServe(":8080", idk)

}
