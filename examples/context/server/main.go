package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", loggingHandler(http.HandlerFunc(index)))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Print("Index Handler started")
	defer log.Print("Index Handler ended")
	log.Print(r.Context().Value("user"))
	ctx := r.Context()
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello Gopher")
	case <-ctx.Done():
		err := ctx.Err()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// loggingHandler is an HTTP Middleware that logs HTTP requests.
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before executing given Handler
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		ctx := context.WithValue(r.Context(), "user", "shijuvar")
		// Calls the next handler by providing the Context
		next.ServeHTTP(w, r.WithContext(ctx))
		// Middleware logic after executing given Handler
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
