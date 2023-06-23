package main

import (
	"fmt"
	"math"
)

type shapes interface {
	area() float64
}

type square struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r *square) area() (area float64) {
	area = r.height * r.width
	fmt.Printf("Area: %v", area)
	return
}
func (r circle) area() (area float64) {
	area = r.radius * 2 * math.Pi
	fmt.Printf("Area: %v", area)
	return
}
func info(g shapes) {
	g.area()
}

func main() {
	// big_square := square{10.0, 20.0}
	small_circle := circle{1.0}
	small_circle.area()

}
