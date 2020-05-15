package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

func PanicRecovery(logger *log.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				err := recover()
				if err != nil {
					logger.Print("Error", err)
					jsonError, _ := json.Marshal(map[string]string{
						"error": "There was an internal server error",
					})

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(jsonError)
				}

			}()

			next.ServeHTTP(w, r)

		})
	}
}
