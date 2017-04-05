package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("deferred function 1, value of i is", i)
	defer fmt.Println("deferred function 2")
	defer fmt.Println("deferred function 3")
	i++
	fmt.Println("The current value of i is", i)
}
