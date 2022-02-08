package calculator

type Calculator struct {
	acc int
}

func (c *Calculator) Do(input int, op string) int {
	switch op {
	case "+":
		c.acc = c.acc + input
	case "-":
		c.acc = c.acc - input
	case "*":
		c.acc = c.acc * input
	}
	return c.acc
}
