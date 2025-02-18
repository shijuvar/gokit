package domain

type User[T any] struct {
	ID   T
	Name string
}

type Customer struct {
	Name  string
	Email string
}
