// Package calc provides a simple calculator
package calc

import "math"

// Sum returns sum of integer values
func Sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

// Average returns average of integer values
// The output provides a float64 value in two decimal points
func Average(nums ...int) float64 {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	result := float64(sum) / float64(len(nums))
	pow := math.Pow(10, float64(2))
	digit := pow * result
	round := math.Floor(digit)
	return round / pow

}
