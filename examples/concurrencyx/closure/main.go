package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var values []string = []string{"a", "b", "c"}

func main() {
	errOnClosure()
	closureWithParam()
	closureWithNewVar()
	closureWithWaitGroup()
	closureWithErrGroup()
}

func errOnClosure() {
	fmt.Println("errOnClosure")
	done := make(chan struct{})
	defer close(done)
	// data race: the variable v is shared by len(values) of goroutines.
	// each iteration of the loop uses the same instance of the variable v,
	// so each closure shares that single variable.
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- struct{}{}
		}()
	}

	//wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

func closureWithNewVar() {
	fmt.Println("closureWithNewVar")
	done := make(chan struct{})
	defer close(done)
	for _, v := range values {
		v := v // create a new 'v'.
		go func() {
			fmt.Println(v)
			done <- struct{}{}
		}()
	}

	//wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

// closureWithParam uses a local variable and pass the string
// as a parameter when starting the goroutine
func closureWithParam() {
	fmt.Println("closureWithParam")
	done := make(chan struct{})
	defer close(done)
	for _, v := range values {
		go func(s string) { // Use a local variable.
			fmt.Println(s)
			done <- struct{}{}
		}(v)
	}

	//wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

func closureWithWaitGroup() {
	fmt.Println("closureWithWaitGroup")
	var wg sync.WaitGroup
	for _, v := range values {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			fmt.Println(s)
		}(v)
	}
	wg.Wait()
}
func closureWithErrGroup() {
	fmt.Println("closureWithErrGroup")
	g := new(errgroup.Group)
	for _, v := range values {
		v := v
		g.Go(func() error {
			fmt.Println(v)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Error from ErrGroup:", err)
	}
}
