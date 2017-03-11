package main

import (
	"fmt"
	"reflect"
)

func main() {
	// type conversion: dealing with a type
	// The expression T(v) converts the value v to the type T
	i := 100
	f := float64(i)
	fmt.Println(reflect.TypeOf(f))
	// type assertion: dealing with an interface
	// A type assertion provides access to an interface value's underlying concrete value
	var x interface{} = float64(100)
	y := x.(float64)
	fmt.Println(reflect.TypeOf(y))
}
