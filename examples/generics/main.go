package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Number interface {
	~int | ~int32 | ~int64 | ~float64
}
// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	fmt.Printf("Max with integer: %v\n",Max([]num{10, 15, 4, 25, 16, 18, 2}))
	fmt.Printf("Max with float: %v\n",Max([]float64{6.2, 4.1, 6.2, 9.6, 8.2, 1.5, 4.7}))
	fmt.Printf("Contains with string: %v\n", Contains([]string{"one", "two", "three"},"two"))
	fmt.Printf("Contains with int: %v\n", Contains([]int{100, 200, 300},50))

	randomNumbers:=make([]int, 5, 5)
	for i:=0;i<5; i++ {
		randomNumbers[i]=rand.Intn(25)
	}
   fibvalues:= MapSlice(randomNumbers,findFibonacci)
   fmt.Println(fibvalues)
   sqrs:= MapSlice(randomNumbers, func(num int) int {
	   return num * num
   })
	fmt.Println(sqrs)

}

type fibvalue struct {
	input, value int
}
func findFibonacci(num int) fibvalue {
	input := float64(num)
	// Fibonacci using Binet's formula
	Phi := (1 + math.Sqrt(5)) / 2
	phi := (1 - math.Sqrt(5)) / 2
	result := (math.Pow(Phi, input) - math.Pow(phi, input)) / math.Sqrt(5)
	return fibvalue {
		input: num,
		value: int(result),
	}
}
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type num int
func Max[T Number](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}
func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func MapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}