package utils

import (
	"net/http"
)

// EnableCORS is a middleware function to handle CORS headers
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for requests from https://project13b-web-fluffy-kittens.onrender.com
		w.Header().Set("Access-Control-Allow-Origin", "https://project13b-web-fluffy-kittens.onrender.com")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
