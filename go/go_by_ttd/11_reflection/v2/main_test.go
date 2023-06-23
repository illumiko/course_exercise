package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []Case{
		{
			Name: "test with one varible",
			Input: struct {
				Name string
			}{"John"},
			Expected: []string{"John"},
		},

		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {

			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("\nexpected: %v\ngot:%v\n", test.Expected, got)

			}

		})
	}
}

type Case struct {
	Name     string
	Input    interface{}
	Expected []string
}
