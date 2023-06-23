package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("\nTest response", func(t *testing.T) {

		fastServer := mockServer(0 * time.Millisecond)
		slowServer := mockServer(1 * time.Millisecond)
		defer fastServer.Close()
		defer slowServer.Close()

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		response, _ := RaceWebsite(slowUrl, fastUrl)
		expect := fastUrl

		if response != expect {
			t.Errorf("\nExpect: %v\nResponse: %v\n", expect, response)
		}

	})

	t.Run("\nTest timeoout", func(t *testing.T) {
		fastServe := mockServer(25 * time.Millisecond)

		defer fastServe.Close()

		fastUrl := fastServe.URL

		_, response := RaceWebsite(fastUrl, fastUrl)

		if response == nil {
			t.Error("No error. Expected error", response)
		}
	})

}

func mockServer(t time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
