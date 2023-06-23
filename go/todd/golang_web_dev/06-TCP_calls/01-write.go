package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, "Hello World")
		fmt.Fprint(conn, "HELLO FMT")
		conn.Close() //dialing into tcp doesnt work without this

	}
}
