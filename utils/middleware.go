package utils

import (
	"net/http"
)

func JsonRoute(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Invalid content type", http.StatusBadRequest)
		}

		h.ServeHTTP(w, r)
	})
}
