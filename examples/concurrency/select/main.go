package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type (
	fibvalue struct {
		input, value int
	}
	squarevalue struct {
		input, value int
	}
)

func generateSquare(sqrs chan<- squarevalue) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		num := rand.Intn(50)
		sqrs <- squarevalue{
			input: num,
			value: num * num,
		}
	}
	close(sqrs)
}
func generateFibonacci(fibs chan<- fibvalue) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		num := float64(rand.Intn(50))
		// Fibonacci using Binet's formula
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, num) - math.Pow(phi, num)) / math.Sqrt(5)
		fibs <- fibvalue{
			input: int(num),
			value: int(result),
		}
	}
	close(fibs)
}
func printValues(fibs <-chan fibvalue, sqrs <-chan squarevalue) {
	defer wg.Done()
	for {
		if fibs == nil && sqrs == nil {
			break
		}
		select {
		case fib, ok := <-fibs:
			if ok {
				fmt.Printf("Fibonacci value of %d is %d\n", fib.input, fib.value)
			} else {
				fibs = nil
			}
		case sqr, ok := <-sqrs:
			if ok {
				fmt.Printf("Square value of %d is %d\n", sqr.input, sqr.value)
			} else {
				sqrs = nil
			}
		}
	}
}

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {
	wg.Add(3)
	// Create Channels
	fibs := make(chan fibvalue)
	sqrs := make(chan squarevalue)
	// Launching 3 goroutines
	go generateFibonacci(fibs)
	go generateSquare(sqrs)
	go printValues(fibs, sqrs)
	// Wait for completing all goroutines
	wg.Wait()
}
