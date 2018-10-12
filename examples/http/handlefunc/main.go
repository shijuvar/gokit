package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	fmt.Fprintf(w, "Welcome to Go Web Programming\n")
}
func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "net/http package is used to build web apps\n")
}

func main() {
	// NewServeMux returns a new ServeMux.
	mux := http.NewServeMux()
	// HandleFunc registers the handler function for the given pattern.
	mux.HandleFunc("/", index)
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/message", message)
	log.Println("Listening...")
	// http.ListenAndServe(":8080", mux)
	// A Server defines parameters for running an HTTP server.
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux, // handler to invoke, http.DefaultServeMux if nil
	}
	// ListenAndServe listens on the TCP network address and
	// then calls Serve to handle requests on incoming connections.
	server.ListenAndServe()
}
