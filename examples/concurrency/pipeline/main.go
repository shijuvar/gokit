package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type fibvalue struct {
	input, value int
}

var wg sync.WaitGroup

func randomCounter(out chan<- int) {
	defer wg.Done()
	var random int
	for x := 0; x < 10; x++ {
		random = rand.Intn(50)
		out <- random
	}
	close(out)
}

func generateFibonacci(out chan<- fibvalue, in <-chan int) {
	defer wg.Done()
	var input float64
	for v := range in {
		input = float64(v)
		// Fibonacci using Binet's formula
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, input) - math.Pow(phi, input)) / math.Sqrt(5)
		out <- fibvalue{
			input: v,
			value: int(result),
		}
	}
	close(out)
}

func printFibonacci(in <-chan fibvalue) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("Fibonacci value of %d is %d\n", v.input, v.value)
	}
}

func main() {
	// Add 3 into WaitGroup Counter
	wg.Add(3)
	// Declare Channels
	randoms := make(chan int)
	fibs := make(chan fibvalue)
	// Launching 3 goroutines
	go randomCounter(randoms)
	go generateFibonacci(fibs, randoms)
	go printFibonacci(fibs)
	// Wait for completing all goroutines
	wg.Wait()
}
