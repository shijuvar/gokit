package router

import (
	"github.com/gorilla/mux"

	"github.com/shijuvar/gokit/examples/http-app/pkg/controller"
	"github.com/shijuvar/gokit/examples/http-app/pkg/postgres"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router, store postgres.DataStore) *mux.Router {

	userStore := postgres.UserStore{Store: store}
	userController := controller.UserController{Store: userStore}
	router.Handle("/users", controller.ResponseHandler(userController.PostUser)).Methods("POST")
	router.Handle("/users/login", controller.ResponseHandler(userController.PostLogin)).Methods("POST")
	return router
}
