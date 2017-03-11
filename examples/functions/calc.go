package main

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func main() {
	x, y := 20, 10

	result := Add(x, y)
	fmt.Println("[Add]:", result)

	result = Subtract(x, y)
	fmt.Println("[Subtract]:", result)
}
