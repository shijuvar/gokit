package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type (
	Response struct {
		Data interface{} `json:"data"`
	}
	handler struct {
		store map[string]note
	}
	note struct {
		Id          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedOn   time.Time `json:"createdon"`
	}
)

// Generic handler for writing response header and body
func responseHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)
		if err != nil {
			data = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if data != nil {
			err = json.NewEncoder(w).Encode(Response{Data: data})
			if err != nil {
				log.Printf("could not encode response to output: %v", err)
			}
		}

	})
}

//Variable to generate key for the collection
var id int = 0

//HTTP Post - /api/notes
func (h handler) postNoteHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var note note
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode JSON request body: %v", err)

	}

	note.CreatedOn = time.Now()
	id++
	note.Id = id
	k := strconv.Itoa(id)
	h.store[k] = note
	return note, http.StatusCreated, nil
}

//HTTP Get - /api/notes
func (h handler) getNoteHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var notes []note
	for _, v := range h.store {
		notes = append(notes, v)
	}
	return notes, http.StatusOK, nil
}

//HTTP Put - /api/notes/{id}
func (h handler) putNoteHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpd note
	// Decode the incoming Note json
	err = json.NewDecoder(r.Body).Decode(&noteToUpd)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode JSON request body: %v", err)
	}
	if note, ok := h.store[k]; ok {
		noteToUpd.CreatedOn = note.CreatedOn
		//delete existing item and add the updated item
		delete(h.store, k)
		h.store[k] = noteToUpd
	} else {
		return nil, http.StatusBadRequest, errors.New("could not find out a Note with given key")
	}
	return nil, http.StatusNoContent, nil
}

//HTTP Delete - /api/notes/{id}
func (h handler) deleteNoteHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	k := vars["id"]
	// Remove from Store
	if _, ok := h.store[k]; ok {
		//delete existing item
		delete(h.store, k)
	} else {
		return nil, http.StatusBadRequest, errors.New("could not find out a Note with given key")
	}
	return nil, http.StatusNoContent, nil
}

//Entry point of the program
func main() {

	h := handler{
		store: make(map[string]note),
	}
	r := mux.NewRouter()
	r.Handle("/api/notes", responseHandler(h.getNoteHandler)).Methods("GET")
	r.Handle("/api/notes", responseHandler(h.postNoteHandler)).Methods("POST")
	r.Handle("/api/notes/{id}", responseHandler(h.putNoteHandler)).Methods("PUT")
	r.Handle("/api/notes/{id}", responseHandler(h.deleteNoteHandler)).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
