package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func main() {
	err := notFoundErrorHappens()
	// Is reports whether any error in err's chain matches target.
	if errors.Is(err, ErrNotFound) {
		fmt.Println(err)
	}
	// Unwraps the error value
	err = errors.Unwrap(err)
	fmt.Println(err)

	err = NotFoundError{
		Name: "File",
		Err:  err,
	}
	err = fileNotFoundError()
	var e NotFoundError
	// As finds the first error in err's chain that matches target,
	// and if so, sets target to that error value and returns true.
	if errors.As(err, &e) {
		fmt.Println(e)
	}
}

func notFoundErrorHappens() error {
	// Wrapping errors with %w
	// Wrapping an error with %w
	// makes it available to errors.Is and errors.As:
	return fmt.Errorf(
		"uh oh! something happened: %w", ErrNotFound,
	)
}

func fileNotFoundError() error {
	err := NotFoundError{
		Name: "File",
		Err:  errors.New("not found"),
	}
	return err
}

type NotFoundError struct {
	Name string
	Err  error
}

func (e NotFoundError) Error() string {
	return e.Name + ": not found"
}

func (e NotFoundError) Unwrap() error {
	return e.Err
}
