package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	response(conn)
}
func response(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		scaned := scanner.Text()
		fmt.Fprintln(conn, scaned)
	}

	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.")
}
