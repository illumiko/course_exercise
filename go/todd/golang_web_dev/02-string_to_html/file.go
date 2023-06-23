package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	data := "Hello Wolrd"
	tpl := fmt.Sprint(`<h1>` + data + `</h1>`)
	if res, err := os.Create("index.html"); err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		defer res.Close()
		io.Copy(res, strings.NewReader(tpl))
	}
}
