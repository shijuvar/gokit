package main

import (
	"fmt"
)

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Providing four arguments
	total := Sum(1, 2, 3, 4)
	fmt.Println("The Sum is:", total)

	// Providing three arguments
	total = Sum(5, 7, 8)
	fmt.Println("The Sum is:", total)

	// without arguments
	total = Sum()
	fmt.Println("The Sum is:", total)

	// Providing a Slice as an argument
	nums := []int{1, 2, 3, 4, 5}
	total = Sum(nums...)
	fmt.Println("The Sum is:", total)
}
