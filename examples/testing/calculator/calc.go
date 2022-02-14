package calculator

type Calculator struct {
	acc float64
}

func (c *Calculator) Do(input float64, op string) float64 {
	switch op {
	case "+":
		c.acc = c.acc + input
	case "-":
		c.acc = c.acc - input
	case "*":
		c.acc = c.acc * input
	case "/":
		c.acc = c.acc / input
	}
	return c.acc
}
