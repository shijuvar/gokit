package main

import (
	"fmt"
	"reflect"
)

type personer interface {
	getName()
}
type person struct {
	name string
}

func (p person) getName() {
	fmt.Println(p.name)
}
func (p person) toString() {
	fmt.Println("Person:", p.name)
	fmt.Println(100, "Global Logic,", []string{"fdfs", "sfsdfsd"})
}

// Empty interface: interface{}
func getType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println("Type:", t)
}

func main() {
	getType(100)
	getType("Golang")
	getType(person{name: "russ"})
	var p personer // interface type
	p = person{name: "shijuvar"}
	p.getName()
	// type assertion: assert into concrete type
	asserted := p.(person)
	asserted.getName()
	asserted.toString()

}
