package middleware

import (
	"net/http"
	"time"
)

func Cache(startDate time.Time) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			modifiedSince := r.Header.Get("If-Modified-Since")
			if modifiedSince == "" {
				next.ServeHTTP(w, r)
				return
			}

			sinceDate, err := time.Parse(http.TimeFormat, modifiedSince)
			if err == nil && sinceDate.Before(startDate) {
				w.WriteHeader(http.StatusNotModified)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
