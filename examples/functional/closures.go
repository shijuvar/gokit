package main

import (
	"fmt"
)

func SplitValues(f func(sum int) (int, int)) {
	x, y := f(50)
	fmt.Println(x, y)
}

func main() {
	// Initialise some arbitrary values
	a, b := 5, 8

	// a function value
	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x
		return x, y
	}
	// Passing function value as an argument
	SplitValues(fn)

	// another function value
	fn1 := func(sum int) (int, int) {
		x := sum/a + b
		y := sum - x
		return x, y
	}
	// Passing function value as an argument
	SplitValues(fn1)

}
