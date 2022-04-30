package memstore

import (
	// internal
	"errors"
	"github.com/shijuvar/gokit/examples/http/restapi/model"
	"time"

	// external
	"github.com/gofrs/uuid"
)

// inmemoryRepository provides concrete implementation
// for repository interface
type inmemoryRepository struct {
	noteStore map[string]model.Note
}

func NewInmemoryRepository() (model.Repository, error) {
	return &inmemoryRepository{
		noteStore: make(map[string]model.Note),
	}, nil
}

func (i *inmemoryRepository) isNoteTitleExists(title string) bool {
	for _, v := range i.noteStore {
		if v.Title == title {
			return true
		}
	}
	return false
}
func (i *inmemoryRepository) Create(n model.Note) error {
	if _, ok := i.noteStore[n.NoteID]; ok {
		return errors.New("NoteID exists")
	}
	if i.isNoteTitleExists(n.Title) {
		return model.ErrNoteExists
	}
	n.CreatedOn = time.Now()
	// Create a Version 4 UUID.
	uid, _ := uuid.NewV4()
	n.NoteID = uid.String()
	i.noteStore[n.NoteID] = n
	return nil
}

func (i *inmemoryRepository) Update(id string, n model.Note) error {
	if _, ok := i.noteStore[id]; !ok {
		return errors.New("NoteID doesn't exist")
	}
	n.CreatedOn = time.Now()
	i.noteStore[id] = n
	return nil
}

func (i *inmemoryRepository) Delete(id string) error {
	if _, ok := i.noteStore[id]; !ok {
		return errors.New("NoteID doesn't exist")
	}
	delete(i.noteStore, id)
	return nil
}
func (i *inmemoryRepository) GetById(id string) (model.Note, error) {
	if v, ok := i.noteStore[id]; !ok {
		return model.Note{}, errors.New("NoteID doesn't exist")
	} else {
		return v, nil
	}

}

func (i *inmemoryRepository) GetAll() ([]model.Note, error) {
	notes := make([]model.Note, 0)
	for _, v := range i.noteStore {
		notes = append(notes, v)
	}
	return notes, nil
}
