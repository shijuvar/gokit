package main

import (
	"fmt"
)

func main() {
	x := make([]int, 3, 5)
	x[0] = 10
	x[1] = 20
	x[2] = 30
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := make([]int, 3)
	y[0] = 10
	y[1] = 20
	y[2] = 30
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	z := []int{10, 20, 30}
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z1 := []int{0: 10, 2: 30}
	fmt.Println(len(z1))
	fmt.Println(cap(z1))

	x1 := []int{10, 20, 30}
	y1 := append(x1, 40, 50)
	fmt.Println(x1, y1)
}
