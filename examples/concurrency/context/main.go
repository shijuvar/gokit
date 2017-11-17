package main

import (
	"context"
	"fmt"
	"time"
)

func generateValues(ctx context.Context, counter chan int) {
	n := 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			return
		case counter <- n:
			n++
		}
	}
}

func main() {
	// WithCancel returns a copy of parent with a new Done channel. The returned
	// context's Done channel is closed when the returned cancel function is called
	// or when the parent context's Done channel is closed, whichever happens first.
	//
	// Canceling this context releases resources associated with it, so code should
	// call cancel as soon as the operations running in this Context complete.
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	counter := make(chan int)
	defer cancel()
	go generateValues(ctx, counter)
	for n := range counter {
		fmt.Println(n)
		//if n == 10 {
		//	cancel()
		//}
	}
}
