package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	get := Adder(2, 2)
	expect := 4
	condCheck(t, get, expect)
}
func ExampleAdd() {
	sum := Adder(1, 2)
	fmt.Println(sum)
	//Output: 3
}
func condCheck(t testing.TB, get, expect int) {
	if get != expect {
		t.Errorf("got: %v \n wanted: %v \n", get, expect)
	}
}
