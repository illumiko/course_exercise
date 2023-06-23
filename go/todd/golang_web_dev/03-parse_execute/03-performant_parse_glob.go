package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var directory string

func get_files(directory string) (file_list []string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, files := range files {
		file_list = append(file_list, files.Name())
	}
	return
}

func init() {
	directory = "./templates/"
	tpl = template.Must(template.ParseGlob(directory + "*")) // parseGlob returns the parameters needed for must
}
func main() {
	directory := "./templates/"
	file_list := get_files(directory)
	for _, file := range file_list {
		err := tpl.ExecuteTemplate(os.Stdout, file, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}
