package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Apply(h http.Handler, mw ...Middleware) http.Handler {
	for _, middleware := range mw {
		h = middleware(h)
	}
	return h
}