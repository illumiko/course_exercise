package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) { // {{{
	t.Run("Does SliceSum return the sum of all the ings", func(t *testing.T) {
		arr := []int{2, 2, 3}
		get := Slice_sum(arr)
		expect := 7
		cond_check(t, get, expect, arr)

	})
} // }}}
func Test_sum_all(t *testing.T) {
	check_sum := func(t testing.TB, get, expect []int) {
		t.Helper()
		if !reflect.DeepEqual(get, expect) {
			t.Errorf("\n got: %v \n wanted: %v \n", get, expect)
		}
	}
	t.Run("Testing", func(t *testing.T) {
		arr := []int{1, 2}
		arr2 := []int{3, 2}
		get := Slice_sum_all(arr, arr2)
		expect := []int{3, 5}
		check_sum(t, get, expect)
	})
	t.Run("Testing with 0 vals", func(t *testing.T) {
		arr := []int{1, 2}
		arr2 := []int{0, 0}
		get := Slice_sum_all(arr, arr2)
		expect := []int{3, 0}
		check_sum(t, get, expect)
	})

}
func Example_slice_sum() {
	sum := Slice_sum_all([]int{1, 2}, []int{3, 5})
	fmt.Printf("%v", sum)
	//Output: [3 8]

}
func cond_check(t testing.TB, get, expect int, numbers []int) {
	if get != expect {
		t.Errorf("\n got: %v \n wanted: %v \n given: %v\n", get, expect, numbers)
	}

}
