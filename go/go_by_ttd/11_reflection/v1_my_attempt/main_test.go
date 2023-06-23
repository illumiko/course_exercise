package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	type x struct {
		name string
		age  int
	}

	structX := x{"John", 10}
	Walk(structX, func(s string) {
		repS := strings.Repeat(s, 2)
		structX.name = repS

	})
	expect := x{"JohnJohn", 10}
	if !reflect.DeepEqual(structX, expect) {
		t.Errorf("\nexpect: %v\nresponse: %v", expect, structX)

	}
}
