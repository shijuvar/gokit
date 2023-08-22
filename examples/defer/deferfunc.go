package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return io.ReadAll(f)
}

func main() {
	f, _ := ReadFile("test.txt")
	fmt.Println(string(f))
}
