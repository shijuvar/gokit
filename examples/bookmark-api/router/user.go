package router

import (
	"github.com/gorilla/mux"

	"github.com/shijuvar/gokit/examples/bookmark-api/controller"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controller.Register).Methods("POST")
	router.HandleFunc("/users/login", controller.Login).Methods("POST")
	return router
}
