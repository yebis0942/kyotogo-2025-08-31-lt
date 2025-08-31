package gemini25flash

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// NewHandler returns an http.Handler with two endpoints.
// The root "/" endpoint returns "Hello, World!".
// The "/health" endpoint returns a JSON object with {"status": "ok"}.
func NewHandler() http.Handler {
	// Create a new ServeMux to handle different endpoints.
	mux := http.NewServeMux()

	// Handle the root path "/"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	// Handle the "/health" path
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to application/json.
		w.Header().Set("Content-Type", "application/json")

		// Create a map for the JSON response.
		response := map[string]string{"status": "ok"}

		// Encode the map into JSON and write it to the response writer.
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	return mux
}
