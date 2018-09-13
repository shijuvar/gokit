package domain

import (
	"github.com/pkg/errors"
)

// ToDo: All domain services
func (p Product) Valid() (bool, error) {
	var err error
	if p.Name == "" {
		err = errors.New("Name could not be empty")
	}
	if p.SKU == "" {
		err = errors.New("SKU could not be empty")
	}
	// Do all validation logic here
	if err != nil {
		return false, err
	} else {
		return true, err
	}

}
