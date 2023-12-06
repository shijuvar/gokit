// Example code updated for Go 1.21
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func generateValues(ctx context.Context, counter chan int) {

	defer close(counter)
	n := 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Done channel received")
			if err := ctx.Err(); err != nil {
				fmt.Println("ctx-cancel error:", err)
			}
			if err := context.Cause(ctx); err != nil {
				fmt.Println("ctx-cancel cause:", err)
			}
			return
		case counter <- n:
			n++
		}
	}
}

func main() {
	causeError := errors.New("goroutine is leaking")
	ctx, cancel := context.WithCancelCause(context.Background())
	counter := make(chan int)
	// calling goroutine with ctx-cancel value
	go generateValues(ctx, counter)
	for n := range counter {
		if n == 10 {
			cancel(causeError)
		}
		fmt.Println(n)
	}
	fmt.Println("done")
}
