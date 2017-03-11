package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30, 40, 50}
	y := x[1:3]
	fmt.Println("y:", y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	z := x[:3]
	fmt.Println("z:", z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	x1 := x[:]
	fmt.Println("x1:", x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))

	x1[4] = 75
	fmt.Println("x:", x)
	fmt.Println("x1:", x1)

}
