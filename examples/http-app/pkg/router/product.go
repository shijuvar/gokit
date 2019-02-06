package router

import (
	"github.com/gorilla/mux"

	"github.com/shijuvar/gokit/examples/http-app/pkg/auth"
	"github.com/shijuvar/gokit/examples/http-app/pkg/controller"
	"github.com/shijuvar/gokit/examples/http-app/pkg/postgres"
)

// SetUserRoutes registers routes for user product
func SetProductRoutes(router *mux.Router, store postgres.DataStore) *mux.Router {

	productStore := postgres.ProductStore{Store: store}
	productController := controller.ProductController{Store: productStore}
	productRouter := mux.NewRouter()

	productRouter.Handle("/products", controller.ResponseHandler(productController.PostProduct)).Methods("POST")
	productRouter.Handle("/products", controller.ResponseHandler(productController.GetAllProducts)).Methods("GET")
	productRouter.Handle("/products/{id}", controller.ResponseHandler(productController.GetProductById)).Methods("GET")
	// Applying authorization middleware
	router.PathPrefix("/products").Handler(auth.AuthorizeRequest(productRouter))
	return router
}
