package main

import (
	"fmt"
	"iter"
	"sync"
)

type Slice[V any] []V

func (s Slice[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := range s {
			if !yield(s[i]) {
				return
			}
		}
	}
}
func (s Slice[V]) ReversedAll() iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// Pairs returns an iterator over successive pairs of values from seq.
func Pairs[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v1, ok1 := next()
			if !ok1 {
				return
			}
			v2, ok2 := next()
			// If ok2 is false, v2 should be the
			// zero value; yield one last pair.
			if !yield(v1, v2) {
				return
			}
			if !ok2 {
				return
			}
		}
	}
}
func Pull[V any](seq iter.Seq[V]) {
	// Pull converts the “push-style” iterator sequence seq into a “pull-style” iterator
	// accessed by the two functions next and stop.
	next, stop := iter.Pull(seq)
	defer stop()
	for {
		v, ok := next()
		if !ok {
			break
		}
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
}

// Print prints all elements in a sequence.
func Print2[V any](seq iter.Seq2[V, V]) {
	for k, v := range seq {
		fmt.Printf("%v:%v\t", k, v)
	}
	fmt.Println()
}

// Print prints all elements in a sequence.
func Print[V any](seq iter.Seq[V]) {
	for v := range seq {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
}
func main() {

	s := Slice[int]([]int{10, 20, 30, 40, 50})
	Print(s.All())
	Print(s.ReversedAll())
	Pull(s.All())
	Print2(Pairs(s.All()))
	fmt.Println("\niterateOverSMap")
	iterateOverSMap()
}

func Numbers(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i + 1) {
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
