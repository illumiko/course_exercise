package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template
var templateIndex = "02-index.gohtml"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	data := []int{1, 2, 3, 4, 5}
	tpl.ExecuteTemplate(w, templateIndex, data)
}
func init() {
	tpl = template.Must(template.ParseFiles("./" + templateIndex))
}
func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)

}
