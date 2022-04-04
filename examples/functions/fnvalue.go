package main

import "fmt"

func main() {
	/*
		Go functions are like values so that functions can be
		assigned to variables, and those function values can be used
		for calling other functions which is having function value
		as parameters, or functions can return function as values
	*/

	fn := func(principal, interest float64, period int) float64 {
		var total float64
		total = (principal * interest * float64(period)) / 100
		return total
	}
	result := fn(100, 10, 1)
	fmt.Println(result)
}
