package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// loggingHandler is an HTTP Middleware that logs HTTP requests.
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before executing given Handler
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		// Middleware logic after executing given Handler
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	html :=
		`<doctype html>
        <html>
	<head>
		<title>Hello Gopher</title>
	</head>
	<body>
		<b>Hello Gopher!</b>
        <p>
            <a href="/welcome">Welcome</a> |  <a href="/message">Message</a>
        </p>
	</body>
</html>`
	fmt.Fprintf(w, html)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Programming")
}
func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "net/http package is used to build web apps")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", loggingHandler(http.HandlerFunc(index)))
	mux.Handle("/welcome", loggingHandler(http.HandlerFunc(welcome)))
	mux.Handle("/message", loggingHandler(http.HandlerFunc(message)))
	log.Println("Listening...")
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
