package main

import (
	"errors"
	"fmt"
)

var err1 error = errors.New("err1")
var err2 error = errors.New("err2")

func main() {

	err := errors.Join(err1, err2)
	fmt.Println(err)
	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	}
	if errors.Is(err, err2) {
		fmt.Println("err is err2")
	}
}
