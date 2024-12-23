package utils

import (
	"net/http"
)

// EnableCORS is a middleware function to handle CORS headers
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for requests from specific origins
		origin := r.Header.Get("Origin")
		if origin == "https://project13b-web-fluffy-kittens.onrender.com" || origin == "http://localhost:4200" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Handle preflight OPTIONS requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent) // Use 204 No Content for OPTIONS
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
