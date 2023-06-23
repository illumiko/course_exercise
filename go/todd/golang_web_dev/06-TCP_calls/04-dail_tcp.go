package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	res, err := ioutil.ReadAll(dial)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))

}
