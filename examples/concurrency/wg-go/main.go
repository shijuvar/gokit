// This sample program demonstrates how to create goroutines
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// WaitGroup is used to wait for the program to finish goroutines.
	wg := new(sync.WaitGroup)
	fmt.Println("Start Goroutines")
	// Launch functions as goroutines
	wg.Go(addTable)
	wg.Go(multiTable)
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

// addTable prints addition table for 1 to 10
func addTable() {
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Addition Table for:", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d+%d=%d\t", i, j, i+j)
		}
		fmt.Println("\n")
	}
}

// multiTable prints multiplication table for 1 to 10
func multiTable() {
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Multiplication Table for:", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println("\n")
	}
}
