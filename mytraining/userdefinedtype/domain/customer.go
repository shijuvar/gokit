package domain

type Customer struct {
	ID    string
	Name  string
	Email string
}

type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}
