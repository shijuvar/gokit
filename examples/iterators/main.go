package main

import (
	"fmt"
	"iter"
	"sync"
)

func main() {
	iterateOverSMap()
	iterateWithSeq()
}
func iterateWithSeq() {
	s := []int{1, 2, 3, 4, 5}
	PrintAll(Reversed(s))
}

// PrintAll prints all elements in a sequence.
func PrintAll[V any](s iter.Seq[V]) {
	for v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// Reversed returns an iterator that loops over a slice in reverse order.
/*
An iterator is a function that passes successive elements of a sequence to a callback function,
conventionally named yield.
The function stops either when the sequence is finished or when yield returns false,
indicating to stop the iteration early.
This package defines Seq and Seq2 (pronounced like seek—the first syllable of sequence)
as shorthands for iterators that pass 1 or 2 values per sequence element to yield:
*/
// pushing values to the yield function.
func Reversed[V any](s []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			// Yield returns true if the iterator should continue with the next element in the sequence,
			// false if it should stop.
			if !yield(s[i]) {
				return
			}
		}
	}
}
func iterateOverSMap() {
	// sync.Map type is part of the sync package and provides built-in
	// synchronization to prevent race conditions without the explicit use of mutexes.
	var m sync.Map

	m.Store("alice", 11)
	m.Store("bob", 12)
	m.Store("cindy", 13)

	/*
		go 1.22 explicit call
		m.Range(func(key, val any) bool {
			fmt.Println(key, val)
			return true
		})
	*/
	// go 1.23 - becomes implicit call
	for key, val := range m.Range {
		fmt.Println(key, val)
	}
}
