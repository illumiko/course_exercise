package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func get_file_list(directory string) (file_list []string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		file_list = append(file_list, file.Name())
	}
	return file_list

}

func main() {
	directory := "./templates/"
	file_list := get_file_list(directory)
	template, err := template.ParseGlob(directory + "*")
	if err != nil {
		log.Fatal(err)
	}
	for _, files := range file_list {
		new_file, err := os.Create(files + ".html")
		if err != nil {
			log.Fatal(err)
		}
		defer new_file.Close()
		err = template.ExecuteTemplate(new_file, files, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}
