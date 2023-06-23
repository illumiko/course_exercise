package main

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	mockPingWebsites := func(url string) bool {
		if url == "https://yahoo.net" {
			return false
		}
		return true
	}

	websites := []string{
		"https://google.com",
		"https://youtube.com",
		"https://yahoo.net",
	}
	response := CheckWebsite(mockPingWebsites, websites)
	expect := map[string]bool{
		"https://google.com":  true,
		"https://youtube.com": true,
		"https://yahoo.net":   false,
	}
	if !reflect.DeepEqual(expect, response) {
		t.Errorf("\nExpect: %v\nResponse: %v", expect, response)
	}

}
func BenchmarkCheckWebstei(b *testing.B) {
	var wg sync.WaitGroup
	urls := make([]string, 100)

	slowPingWebsites := func(url string) bool {
		time.Sleep(20 * time.Millisecond)
		return true
	}

	for i := 0; i >= len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			CheckWebsite(slowPingWebsites, urls)
			defer wg.Done()
		}()
	}
	wg.Wait()
}
