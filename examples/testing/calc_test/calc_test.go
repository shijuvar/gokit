package calc_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"

	. "github.com/shijuvar/gokit/examples/testing/calc"
)

// Test case for the function Sum
func TestSum(t *testing.T) {

	testCases := []struct {
		name        string
		input       []int
		expected    int
		shouldError bool
	}{

		{
			name:     "sum with one digit numbers",
			input:    []int{7, 8, 5},
			expected: 20,
		},
		{
			name:     "sum with three digit numbers",
			input:    []int{100, 250, 400},
			expected: 750,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.input...)
			assert.Equal(t, tc.expected, result, fmt.Sprintf("Result: %d, Expected: %d", result, tc.expected))

		})
	}

}

// Test case for function Average
func TestAverage(t *testing.T) {
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {

		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}
