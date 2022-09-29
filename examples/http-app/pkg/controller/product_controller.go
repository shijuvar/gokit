package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shijuvar/gokit/examples/http-app/pkg/domain"
)

type ProductController struct {
	Store domain.ProductStore // This gives extensibility and testability
}

// HTTP Post - /products
func (handler ProductController) PostProduct(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var product domain.Product
	// Decode the incoming json data to note struct
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("Unable to decode JSON request body: %w", err)
	}
	// Persistence
	newProduct, err := handler.Store.Create(product)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("rrror on inserting Product: %w", err)
	}
	return newProduct, http.StatusCreated, nil
}

// HTTP Get - /products
func (handler ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	// ToDo: Write the implementation
	return nil, http.StatusOK, nil

}

// HTTP Get - /products/{id}
func (handler ProductController) GetProductById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	// Get id from the incoming url
	//vars := mux.Vars(r)
	//id := vars["id"]

	// ToDo: Write the implementation
	return nil, http.StatusOK, nil

}
