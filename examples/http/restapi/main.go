package main

import (
	"log"
	"net/http"

	// external
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

//Entry point of the program
func main() {
	logger, _ := zap.NewProduction() // Create Uber's Zap logger
	h := &handler{
		repository: newInmemoryRepository(), // Injecting dependency
		logger:     logger,
	}
	r := initializeRoutes(h) // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}

func initializeRoutes(h *handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/notes", h.getAll).Methods("GET")
	r.HandleFunc("/api/notes/{id}", h.get).Methods("GET")
	r.HandleFunc("/api/notes", h.post).Methods("POST")
	r.HandleFunc("/api/notes/{id}", h.put).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", h.delete).Methods("DELETE")
	return r
}
