package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// ReadFile read the given file
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open file")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read the file")
	}
	return buf, nil
}

func main() {
	data, err := ReadFile("data.txt")
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(data))
}
