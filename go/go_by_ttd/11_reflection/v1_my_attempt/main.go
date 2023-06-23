package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	reflectVal := reflect.ValueOf(x)
	NumFieldsVal := reflectVal.NumField()
	for i := 0; i < NumFieldsVal; i++ {
		fieldType := reflectVal.Field(i).Kind()

		if fieldType == reflect.String {
			fn(reflectVal.Field(i).String())
		}
	}
}

func main() {
}
