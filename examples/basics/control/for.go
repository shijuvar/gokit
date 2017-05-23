package main

import "fmt"

func main() {
	sum()
	sum1()
}
func sum() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func sum1() {
	sum := 1
	for sum < 50 {
		sum += sum
	}
	fmt.Println(sum)
}
