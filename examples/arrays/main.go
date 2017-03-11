package main

import (
	"fmt"
)

func main() {
	// Declare arrays
	var x [5]int
	// Assign values at specific index
	x[0] = 5
	x[4] = 25
	fmt.Println("Value of x:", x)

	x[1] = 10
	x[2] = 15
	x[3] = 20
	fmt.Println("Value of x:", x)

	// Declare and initialize array with array literal
	y := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Value of y:", y)
	fmt.Println("Length of y:", len(y))

	// Array literal with ...
	z := [...]int{10, 20, 30, 40, 50}
	fmt.Println("Value of z:", z)
	fmt.Println("Length of z:", len(z))

	// Initialize values at specific index with array literal
	langs := [4]string{0: "Go", 3: "Julia"}
	fmt.Println("Value of langs:", langs)
	// Assign values to remain positions
	langs[1] = "Rust"
	langs[2] = "Scala"

	// Iterate over the elements of array
	fmt.Println("Value of langs:", langs)
	fmt.Println("\nIterate over arrays\n")
	for i := 0; i < len(langs); i++ {
		fmt.Printf("langs[%d]:%s \n", i, langs[i])
	}
	fmt.Println("\n")

	// Iterate over the elements of array using range
	for k, v := range langs {
		fmt.Printf("langs[%d]:%s \n", k, v)
	}
	for k := range langs {
		fmt.Printf("Index:%d \n", k)
	}
	for _, v := range langs {
		fmt.Printf("Value:%s \n", v)
	}
}
