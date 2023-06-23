package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template
var file_name string
var pfn = template.FuncMap{
	"sq":   square,
	"sqrt": square_rt,
}

func square(num float64) (val float64) {
	val = num * num
	return
}
func square_rt(num float64) (val float64) {
	val = math.Sqrt(num)
	return
}

func init() {
	file_name = "03-pipeline.gohtml"
	tpl = template.Must(template.New("").Funcs(pfn).ParseFiles("./templates/" + file_name))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, file_name, float64(4))
	if err != nil {
		log.Fatal(err)
	}
}
