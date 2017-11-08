package main

import (
	"fmt"
	"unsafe"
)

func main() {
	f := 100.00
	num := 100
	str := "Gopher"
	m := map[string]string{
		"EL": "Greek",
		"EN": "English",
		"ES": "Spanish",
		"FR": "French",
		"HI": "Hindi",
	}
	// unsafe.Sizeof function returns the size in bytes
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Sizeof(num))
	fmt.Println(unsafe.Sizeof(str))
	fmt.Println(unsafe.Sizeof(m))

}
