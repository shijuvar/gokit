package postgres

import "github.com/shijuvar/gokit/examples/http-app/pkg/domain"

// ProductStore provides persistence logic for "products" table
type ProductStore struct {
	Store DataStore
}

// ToDO: Write CRUD operations here
// Create creates a new Product
func (productStore ProductStore) Create(product domain.Product) (domain.Product, error) {
	// ToDo: Write the code here
	return domain.Product{}, nil
}
