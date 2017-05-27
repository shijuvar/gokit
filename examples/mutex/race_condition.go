package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var sum = 0

func Sum(label string) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 25000; i++ {
			sum = sum + 1
		}

		fmt.Println("From "+label+":", sum)
	}()
}

func main() {
	processes := []string{"Process-1", "Process-2", "Process-3", "Process-4"}
	for _, p := range processes {
		Sum(p)
	}
	wg.Wait()
	fmt.Println("Final Sum:", sum)
}
