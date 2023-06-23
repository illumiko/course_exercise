package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./template/*"))
}

func main() {
	http.HandleFunc("/data", handleData)
	http.HandleFunc("/redirect", handleRedirect)
	http.HandleFunc("/", handleIndex)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "method: "+r.Method)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	status := http.StatusSeeOther
	//change this to see other redirect codes
	// ~/Documents/Projects/Course/Todd/golang_web_dev/completed/028_redirect/status-codes.html
	http.Redirect(w, r, "/redirect", status)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

}
