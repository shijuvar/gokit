package controller

import (
	"encoding/json"
	"net/http"

	util "github.com/shijuvar/gokit/examples/http-app/pkg/apputil"
)

// response used to send HTTP responses
type response struct {
	Data interface{} `json:"data"`
}

// Generic handler for writing response header and body for all handler functions
func ResponseHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r) // execute application handler
		if err != nil {
			data = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if data != nil {
			// Send JSON response back to the client application
			err = json.NewEncoder(w).Encode(response{Data: data})
			if err != nil {
				util.Error.Printf("Error from Handler: %s\n", err.Error())
			}
		}

	})
}
