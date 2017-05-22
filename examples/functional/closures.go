package main

import (
	"fmt"
)

func SplitValues(f func(sum int) (int, int)) {
	x, y := f(35)
	fmt.Println(x, y)

	x, y = f(50)
	fmt.Println(x, y)
}

func main() {
	a, b := 5, 8
	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x
		return x, y
	}

	// Passing function value as an argument to another function
	SplitValues(fn)

	// Calling the function value by providing argument
	x, y := fn(20)
	fmt.Println(x, y)
}
