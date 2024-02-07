package main

import "fmt"

func main() {
	// The values of i in this countdown program go from 0 to 9
	for i := range 10 {
		fmt.Println(i)
	}
	fmt.Println("ranging over integers is added in Go 1.22")
}
