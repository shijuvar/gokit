package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// in organises input value for Calc
type in struct {
	input float64
	op    string
}

// Test organises values for tests
type Test struct {
	name        string
	in          in
	out         float64
	shouldError bool
}

// TestCalcAdd tests the behavior of Calc with Add operation
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
	var result float64
	for _, test := range tests {

		result = c.Do(test.in.input, test.in.op)
		//if test.out != result {
		//	t.Errorf(fmt.Sprintf("%s: %v", test.name, test.in))
		//}
		// Assertion with assert package
		assert.Equal(
			t,
			test.out,
			result,
			fmt.Sprintf("%s: %v", test.name, test.in),
		)
	}
}

// TestCalcMultipleOps tests the behavior of Calc with multiple operations
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
		Test{
			name: "fourth input for subtraction",
			in:   in{input: 4, op: "/"},
			out:  5,
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

// TestCalcMultipleOpsWithSubTest tests the behavior of Calc with multiple operations
// with sub-level tests
func TestCalcMultipleOpsWithSubTest(t *testing.T) {
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
		Test{
			name: "fourth input for subtraction",
			in:   in{input: 4, op: "/"},
			out:  5,
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

// BenchmarkCalc benchmarks Calc
func BenchmarkCalc(b *testing.B) {
	test := Test{
		in:  in{input: 5, op: "+"},
		out: 5,
	}
	for i := 0; i < b.N; i++ {
		c := &Calculator{}
		c.Do(test.in.input, test.in.op)
	}
}

// BenchmarkCalcMultipleOpsWithSubTest benchmarks with sub-level benchmark
func BenchmarkCalcMultipleOpsWithSubTest(b *testing.B) {
	var tests = []Test{
		Test{
			name: "input for addition",
			in:   in{input: 10, op: "+"},
			out:  10,
		},
		Test{
			name: "input for multiplication",
			in:   in{input: 5, op: "*"},
			out:  50,
		},
		Test{
			name: "input for subtraction",
			in:   in{input: 30, op: "-"},
			out:  20,
		},
		Test{
			name: "third input for subtraction",
			in:   in{input: 4, op: "/"},
			out:  5,
		},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := &Calculator{}
				c.Do(test.in.input, test.in.op)
			}
		})
	}
}
