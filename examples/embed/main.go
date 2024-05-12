package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var fileString string

//go:embed hello.txt
var fileByte []byte

//go:embed hello.txt
var f embed.FS

func main() {
	fmt.Println(fileString)
	fmt.Println(string(fileByte))
	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))
}
