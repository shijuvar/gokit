package main

import (
	"fmt"
)

func main() {
	x := make([]int, 2, 5)
	x[0] = 10
	x[1] = 20
	fmt.Println("Slice x:", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))
	// Create a bigger slice
	x = append(x, 30, 40, 50)
	fmt.Println("Slice x after appending data:", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))

	x = append(x, 60, 70, 80)
	fmt.Println("Slice x after appending data for the second time:", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))

}
