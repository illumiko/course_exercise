package main

import "fmt"

func main() {
	data := "Hello Wolrd"
	template := `<h1>` + data + `</h1>`
	fmt.Println(template)
}
