package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	finish := make(chan struct{}) // Empty Struct occupies zero bytes
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		defer done.Done()
		select {
		case <-time.After(5 * time.Minute):
			fmt.Println("time out")
		case <-finish: // non-blocking when you close the channel
		}
	}()
	t0 := time.Now()
	sleep := rand.Int63n(20)
	time.Sleep(time.Duration(sleep) * time.Second)
	close(finish)
	done.Wait()
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
}
