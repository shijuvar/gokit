package main

import (
	"fmt"
	"log"
	"net/http"
)

type textHandler struct {
	responseText string
}

func (th *textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, th.responseText)
}

type indexHandler struct {
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &indexHandler{})

	thWelcome := &textHandler{"Welcome to Go Web Programming"}
	mux.Handle("/welcome", thWelcome)

	thMessage := &textHandler{"net/http package is used to build web apps"}
	mux.Handle("/message", thMessage)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
