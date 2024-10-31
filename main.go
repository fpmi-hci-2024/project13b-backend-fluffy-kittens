package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler handles requests to the "/hello" endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", helloHandler) // Define a handler for "/hello"

	fmt.Println("Starting server on :8080")
	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
