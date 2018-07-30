package main

import (
	"fmt"
)

func main() {

	team := struct {
		teams   []string
		manager string
	}{
		[]string{"Dev", "QA", "Ops"},
		"Keerthi",
	}
	fmt.Printf("%+v\n", team)

}
