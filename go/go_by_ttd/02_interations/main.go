package main

// returns char repeated over n times
func Repeat(n int, char string) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += char
	}
	return repeated
}
