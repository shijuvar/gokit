package model

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("No records found")
var ErrNoteExists = errors.New("Note title exists")

type Note struct {
	NoteID      string    `json:"noteid,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon,omitempty"`
}

// CRUD interface
type Repository interface {
	Create(Note) error
	Update(string, Note) error
	Delete(string) error
	GetById(string) (Note, error)
	GetAll() ([]Note, error)
}
