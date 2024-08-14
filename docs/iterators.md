## Iterators 
Iterators was introduced in Go 1.23. An iterator is a function that passes successive elements of a sequence to a callback function, conventionally named yield. The function stops either when the sequence is finished or when yield returns false, indicating to stop the iteration early. This package defines Seq and Seq2 (pronounced like seek—the first syllable of sequence) as shorthands for iterators that pass 1 or 2 values per sequence element to yield:

```go
type (
	Seq[V any]     func(yield func(V) bool)
	Seq2[K, V any] func(yield func(K, V) bool)
)
```

Seq2 represents a sequence of paired values, conventionally key-value or index-value pairs.

Yield returns true if the iterator should continue with the next element in the sequence, false if it should stop.

Here's an iterator with Seq[V any]: 
```go
// Reversed returns an iterator that loops over a slice in reverse order.
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
```


Iterator functions are most often called by a range loop, as in:

```go
func main() {
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
```