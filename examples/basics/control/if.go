package main

import (
	"flag"
	"fmt"
)

func main() {
	input := flag.Int("input", 0, "an integer value")
	flag.Parse()
	if *input%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
