package main

import "fmt"

func main() {
	fmt.Println(truncate([]int{10, 20, 30, 40, 50}, 2))
	fmt.Println(truncateWithCopy([]int{10, 20, 30, 40, 50}, 2))
	fmt.Println(truncateWithAppend([]int{10, 20, 30, 40, 50}, 2))
	fmt.Println(genericTruncate([]int{10, 20, 30, 40, 50}, 2))
	fmt.Println(genericTruncate([]string{"A", "B", "C", "D", "E"}, 2))
	fmt.Println(genericTruncate([]float64{10.65, 20.00, 30.50, 40.55, 50.75}, 2))

}

func truncate(sl []int, i int) []int {
	// Remove the element at index i from sl.
	sl[i] = sl[len(sl)-1] // Copy last element to index i.
	sl[len(sl)-1] = 0     // Erase last element (write zero value).
	sl = sl[:len(sl)-1]   // Truncate slice.
	return sl
}
func truncateWithCopy(sl []int, i int) []int {
	// Remove the element at index i from sl.
	copy(sl[i:], sl[i+1:]) // Shift sl[i+1:] left one index.
	sl[len(sl)-1] = 0      // Erase last element (write zero value).
	sl = sl[:len(sl)-1]    // Truncate slice.
	return sl
}
func truncateWithAppend(sl []int, i int) []int {
	sl = append(sl[:i], sl[i+1:]...)
	return sl
}

type valuetype interface {
	~int | ~int32 | ~int64 | ~float64 | ~string
}

func genericTruncate[V valuetype](sl []V, i int) []V {
	sl = append(sl[:i], sl[i+1:]...)
	return sl
}
