package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./template/*"))
}
func main() {
	http.HandleFunc("/", handleIndex)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	var str string
	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("q")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		// fmt.Println("file: ", file, "\n", "header: ", header, "\n", "err: ", err)
		fmt.Printf("header type: %T\nfile type: %T", header, file)
		rFile, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		dst, err := os.Create(filepath.Join("./user_files/", "copy_"+header.Filename))
		if err != nil {
			log.Fatal(err)
		}
		defer dst.Close()
		_, err = dst.Write(rFile)
		if err != nil {
			log.Fatal(err)
		}
		str = string(rFile)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", str)

}
