package main

import (
	"fmt"
	"sync"

	"github.com/sourcegraph/conc"
)

func main() {
	fmt.Println("with std lib")
	withStdLib()
	fmt.Println("with conc")
	withConc()
}

func withConc() {
	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Go(func() {
			fmt.Println(i)
		})
	}
	wg.Wait()
}
func withStdLib() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
