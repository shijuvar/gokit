package main

import (
	"fmt"
)

func main() {
	// Declares a nil map
	var chapts map[int]string

	// Initialize map with make function
	chapts = make(map[int]string)

	// Add data as key/value pairs
	chapts[1] = "Beginning Go"
	chapts[2] = "Go Fundamentals"
	chapts[3] = "Structs and Interfaces"

	// Iterate over the elements of map using range
	for k, v := range chapts {
		fmt.Printf("Key: %d Value: %s\n", k, v)
	}

	// Declare and initialize map using map literal
	langs := map[string]string{
		"EL": "Greek",
		"EN": "English",
		"ES": "Spanish",
		"FR": "French",
		"HI": "Hindi",
	}

	// Delete an element
	delete(langs, "EL")
	// Lookout an element with key
	if lan, ok := langs["EL"]; ok {
		fmt.Println(lan)
	} else {
		fmt.Println("\nKey doesn't exists")
	}

	// Update an existing item of map
	langs["EN"] = "US English"
	for k, v := range langs {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
}
