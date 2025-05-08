package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	for i, v := range slices.Backward(names) {
		fmt.Println(i, ":", v)
	}
}
