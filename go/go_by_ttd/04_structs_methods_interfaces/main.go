package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width  float64
	lenght float64
}
type Circle struct {
	radius float64
}
type Shape interface {
	Get_area() float64
	Get_perimeter() float64
}

func (r Rectangle) Get_perimeter() float64 {
	return 2 * (r.lenght + r.width)
}

func (r Rectangle) Get_area() float64 {
	return r.lenght * r.width
}
func (r Circle) Get_area() float64 {
	return math.Pi * r.radius * r.radius
}
func (r Circle) Get_perimeter() float64 {
	return math.Pi * r.radius * 2
}

func main() {
	rect := Rectangle{10.0, 2.0}
	fmt.Println(Rectangle.Get_perimeter(rect))

}
