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
		return nil, err
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
		os.Exit(1)
	}
	fmt.Print(string(data))
}
