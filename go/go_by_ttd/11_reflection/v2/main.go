package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	reflectVal := reflect.ValueOf(x)

	for i := 0; i < reflectVal.NumField(); i++ {
		field := reflectVal.Field(i)

		switch field.Kind() {

		case reflect.String:
			fn(field.String())

		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}
