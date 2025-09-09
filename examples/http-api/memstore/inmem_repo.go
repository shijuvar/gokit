package memstore

import (
	// internal
	"errors"
	"time"

	// external
	"github.com/gofrs/uuid"

	"github.com/shijuvar/gokit/examples/http-api/model"
)

// inmemoryRepository provides concrete implementation for repository interface
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
func (i *inmemoryRepository) Create(n model.Note) (string, error) {
	if _, ok := i.noteStore[n.NoteID]; ok {
		return "", errors.New("NoteID exists")
	}
	if i.isNoteTitleExists(n.Title) {
		return "", model.ErrNoteExists
	}
	n.CreatedOn = time.Now()
	// Create a Version 4 UUID.
	uid, _ := uuid.NewV4()
	n.NoteID = uid.String()
	i.noteStore[n.NoteID] = n
	return n.NoteID, nil
}

func (i *inmemoryRepository) Update(id string, n model.Note) error {
	if _, ok := i.noteStore[id]; !ok {
		return model.ErrNoteNotExists
	}

	n.NoteID = id
	n.CreatedOn = time.Now()
	i.noteStore[id] = n
	return nil
}

func (i *inmemoryRepository) Delete(id string) error {
	if _, ok := i.noteStore[id]; !ok {
		return model.ErrNoteNotExists
	}
	delete(i.noteStore, id)
	return nil
}
func (i *inmemoryRepository) GetById(id string) (model.Note, error) {
	if v, ok := i.noteStore[id]; !ok {
		return model.Note{}, model.ErrNoteNotExists
	} else {
		return v, nil
	}

}

func (i *inmemoryRepository) GetAll() ([]model.Note, error) {
	if len(i.noteStore) == 0 {
		return nil, model.ErrNotFound
	}
	notes := make([]model.Note, 0)
	for _, v := range i.noteStore {
		notes = append(notes, v)
	}
	return notes, nil
}
