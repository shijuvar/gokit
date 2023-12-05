// Example code updated for Go 1.21
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func generateValues(ctx context.Context, counter chan int) {
	n := 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("context error:", err)
			}
			if err := context.Cause(ctx); err != nil {
				fmt.Println("context cause:", err)
			}
			return
		case counter <- n:
			n++
		}
	}
}

func main() {
	causeError := errors.New("timeout of goroutine")
	ctx, cancel := context.WithTimeoutCause(context.Background(), time.Second*5, causeError)
	defer cancel()
	counter := make(chan int)
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("executing stop function")
		// closing the channel
		close(counter)

	})
	defer stop()
	go generateValues(ctx, counter)
	for n := range counter {
		fmt.Println(n)
	}
	fmt.Println("done")
}
