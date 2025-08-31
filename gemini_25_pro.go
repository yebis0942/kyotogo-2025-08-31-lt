package main

import (
	"fmt"
	"net/http"
)

// createRouter creates and returns a new http.Handler with the defined endpoints.
// This approach separates the routing logic from the server setup.
func NewServer_gemini_25_pro() http.Handler {
	// http.NewServeMux() creates a new request multiplexer (router).
	mux := http.NewServeMux()

	// Handle the root endpoint "/".
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ServeMux matches "/" for any path not otherwise handled.
		// This check ensures we only respond to the exact "/" path.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		// Fprint writes the string to the ResponseWriter.
		fmt.Fprint(w, "Hello, world!")
	})

	// Handle the "/health" endpoint for health checks.
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to indicate a JSON response.
		w.Header().Set("Content-Type", "application/json")
		// Write the HTTP status code 200 OK.
		w.WriteHeader(http.StatusOK)
		// Write the JSON body. For simple, static JSON, a raw string is efficient.
		fmt.Fprint(w, `{"status": "ok"}`)
	})

	return mux
}
