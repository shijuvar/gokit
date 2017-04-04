package main

import "fmt"

func main() {
	// Declare a nil slice
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10, 20, 30)
	fmt.Println("Slice x after appending data:", x)
	fmt.Println(len(x), cap(x))
}
