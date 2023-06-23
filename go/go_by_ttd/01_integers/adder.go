package integers

import "fmt"

func main() {
	fmt.Println(Adder(2, 2))
}

// prints the sum of two integers
func Adder(x, y int) int {
	return x + y
}
