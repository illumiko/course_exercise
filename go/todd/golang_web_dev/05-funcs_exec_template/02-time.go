package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(pfn).ParseFiles("./templates/02-time.gohtml"))
}

var pfn = template.FuncMap{
	"format": time_format,
}

func time_format(t time.Time) string {
	return t.Format("ANSIC")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "02-time.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
