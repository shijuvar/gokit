package main

import (
	"fmt"

	domain "gokit/examples/generic-alias/domain-v2"
)

func main() {

	user := domain.User[string]{
		ID:   "u101",
		Name: "john",
	}
	fmt.Println(user)
}
