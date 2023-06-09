package main

import (
	"fmt"
	"runtime/debug"
)

type caughtPanicError struct {
	val   any
	stack []byte
}

func (e *caughtPanicError) Error() string {
	return fmt.Sprintf(
		"panic: %q\n%s",
		e.val,
		string(e.stack),
	)
}

func main() {
	done := make(chan error)
	go func() {
		defer func() {
			if v := recover(); v != nil {
				done <- &caughtPanicError{
					val:   v,
					stack: debug.Stack(),
				}
			} else {
				done <- nil
			}
		}()
		panic("just pancking")
	}()
	err := <-done
	if err != nil {
		panic(err)
	}
}
