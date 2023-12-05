package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	wg.Add(2)
	/*
		AfterFunc arranges to call f in its own goroutine after ctx is done (cancelled or timed out).
		If ctx is already done, AfterFunc calls f immediately in its own goroutine.
	*/
	stop := context.AfterFunc(ctx, func() {
		defer wg.Done()
		fmt.Println("clean-up code")
	})

	go doSomething(ctx, wg)
	cancel()
	// Calling the returned stop function stops the association of ctx with f
	stop()
	if !stop() {
		fmt.Println("stop func has been started in its own goroutine; or stop func was already stopped.")
	}
	wg.Wait()
}

func doSomething(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("Done received")
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	}
}
