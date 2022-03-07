package main

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var errNotFound = errors.New("No records found")

type note struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	NoteID      string             `json:"noteid,omitempty" bson:"noteid,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedOn   time.Time          `json:"createdon,omitempty" bson:"createdon,omitempty"`
}

// CRUD interface
type repository interface {
	create(note) error
	update(string, note) error
	delete(string) error
	getById(string) (note, error)
	getAll() ([]note, error)
}
