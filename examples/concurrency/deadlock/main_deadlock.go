package main

import (
	"fmt"
)

func main() {
	// Declare a unbuffered channel
	counter := make(chan int)
	// This will create a deadlock
	counter <- 10          // Send operation to a channel from main goroutine
	fmt.Println(<-counter) // Receive operation from the channel
}
