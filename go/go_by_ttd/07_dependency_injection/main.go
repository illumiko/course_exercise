package main

import (
	"bytes"
	"fmt"
)

// This function works as a greeter to some bla bla
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
	
}
