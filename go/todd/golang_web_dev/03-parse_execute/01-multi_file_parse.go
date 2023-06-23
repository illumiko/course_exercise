package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	directory := "./templates/"
	template_name_one := "template.gohtml"
	template_name_two := "02-template.gohtml"
	tpl, err := template.ParseFiles(directory + template_name_one)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
	tpl, err = template.ParseFiles(directory + template_name_two)
	err = tpl.ExecuteTemplate(os.Stdout, template_name_two, nil)
	if err != nil {
		log.Fatal(err)
	}

}
