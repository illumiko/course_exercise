package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// A simple replication of the cat functionality proviced in the unix toolset
// this helps to understand how to read files on go
func cat(fname string) error {
	file, err := os.Open(fname)
	_, err = io.Copy(os.Stdout, file)

	file.Close()
	return err
}

func wc(fname string) map[string]int {
	var word_count, char_count, line_count int

	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		s := scan.Text()

		word_count += len(strings.Fields(s))
		char_count += len(s)
		line_count++
	}
	return map[string]int{"lc": line_count, "wc": word_count, "cc": char_count}
}

func main() {
	switch os.Args[1] {
	case "wc":
		parameter := string(os.Args[2])
		parameter_check := "-" == parameter[0:1]
		if parameter_check {
			for _, files := range os.Args[3:] {
				info := wc(files)
				fmt.Fprintf(os.Stdout, "%v", info[parameter[1:]])
			}
		} else {
			for _, files := range os.Args[2:] {
				info := wc(files)
				fmt.Fprintf(os.Stdout, "%v", info)
			}
		}
	case "cat":
		for _, args := range os.Args[2:] {
			if err := cat(args); err != nil {
				fmt.Fprint(os.Stderr, err)
				continue
			}
		}
	}

}
