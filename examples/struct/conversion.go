package main

import (
	"fmt"
)

type person struct {
	firstName, lastName string
}

type employee struct {
	firstName, lastName string
}

type developer struct {
	firstName, lastName string
	skills              []string
}

func main() {
	p := person{"Shiju", "Varghese"}
	e := employee(p)
	//d := developer(p) // This is not possible
	fmt.Println(e)

}
