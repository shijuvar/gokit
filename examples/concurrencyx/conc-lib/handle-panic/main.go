package main

import "github.com/sourcegraph/conc"

func main() {
	var wg conc.WaitGroup
	wg.Go(func() {
		panic("just panicking")

	})
	// panics with a nice stacktrace
	wg.Wait()
}
