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
		go handle(conn)
	}

}
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		scanned := scanner.Text()
		fmt.Print(scanned)
		fmt.Fprint(conn, "see mom i can read and write \n")
	}
	defer conn.Close()

}
