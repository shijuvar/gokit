package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type in struct {
	input int
	op    string
}
type Test struct {
	name string
	in   in
	out  int
}

func TestCalcAdd(t *testing.T) {
	var tests = []Test{
		Test{
			name: "first input",
			in:   in{input: 5, op: "+"},
			out:  5,
		},
		Test{
			name: "second input",
			in:   in{input: 10, op: "+"},
			out:  15,
		},
		Test{
			name: "third input",
			in:   in{input: 35, op: "+"},
			out:  50,
		},
	}
	c := &Calculator{}
	for _, test := range tests {
		// Assertion with assert package
		assert.Equal(
			t,
			test.out,
			c.Do(test.in.input, test.in.op),
			fmt.Sprintf("%s: %v", test.name, test.in),
		)
	}
}

// TestCalcMultipleOps tests
func TestCalcMultipleOps(t *testing.T) {
	var tests = []Test{
		Test{
			name: "first input for addition",
			in:   in{input: 10, op: "+"},
			out:  10,
		},
		Test{
			name: "second input for multiplication",
			in:   in{input: 5, op: "*"},
			out:  50,
		},
		Test{
			name: "third input for subtraction",
			in:   in{input: 30, op: "-"},
			out:  20,
		},
	}
	c := &Calculator{}
	for _, test := range tests {
		// Assertion with assert package
		assert.Equal(
			t,
			test.out,
			c.Do(test.in.input, test.in.op),
			fmt.Sprintf("%s: %v", test.name, test.in),
		)
	}
}

// Sub-level tests with table test
func TestCalcMultipleOpsWithSub(t *testing.T) {
	var tests = []Test{
		Test{
			name: "first input for addition",
			in:   in{input: 10, op: "+"},
			out:  10,
		},
		Test{
			name: "second input for multiplication",
			in:   in{input: 5, op: "*"},
			out:  50,
		},
		Test{
			name: "third input for subtraction",
			in:   in{input: 30, op: "-"},
			out:  20,
		},
	}
	c := &Calculator{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(
				t,
				test.out,
				c.Do(test.in.input, test.in.op),
				fmt.Sprintf("%s: %v", test.name, test.in),
			)
		})
	}
}
