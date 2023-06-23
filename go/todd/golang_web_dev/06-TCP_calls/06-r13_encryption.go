package main

import (
	"bufio"
	"fmt"
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		getText := strings.ToLower(scanner.Text())
		bt := []byte(getText)
		encrypted := encrypt(bt)
		fmt.Fprintf(conn, "Result: %s -> %s \n", getText, encrypted)
	}
}
func encrypt(data []byte) (encrypted []byte) {
	encrypted = make([]byte, len(data))
	for i, v := range data {
		if v <= 109 {
			encrypted[i] = v + 13
		} else {
			encrypted[i] = v - 13
		}
	}
	return

}
