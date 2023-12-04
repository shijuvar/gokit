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
	/*
		WithCancel returns a copy of parent with a new Done channel.
		The returned context's Done channel is closed when the returned cancel function is called or
		when the parent context's Done channel is closed, whichever happens first.

		Canceling this context releases resources associated with it,
		so code should call cancel as soon as the operations running in this Context complete.
	*/

	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	/*
		WithCancelCause behaves like WithCancel but returns a CancelCauseFunc instead of a CancelFunc.
		Calling cancel with a non-nil error (the "cause") records that error in ctx;
		it can then be retrieved using Cause(ctx).
		Calling cancel with nil sets the cause to Canceled.
	*/
	ctx, cancel := context.WithCancelCause(context.Background())
	causeError := errors.New("this is a leaking goroutine")

	counter := make(chan int)
	go generateValues(ctx, counter)
	for n := range counter {
		if n == 10 {
			cancel(causeError)
		}
		fmt.Println(n)
	}
	fmt.Println("done")
}
