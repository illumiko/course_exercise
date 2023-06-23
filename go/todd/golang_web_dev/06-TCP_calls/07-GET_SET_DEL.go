package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	defer conn.Close()
	//Guide
	io.WriteString(conn, "GET [insert key] \nSET [insert key] [insert value] \nDEL [insert key]\n")
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		user_input := strings.Fields(ln)
		switch user_input[0] {
		case "SET":
			if len(user_input) < 3 {
				fmt.Fprintln(conn, "MISSING ARGUMENTS\n")
				continue
			}
			key := user_input[1]
			value := user_input[2]
			data[key] = value
		case "GET":
			key := user_input[1]
			fmt.Fprintf(conn, "key -> %v \n", data[key])
		case "DEL":
			key := user_input[1]
			delete(data, key)
		case "LS":
			if len(data) <= 0 {
				fmt.Fprintln(conn, "database empty")
			} else {
				fmt.Fprintf(conn, "%v -> %v \n", "index", "value")
				for i, v := range data {
					fmt.Fprintf(conn, "%v -> %v\n", i, v)
				}
			}
		}

	}

}
