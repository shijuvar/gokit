package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30}
	fmt.Printf("[Slice:x] Length is %d Capacity is %d\n", len(x), cap(x))
	// Create a bigger slice
	y := make([]int, 5, 10)
	copy(y, x)
	fmt.Printf("[Slice:y] Length is %d Capacity is %d\n", len(y), cap(y))
	fmt.Println("Slice y after copying:", y)
	y[3] = 40
	y[4] = 50
	fmt.Println("Slice y after adding elements:", y)
}
