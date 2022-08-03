package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
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
	fmt.Println("closureWithVar")
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

func closureWithParam() {
	fmt.Println("closureWithParam")

	done := make(chan struct{})
	defer close(done)
	for _, v := range values {
		v := v
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

func closureWithWaitGroup() {
	fmt.Println("closureWithWaitGroup")
	var wg sync.WaitGroup
	for _, v := range values {
		v := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
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
