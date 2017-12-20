package main

import (
	"fmt"
)

func main() {
	teams := []string{"Dev", "QA", "Ops"}
	manager := "Keerthi"

	team := struct {
		teams   []string
		manager string
	}{
		teams,
		manager,
	}
	fmt.Printf("%+v", team)
}
