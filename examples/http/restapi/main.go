package main

import (
	"log"
	"net/http"

	// external
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	apphttp "github.com/shijuvar/gokit/examples/http/restapi/http"
	"github.com/shijuvar/gokit/examples/http/restapi/memstore"
)

//Entry point of the program
func main() {
	logger, _ := zap.NewProduction()              // Create Uber's Zap logger
	repo, err := memstore.NewInmemoryRepository() // With in-memory database
	//repo, err := newMongoNoteRepository() // With MongoDB database
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &apphttp.NoteHandler{
		Repository: repo, // Injecting dependency
		Logger:     logger,
	}
	router := initializeRoutes(h) // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}

func initializeRoutes(h *apphttp.NoteHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/notes", h.GetAll).Methods("GET")
	r.HandleFunc("/api/notes/{id}", h.Get).Methods("GET")
	r.HandleFunc("/api/notes", h.Post).Methods("POST")
	r.HandleFunc("/api/notes/{id}", h.Put).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", h.Delete).Methods("DELETE")
	return r
}
