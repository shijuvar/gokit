package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shijuvar/gokit/examples/testing/httpbdd/controllers"
	"github.com/shijuvar/gokit/examples/testing/httpbdd/store"
)

func setUserRoutes() *mux.Router {
	r := mux.NewRouter()
	// Resolve dependencies
	userStore := &store.MongoUserStore{}
	controller := controllers.Handler{
		Store: userStore, // Injecting dependencies
	}
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":8080", setUserRoutes())
}
