package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shijuvar/gokit/examples/testing/httpbdd/model"
)

// GetUsers serves requests for Http Get to "/users"
func GetUsers(store model.UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := store.GetUsers()
		users, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(users)
	})

}

// CreateUser serves requests for Http Post to "/users"
func CreateUser(store model.UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		// Decode the incoming User json
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Fatalf("[Controllers.CreateUser]: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Insert User entity into User Store
		err = store.AddUser(user)
		if err != nil {
			if err == model.ErrorEmailExists {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
