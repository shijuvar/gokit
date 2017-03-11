package main

import (
	"fmt"

	"github.com/shijuvar/go-recipes/ch01/strutils"
)

func main() {
	str1, str2 := "Golang", "gopher"
	// Convert to upper case
	fmt.Println("To Upper Case:", strutils.ToUpperCase(str1))

	// Convert to lower case
	fmt.Println("To Lower Case:", strutils.ToLowerCase(str1))

	// Convert first letter to upper case
	fmt.Println("To First Upper:", strutils.ToFirstUpper(str2))
}
