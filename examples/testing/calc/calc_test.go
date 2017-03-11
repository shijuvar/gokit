package calc

import (
	"fmt"
	"testing"
	"time"
)

// Test case for the function Sum
func TestSum(t *testing.T) {
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {

		t.Errorf("Result: %d, Expected: %d", result, expected)
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

// Benchmark for function Sum
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(7, 8, 10)
	}
}

// Benchmark for function Average
func BenchmarkAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average(7, 8, 10)
	}
}

func TestLongRun(t *testing.T) {
	// Checks whether the short flag is provided
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	// Long running implementation goes here
	time.Sleep(5 * time.Second)
}

// Test case for the function Sum to be executed in parallel
func TestSumInParallel(t *testing.T) {
	t.Parallel()
	// Delaying 1 second for the sake of demonstration
	time.Sleep(1 * time.Second)
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {

		t.Errorf("Result: %d, Expected: %d", result, expected)
	}

}

// Test case for the function Sum to be executed in parallel
func TestAverageInParallel(t *testing.T) {
	t.Parallel()
	// Delaying 1 second for the sake of demonstration
	time.Sleep(2 * time.Second)
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {

		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}

// Example code for function Sum
func ExampleSum() {
	fmt.Println(Sum(7, 8, 10))
	// Output: 25
}

// Example code for function Average
func ExampleAverage() {
	fmt.Println(Average(7, 8, 10))
	// Output: 8.33
}
