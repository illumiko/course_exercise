package main

import (
	"fmt"
)

func main() {
	fmt.Println(Slice_sum([]int{2, 3}))
	fmt.Println(Slice_sum_all([]int{2, 3, 4}, []int{1, 2, 3}))

}
func Slice_sum_all(x ...[]int) []int {
	get := []int{}
	lols := "Hello"
	fmt.Println(lols)
	for _, val := range x {
		get = append(get, Slice_sum(val))
	}
	return get
}
func Slice_sum(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}
