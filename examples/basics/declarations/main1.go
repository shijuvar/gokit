package main

import "fmt"

// Declare constant
const Title string = "Person Details"

// Declare package variable
var Country string = "USA"

func main() {
	var fname, lname string = "Shiju", "Varghese"
	var age int = 35
	// Print constant variable
	fmt.Println(Title)
	// Print local variables
	fmt.Println("First Name:", fname)
	fmt.Println("Last Name:", lname)
	fmt.Println("Age:", age)
	// Print package variable
	fmt.Println("Country:", Country)
}
