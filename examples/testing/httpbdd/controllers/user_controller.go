package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shijuvar/gokit/examples/testing/httpbdd/model"
)

type Handler struct {
	// Dependencies and States
	Store model.UserStore
}

// GetUsers serves requests for Http Get to "/users"
func (handler Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	data := handler.Store.GetUsers()
	users, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

// CreateUser serves requests for Http Post to "/users"
func (handler Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("[Controllers.CreateUser]: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Insert User entity into User Store
	err = handler.Store.AddUser(user)
	if err != nil {
		if err == model.ErrorEmailExists {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}
