package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0], name)
	str := fmt.Sprint(`<h1>` + name + `</h1>`)
	if res, err := os.Create("os-args.html"); err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		io.Copy(res, strings.NewReader(str))
		res.Close() // closes os.create
	}
	//reading created html file
	if tpl, err := template.ParseFiles("os-args.html"); err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		if err = tpl.Execute(os.Stdout, nil); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
