package main

import (
	"fmt"
)

func main() {
	x := make([]int, 3, 5)
	x[0] = 10
	x[1] = 20
	x[2] = 30
	println(x, "x")

	y := make([]int, 3)
	y[0] = 10
	y[1] = 20
	y[2] = 30
	println(y, "y")

	// Slice literal
	z := []int{10, 20, 30}
	println(z, "z")

	z1 := []int{0: 10, 2: 30}
	println(z1, "z1")

	slice := []int{10, 20, 30}
	// Appending slice
	slice = append(slice, 40, 50)
	slice = append(slice, []int{60, 70, 80}...)
	println(slice, "slice")

	sliced := slice[1:3] // i, j-i, cap cap - i
	println(sliced, "sliced")

}

// Helper function that prints array, length and capacity
func println(s []int, label string) {
	fmt.Println()
	fmt.Println("Slice:", label)
	fmt.Println("Array:", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
}
