package main

import "fmt"

type GenType[T any] struct {
	Data []T
}

// Data is a generic struct with two generic types
type Data[K comparable, V any] struct {
	Key   K
	Value V
}

// getSliceData returns a slice of Data[K, V]
func getSliceData[K comparable, V any](m map[K]V) []Data[K, V] {
	// define a slice with Data type passing K, V type parameters
	d := make([]Data[K, V], len(m))
	i := 0
	for k, v := range m {
		// creating value of generic type of Data
		newData := Data[K, V]{}
		newData.Key = k
		newData.Value = v
		d[i] = newData
		i++
	}
	return d
}

func main() {

	// passing int as type parameter for GenType
	modelInt := GenType[int]{Data: []int{1, 2, 3, 4}}
	fmt.Println(modelInt.Data) // [1 2 3 4]

	// passing string as type parameter for GenType
	modelStr := GenType[string]{Data: []string{"a", "b", "c", "d"}}
	fmt.Println(modelStr.Data) // [a b c d]

	data := getSliceData(map[string]string{
		"go":  "Go programming language",
		"rs":  "Rust programming language",
		"zig": "Zig programming language",
	})
	for _, v := range data {
		fmt.Printf("Key: %s Value:%s\n", v.Key, v.Value)
	}

	data1 := getSliceData(map[int]string{
		1: "Go programming language",
		2: "Rust programming language",
		3: "Zig programming language",
	})
	for _, v := range data1 {
		fmt.Printf("Key: %d Value:%s\n", v.Key, v.Value)
	}
}
