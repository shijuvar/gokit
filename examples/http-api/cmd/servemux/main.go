package main

import (
	"log"
	"net/http"

	apphttp "github.com/shijuvar/gokit/examples/http-api/httpmux"
	"github.com/shijuvar/gokit/examples/http-api/memstore"
)

// Entry point of the program
func main() {
	repo, err := memstore.NewInmemoryRepository() // With in-memory database
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &apphttp.NoteHandler{
		Repository: repo, // Injecting dependency
	}
	router := initializeRoutes(h) // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}

func initializeRoutes(h *apphttp.NoteHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/notes", h.GetAll)
	mux.HandleFunc("GET /api/notes/{id}", h.Get)
	mux.HandleFunc("POST /api/notes", h.Post)
	mux.HandleFunc("PUT /api/notes/{id}", h.Put)
	mux.HandleFunc("DELETE /api/notes/{id}", h.Delete)
	return mux
}
