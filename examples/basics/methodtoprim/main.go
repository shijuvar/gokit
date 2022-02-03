// Add methods to primitive types
package main

import "fmt"

type Number int

func (n Number) Positive() bool {
	return n > 0
}

func main() {
	num1 := Number(100)
	num2 := Number(-50)
	fmt.Printf("%d is positive: %t\n", num1, num1.Positive())
	fmt.Printf("%d is positive: %t\n", num2, num2.Positive())
}
