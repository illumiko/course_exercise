package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//listens for a call on 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close() // ends the call after the work is done
	for {             // listens to the call
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go read(conn) // response to the call
	}

}
func read(conn net.Conn) {
	scanner := bufio.NewScanner(conn) //creaters scanner
	for scanner.Scan() {              //scans the text
		res := scanner.Text() //puts the scanned text into a variable
		fmt.Print(res + "\n")
	}
	defer conn.Close()
	//due to the for loop of conn the server never stops listening
	// and the scanner loop also continues to scaan  for the loop hence code here isnt executed
	fmt.Print("DONE")

}
