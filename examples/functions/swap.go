package main

import (
	"fmt"
)

func Swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x, y := "Shiju", "Varghese"
	fmt.Println("Before Swap:", x, y)

	x, y = Swap(x, y)
	fmt.Println("After Swap:", x, y)
}
