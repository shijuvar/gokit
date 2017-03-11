package model

import "errors"

// ErrorEmailExists is an error value for duplicate email id
var ErrorEmailExists = errors.New("Email Id is exists")

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
