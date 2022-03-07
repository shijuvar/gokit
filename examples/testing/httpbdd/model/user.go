package model

import "errors"

//go:generate mockgen -destination=../withgomock/mock_store.go -package=withgomock github.com/shijuvar/gokit/examples/testing/httpbdd/model UserStore

// ErrorEmailExists is an error value for duplicate email id
var ErrorEmailExists = errors.New("Email NoteID is exists")

// User model
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// UserStore provides a contract for Data Store for User entity
type UserStore interface {
	GetUsers() []User
	AddUser(User) error
}
