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
	/*
		Package errgroup provides synchronization, error propagation, and Context cancelation
		for groups of goroutines working on subtasks of a common task.
		A Group is a collection of goroutines working on subtasks that are part of the same overall task.
	*/
	g := new(errgroup.Group)
	for _, v := range values {
		v := v
		/*
			Go calls the given function in a new goroutine.
			The first call to return a non-nil error cancels the group's ctx-cancel,
			if the group was created by calling WithContext.
			The error will be returned by Wait.
		*/
		g.Go(func() error {
			fmt.Println(v)
			return nil
		})
	}
	/*
		Wait blocks until all function calls from the Go method have returned,
		then returns the first non-nil error (if any) from them.
	*/
	if err := g.Wait(); err != nil {
		fmt.Println("Error from ErrGroup:", err)
	}
}
