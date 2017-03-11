package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func main() {
	f, _ := ReadFile("test.txt")
	fmt.Println("%s", f)
	fmt.Println(string(f))
}
