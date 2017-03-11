package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestGetUsers test HTTP Get to "/users" using ResponseRecorder
func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

// TestGetUsersWithServer test HTTP Get to "/users" using Server
func TestGetUsersWithServer(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()
	usersURL := fmt.Sprintf("%s/users", server.URL)
	request, err := http.NewRequest("GET", usersURL, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}
