package main

import (
	"fmt"
)

func main() {
	// Declare a unbuffered channel
	counter := make(chan int)
	// Creates a buffered channel with capacity of 3
	nums := make(chan int, 3)
	go func() {
		// Send value to the unbuffered channel
		counter <- 1
		close(counter) // Closes the channel
	}()

	go func() {
		// Send values to the buffered channel
		nums <- 10
		nums <- 30
		nums <- 50
	}()
	// Read the value from unbuffered channel
	fmt.Println(<-counter)
	val, ok := <-counter // Trying to read from closed channel
	if ok {
		fmt.Println(val) // This won't execute
	}
	// Read the 3 buffered values from the buffered channel
	fmt.Println(<-nums)
	fmt.Println(<-nums)
	fmt.Println(<-nums)
	close(nums) // Closes the channel
}
