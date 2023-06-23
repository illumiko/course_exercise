package main

import (
	"fmt"
	"net/http"
)

type handle int

func (m handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func main() {
	var h handle
	http.ListenAndServe(":8080", h)

}
