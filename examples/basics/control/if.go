package main

import (
	"fmt"
)

func main() {
	isEven(8)
}
func isEven(input int) {
	if input%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
