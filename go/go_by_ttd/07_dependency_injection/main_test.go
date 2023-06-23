package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	check_condition := func(t testing.TB, response string, expect string) {
		t.Helper()
		if expect != response {
			t.Errorf("Expect: %v\n response: %v", expect, response)
		}
	}
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")
	response := buffer.String()
	respond := "Hello, Chris"
	check_condition(t, response, respond)

}
