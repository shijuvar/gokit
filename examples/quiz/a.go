package main

import "fmt"

type Text struct {
	s string
}

func (x *Text) String() string {
	if x == nil {
		return "<empty>"
	}
	return x.s
}

func show(f func() string) {
	fmt.Println("Result:", f())
}

func main() {
	var x *Text
	show(x.String)
}
