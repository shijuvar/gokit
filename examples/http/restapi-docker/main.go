package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



type handler struct {
	repository repository
}
//Store for the Notes collection

//Variable to generate key for the collection
var id int = 0

//HTTP Post - /api/notes
func (h handler) post(w http.ResponseWriter, r *http.Request) {
	var note note
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	if err:= h.repository.create(note); err!=nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//HTTP Get - /api/notes
func (h handler) getAll(w http.ResponseWriter, r *http.Request) {
	// Get all
	if notes, err := h.repository.getAll(); err !=nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(notes)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
//HTTP Get - /api/notes/{id}
func (h handler) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Get by id
	if note, err := h.repository.getById(id); err !=nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(note)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//HTTP Put - /api/notes/{id}
func (h handler) put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var note note
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatal(err)
	}
	// Update
	if err:= h.repository.update(id,note ); err!=nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//HTTP Delete - /api/notes/{id}
func (h handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// delete
	if err:= h.repository.delete(id); err!=nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//Entry point of the program
func main() {
	r := mux.NewRouter()
	h:= handler{
		repository: newInmemoryRepository(), // Injecting dependency
	}
	r.HandleFunc("/api/notes", h.getAll).Methods("GET")
	r.HandleFunc("/api/notes/{id}", h.get).Methods("GET")
	r.HandleFunc("/api/notes", h.post).Methods("POST")
	r.HandleFunc("/api/notes/{id}", h.put).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", h.delete).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
