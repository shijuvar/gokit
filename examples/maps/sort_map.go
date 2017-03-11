package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initialize map with make function
	chapts := make(map[int]string)

	// Add data as key/value pairs
	chapts[1] = "Beginning Go"
	chapts[2] = "Go Fundamentals"
	chapts[3] = "Structs and Interfaces"
	for k, v := range chapts {
		fmt.Println(k, v)
	}

	// Slice for specifying the order of the map
	var keys []int
	// Appending keys of the map
	for k := range chapts {
		keys = append(keys, k)
	}
	// Ints sorts a slice of ints in increasing order.
	sort.Ints(keys)
	// Iterate over the map with an order
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", chapts[k])
	}
}
