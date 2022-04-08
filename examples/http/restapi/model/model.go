package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks/mock_repository.go -package=mocks github.com/shijuvar/gokit/examples/http/restapi/model Repository

var ErrNotFound = errors.New("No records found")
var ErrNoteExists = errors.New("Note title exists")

type Note struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	NoteID      string             `json:"noteid,omitempty" bson:"noteid,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedOn   time.Time          `json:"createdon,omitempty" bson:"createdon,omitempty"`
}

// CRUD interface
type Repository interface {
	Create(Note) error
	Update(string, Note) error
	Delete(string) error
	GetById(string) (Note, error)
	GetAll() ([]Note, error)
}
