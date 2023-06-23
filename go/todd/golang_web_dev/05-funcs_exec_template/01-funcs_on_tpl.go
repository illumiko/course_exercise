package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var passed_funcs = template.FuncMap{
	"length":     check_lenght,
	"loop_to_10": loop_to_10,
}

func check_lenght(s string) int {
	return len(s)
}
func loop_to_10() (x []int) {
	for i := 0; i <= 10; i++ {
		x = append(x, i)
	}
	return
}

func init() {
	tpl = template.Must(template.New("").Funcs(passed_funcs).ParseFiles("./templates/01-funcs_on_tpl.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "01-funcs_on_tpl.gohtml", "Hello World")
	if err != nil {
		log.Fatal(err)
	}

}
