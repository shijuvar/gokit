package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var f embed.FS

func main() {
	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))
}
