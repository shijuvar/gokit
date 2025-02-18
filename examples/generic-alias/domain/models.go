package domain

type User[T any] struct {
	ID   T
	Name string
}
