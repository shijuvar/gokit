package main

import (
	"cmp"
	"fmt"
)

func main() {
	var x, y int
	m := min(x) // m == x
	fmt.Println(m)
	m = min(x, y) // m is the smaller of x and y
	fmt.Println(m)
	m = max(x, y, 10) // m is the larger of x and y but at least 10
	fmt.Println(m)
	c := max(1, 2.0, 10) // c == 10.0 (floating-point kind)
	fmt.Println(c)
	f := max(0, float32(x)) // type of f is float32
	fmt.Println(f)
	t := max("", "foo", "bar") // t == "foo" (string kind)
	fmt.Println(t)

	/*
		The built-in function clear takes an argument of map, slice, or type parameter type,
		and deletes or zeroes out all elements.
			If the type of the argument to clear is a type parameter,
			all types in its type set must be maps or slices,
			and clear performs the operation corresponding to the actual type argument.
			If the map or slice is nil, clear is a no-op.
	*/
	kv := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
	}
	clear(kv)            // deletes all entries, resulting in an empty map (len(m) == 0)
	fmt.Println(len(kv)) // output: 0
	s := []int{10, 20, 30, 40, 50}
	clear(s)            // sets all elements up to the length of s to the zero value of T
	fmt.Println(len(s)) // output: 5
	fmt.Println(s[0])   // output: 0

	a, b := 10, 20
	v := cmp.Compare(a, b)
	fmt.Println(v)
	v = cmp.Compare(a, 10)
	fmt.Println(v)

}
