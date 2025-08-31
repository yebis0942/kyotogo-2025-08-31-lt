package gpt5

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// NewHandler returns an http.Handler with two routes: "/" and "/health".
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	// Root endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	// Health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]string{"status": "ok"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	})

	return mux
}

func main() {
	const addr = ":8080"

	// Create server
	server := &http.Server{
		Addr:    addr,
		Handler: NewHandler(),
	}

	log.Printf("Starting server on %s", addr)

	// Start server with error handling
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start server: %v", err)
	}
}
