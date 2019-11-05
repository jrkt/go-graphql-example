package graphql

import (
	"log"
	"net/http"
)

func LogMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("logging middleware method=[%s] url=[%s] referer=[%s]\n", r.Method, r.URL, r.Referer())
		h.ServeHTTP(w, r)
	}
}
