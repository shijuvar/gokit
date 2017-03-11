package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// User model
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// getUsers serves requests for Http Get to "/users"
func getUsers(w http.ResponseWriter, r *http.Request) {
	data := []User{
		User{
			FirstName: "Shiju",
			LastName:  "Varghese",
			Email:     "shiju@xyz.com",
		},

		User{
			FirstName: "Irene",
			LastName:  "Rose",
			Email:     "irene@xyz.com",
		},
	}
	users, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	http.ListenAndServe(":8080", r)
}
