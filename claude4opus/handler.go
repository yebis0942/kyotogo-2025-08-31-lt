package claude4opus

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// NewHandler returns an http.Handler with two endpoints:
// - / returns "Hello, world!"
// - /health returns {"status": "ok"}
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	// Root endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only handle exact path "/"
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"status": "ok",
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	return mux
}

func main() {
	const port = 8080

	// Create the handler
	handler := NewHandler()

	// Create the server with the handler
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	// Log that the server is starting
	log.Printf("Starting server on port %d...", port)

	// Start the server and handle any errors
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
