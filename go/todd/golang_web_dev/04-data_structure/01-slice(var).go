package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./template/01-slice(var).gohtml"))
}
func main() {
	animals := []string{"monke", "donke", "pande"}
	err := tpl.Execute(os.Stdout, animals)
	if err != nil {
		log.Fatal(err)
	}

}
