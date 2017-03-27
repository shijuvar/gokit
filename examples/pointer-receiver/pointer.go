// Pointers in Go
package main

import "fmt"

func main() {

	// Get value and address of a variable
	num := 10
	fmt.Printf("[num] Value = %d, Address = %p\n", num, &num)

	// Declare a pointer variable
	var p *int
	// Initialize pointer variable
	p = &num

	// Print value of pointer variable p
	fmt.Printf("[p] Value = %p\n", p)
	// Print address of pointer variable p
	fmt.Printf("[p] Address = %p\n", &p)
	// Print value of variable num using pointer p
	fmt.Printf("[p] Value of num = %d\n", *p)
	// Print address of num using pointer p
	fmt.Printf("[p] Address of num = %p\n", &(*p))
}