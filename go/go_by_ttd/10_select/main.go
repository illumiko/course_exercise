package main

import (
	"fmt"
	"net/http"
	"time"
)

func RaceWebsite(a, b string) (winner string, error error) {
	select {
	case <-pingWebsite(a):
		return a, nil
	case <-pingWebsite(b):
		return b, nil
	case <-time.After(5 * time.Millisecond):
		return "", fmt.Errorf("Time out: Website took too long to response\n a: %v, b%v", a, b)
	}
}

func pingWebsite(url string) chan struct{} {
	responseChan := make(chan struct{})
	go func() {
		http.Get(url)
		close(responseChan)
	}()
	return responseChan
}
