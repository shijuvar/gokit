package main

import "fmt"

func getType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("unknown")
	}
}
func main() {
	getType(25)
	getType(true)
	getType("Gopher")
}
