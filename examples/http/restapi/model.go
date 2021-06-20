package main

import "time"

type note struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

// CRUD interface
type repository interface {
	create(note) error
	update(string, note) error
	delete(string) error
	getById(string) (note, error)
	getAll() ([]note, error)
}
