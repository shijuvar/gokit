package routers

import (
	"github.com/gorilla/mux"

	"github.com/shijuvar/gokit/examples/bookmark-api/controllers"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
