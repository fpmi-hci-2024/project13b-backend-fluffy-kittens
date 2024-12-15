package utils

import (
	"net/http"
)

// EnableCORS is a middleware function to handle CORS headers
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for requests from https://project13b-web-fluffy-kittens.onrender.com
		origin := r.Header.Get("Origin")

		// Allow requests from specific origins
		if origin == "https://project13b-web-fluffy-kittens.onrender.com" || origin == "http://localhost:4200" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
