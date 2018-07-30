package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s struct{} // empty struct consumes zero bytes,
	fmt.Println(unsafe.Sizeof(s))
	var x [100000]struct{}
	fmt.Println(unsafe.Sizeof(x))

	m := make(map[string]struct{})
	m["a"] = struct{}{} // assign empty struct value

	if _, ok := m["a"]; ok {
		fmt.Println("Value exists")
	} else {
		fmt.Println("Value doesn't exists")
	}

	for i := range N(5) {
		fmt.Println(i)
	}

}

// N returns a slice of n 0-sized elements
func N(n int) []struct{} {
	return make([]struct{}, n)
}
