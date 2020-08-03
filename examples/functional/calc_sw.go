package main

import (
	"fmt"
)

type SwCalculator struct {
	acc float64
}

func (c *SwCalculator) Do(input float64, op string) float64 {
	switch op {
	case "+":
		c.acc = c.acc + input
	case "-":
		c.acc = c.acc - input
	case "*":
		c.acc = c.acc * input
		// Cases for each supported operations
	}
	return c.acc
}
func main() {
	var c SwCalculator
	fmt.Println(c.Do(5, "+"))
	fmt.Println(c.Do(2, "*"))
	fmt.Println(c.Do(4, "-"))

}
