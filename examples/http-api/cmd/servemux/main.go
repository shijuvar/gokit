package main

import (
	"log"
	"log/slog"
	"net/http"

	apphttp "github.com/shijuvar/gokit/examples/http-api/httpmux"
	"github.com/shijuvar/gokit/examples/http-api/memstore"
	"github.com/shijuvar/gokit/examples/http-api/middleware"
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

	logger := slog.Default()
	// Adding middleware handlers
	router = middleware.Apply(router,
		middleware.RateLimiter(200),
		middleware.PanicRecovery(logger),
	)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}

func initializeRoutes(h *apphttp.NoteHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/notes", h.GetAll)
	mux.HandleFunc("GET /api/notes/{id}", h.Get)
	mux.HandleFunc("POST /api/notes", h.Post)
	mux.HandleFunc("PUT /api/notes/{id}", h.Put)
	mux.HandleFunc("DELETE /api/notes/{id}", h.Delete)
	return mux
}
