package main

import (
	"bufio"
	"fmt"
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
	newScanner := bufio.NewScanner(conn)
	for newScanner.Scan() {
		ln := newScanner.Text()
		fmt.Fprintln(conn, "Ditto: "+ln)
		if ln == "" {
			break
		}
	}
	fmt.Fprintf(conn, "conn ended\n")
}
