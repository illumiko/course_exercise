package main

import "testing"

func condCheck(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted: %v \n got: %v", want, got)
	}

}
func TestHello(t *testing.T) {
	got := HelloWorld("", "eyanat")
	want := "hello " + "eyanat"
	condCheck(t, got, want)
	t.Run("Return, hello world; when passed an empty string", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "hello " + "world"
		condCheck(t, got, want)

	})
	t.Run("Return hello world in bangla", func(t *testing.T) {
		got := HelloWorld("Bangla", "Prithibi")
		want := "hello " + "Prithibi"
		condCheck(t, got, want)
	})
}
