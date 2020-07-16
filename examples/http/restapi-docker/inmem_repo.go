package main

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type inmemoryRepository struct  {
	noteStore map[string]note
}

func newInmemoryRepository() repository {
	return &inmemoryRepository {
		noteStore : make(map[string]note),
	}
}
func (i *inmemoryRepository) create(n note) error {
	if _, ok := i.noteStore[n.Id]; ok {
		return errors.New("Id exists")
	}
	n.CreatedOn = time.Now()
	// Create a Version 4 UUID.
	uid, _ := uuid.NewV4()
	n.Id=uid.String()
	i.noteStore[n.Id] =n
	return nil
 }

func (i *inmemoryRepository) update(id string, n note) error {
	if _, ok := i.noteStore[id]; !ok {
		return errors.New("Id doesn't exist")
	}
	i.noteStore[id] =n
	return nil
}

func (i *inmemoryRepository) delete(id string) error {
	if _, ok := i.noteStore[id]; !ok {
		return errors.New("Id doesn't exist")
	}
	delete(i.noteStore,id)
	return nil
}
func (i *inmemoryRepository) getById(id string) (note,error) {
	if v, ok := i.noteStore[id]; !ok {
		return note{}, errors.New("Id doesn't exist")
	} else {
		return v,nil
	}

}

func (i *inmemoryRepository) getAll() ([]note, error) {
	notes:= make([]note,0)
	for _, v :=range i.noteStore {
		notes = append(notes,v)
	}
	return notes, nil
}
