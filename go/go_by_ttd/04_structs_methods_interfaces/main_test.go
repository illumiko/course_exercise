package main

import "testing"

func Test_main(t *testing.T) {
	area_tests := []struct {
		name     string
		shape    Shape
		has_area float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 10.0, lenght: 1.0}, has_area: 10.0},
		{name: "Circle", shape: Circle{radius: 1}, has_area: Circle{radius: 1}.Get_area()},
	}
	for _, val := range area_tests {
		t.Run(val.name, func(t *testing.T) {
			if val.shape.Get_area() != val.has_area {
				t.Errorf("\nget: %.2f \n want: %.2f", val.shape.Get_area(), val.has_area)
			}
		})

	}
}
