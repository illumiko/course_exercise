package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var tpl *template.Template

type method_handle int

func (m method_handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Method string
		Submit url.Values
	}{r.Method, r.Form}
	tpl.ExecuteTemplate(w, "02-ParseForm_method.gohtml", data)
}
func init() {
	tpl = template.Must(template.ParseFiles("./02-ParseForm_method.gohtml"))
}
func main() {
	var handle method_handle
	http.ListenAndServe(":8080", handle)

}
