package main

import (
	// internal
	"errors"
	"time"

	// external
	"github.com/gofrs/uuid"
)

// inmemoryRepository provides concrete implementation
// for repository interface
type inmemoryRepository struct {
	noteStore map[string]note
}

func newInmemoryRepository() (repository, error) {
	return &inmemoryRepository{
		noteStore: make(map[string]note),
	}, nil
}
func (i *inmemoryRepository) create(n note) error {
	if _, ok := i.noteStore[n.NoteID]; ok {
		return errors.New("NoteID exists")
	}
	n.CreatedOn = time.Now()
	// Create a Version 4 UUID.
	uid, _ := uuid.NewV4()
	n.NoteID = uid.String()
	i.noteStore[n.NoteID] = n
	return nil
}

func (i *inmemoryRepository) update(id string, n note) error {
	if _, ok := i.noteStore[id]; !ok {
		return errors.New("NoteID doesn't exist")
	}
	n.CreatedOn = time.Now()
	i.noteStore[id] = n
	return nil
}

func (i *inmemoryRepository) delete(id string) error {
	if _, ok := i.noteStore[id]; !ok {
		return errors.New("NoteID doesn't exist")
	}
	delete(i.noteStore, id)
	return nil
}
func (i *inmemoryRepository) getById(id string) (note, error) {
	if v, ok := i.noteStore[id]; !ok {
		return note{}, errors.New("NoteID doesn't exist")
	} else {
		return v, nil
	}

}

func (i *inmemoryRepository) getAll() ([]note, error) {
	notes := make([]note, 0)
	for _, v := range i.noteStore {
		notes = append(notes, v)
	}
	return notes, nil
}
