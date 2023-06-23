package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	if template, err := template.ParseFiles("./templates/template.gohtml"); err == nil {
		if new_file, err := os.Create("index.html"); err == nil {
			defer new_file.Close()
			if err := template.Execute(new_file, nil); err != nil {
				log.Fatal(err)
			}
		}
	}
}
