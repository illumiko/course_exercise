package main

type PingWebsites func(string) bool
type result struct {
	string
	bool
}

// CheckWebsite takes in a PingWebsites function that checks to see if the
// urls from []string are working and returns a map with the url and a bool
func CheckWebsite(pw PingWebsites, url []string) map[string]bool {
	resultChannel := make(chan result)
	response := make(map[string]bool)
	for _, v := range url {
		go func(url string) {
			resultChannel <- result{url, pw(url)}
		}(v)
	}

	for i := 0; i < len(url); i++ {
		receive := <-resultChannel
		response[receive.string] = receive.bool
	}
	return response
}
