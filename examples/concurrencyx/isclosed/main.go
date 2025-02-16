package main

import "fmt"

// IsClosed provides a generic function to understand that whether the channel is closed
func IsClosed[T any](ch <-chan T) bool {
	select {
	case _, ok := <-ch:
		if !ok {
			return true
		}
	default:
	}

	return false
}

func main() {
	c := make(chan int)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}
