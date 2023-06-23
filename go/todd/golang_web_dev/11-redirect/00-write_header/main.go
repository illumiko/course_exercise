package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/data", handleData)
	http.HandleFunc("/redirect", handleRedirect)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./template/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}
func handleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/redirect")
	w.WriteHeader(http.StatusSeeOther)
	fmt.Println(r.Method)
	//same as http.Redirect but more fundamental
}
func handleRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
}
