package mapstore

import (
	"errors"
	"github.com/shijuvar/gokit/training/userdefinedtype/domain"
)

const customerIdSeq int = 0

// MapStore is an implementation of CustomerStore interface
type MapStore struct {
	// An in-memory store with a map
	// Use Customer.ID as the key of map
	store map[string]domain.Customer
}

// Factory method gives a new instance of MapStore
// This is for caller packages to create MapStore instances
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (s *MapStore) Create(c domain.Customer) error {
	if c.ID != "" {
		s.store[c.ID] = domain.Customer{ID: c.ID,
			Name: c.Name, Email: c.Email}
		return nil
	} else {
		return errors.New("Cannot create ID with a nil value, please enter a valid value")
	}
}

func (s *MapStore) Update(k string, c domain.Customer) error {
	if k != "" {
		s.store[k] = c
		return nil
	} else {
		return errors.New("Cannot update Customer with a nil ID value")
	}
}

func (s *MapStore) Delete(k string) error {
	if k != "" {
		delete(s.store, k)
		return nil
	} else {
		return errors.New("Cannot delete Customer with a nil ID value")
	}
	return nil
}

func (s *MapStore) GetById(k string) (domain.Customer, error) {
	if k == "" {
		return domain.Customer{ID: "nil",
			Name:  "nil",
			Email: "nil"}, errors.New("Cannot delete Customer with a nil ID value")
	} else if v, found := s.store[k]; found {
		return v, nil
	} else {
		return domain.Customer{ID: "nil",
			Name:  "nil",
			Email: "nil"}, errors.New("Unexpected outcome")
	}
}

func (s *MapStore) GetAll() ([]domain.Customer, error) {
	values := make([]domain.Customer, len(s.store))
	if s != nil {
		i := 0
		for _, v := range s.store {
			values[i] = v
			i++
		}
		return values, nil
	} else {
		return values, errors.New("No data found in the MapStore")
	}

}
