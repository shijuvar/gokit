package main

import (
	"fmt"

	fav "github.com/shijuvar/gokit/examples/basics/lib"
)

func main() {
	// Print default favorite packages
	fmt.Println("****** Default favorite packages ******\n")
	fav.PrintFavorites()
	// Add couple of favorites
	fav.Add("github.com/dgrijalva/jwt-go")
	fav.Add("github.com/onsi/ginkgo")
	fmt.Println("\n****** All favorite packages ******\n")
	fav.PrintFavorites()
	count := len(fav.GetAll())
	fmt.Printf("Total packages in the favorite list:%d", count)
}
