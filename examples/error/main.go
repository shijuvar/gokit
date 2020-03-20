package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile read the given file
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		// Errorf formats according to a format specifier
		// and returns the string as a value that satisfies error.
		// If the format specifier includes a %w verb with
		// an error operand, the returned error will implement
		// an Unwrap method returning the operand.
		return nil, fmt.Errorf("Failed to open file: %w", err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.New("Failed to read the file")
	}
	return buf, nil
}

func main() {
	data, err := ReadFile("data1.txt")
	if err != nil {
		fmt.Printf("%+v\n", err)
		werror := errors.Unwrap(err)
		if werror != nil {
			fmt.Printf("Wrap: %s\n", werror)
		}
		os.Exit(1)
	}
	fmt.Print(string(data))
}
