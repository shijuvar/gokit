package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shijuvar/gokit/examples/testing/httpbdd/controllers"
	"github.com/shijuvar/gokit/examples/testing/httpbdd/store"
)

func setUserRoutes() *mux.Router {
	r := mux.NewRouter()
	userStore := &store.MongoUserStore{}
	r.Handle("/users", controllers.CreateUser(userStore)).Methods("POST")
	r.Handle("/users", controllers.GetUsers(userStore)).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":8080", setUserRoutes())
}
