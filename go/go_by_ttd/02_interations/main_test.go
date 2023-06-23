package main

import "testing"

func TestRepeatChar(t *testing.T) {
	repeatedString := Repeat(4, "a")
	expectedString := "aaaa"
	conditionCheck(t, repeatedString, expectedString)

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(4, "b")
	}
}

func conditionCheck(t testing.TB, response, expect string) {
	if response != expect {
		t.Errorf("\nResponse: %v\nExpected: %v\n", response, expect)
	}
}
