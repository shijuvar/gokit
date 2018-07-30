package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	finish := make(chan struct{})
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		defer done.Done()
		select {
		case <-time.After(5 * time.Minute):
			fmt.Println("time out")
		case <-finish:
		}
	}()
	t0 := time.Now()
	sleep := rand.Int63n(20)
	time.Sleep(time.Duration(sleep) * time.Second)
	close(finish)
	done.Wait()
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
}
