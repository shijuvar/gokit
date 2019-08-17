package main

import (
	"fmt"
	"sync"
)

var (
	doOnce           sync.Once
	count, singleton int
)

func main() {
	Do()
	Do()
	Do()
}

func Do() {
	// if once.Do(f) is called multiple times, only the first call will invoke f,
	// even if f has a different value in each invocation.
	// once.Do(f) is concurrency safe
	doOnce.Do(func() {
		singleton++
	})
	count++
	fmt.Println("Singleton: ", singleton)
	fmt.Println("Count: ", count)
}
