package claude4sonnet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HealthResponse represents the JSON response for the health endpoint
type HealthResponse struct {
	Status string `json:"status"`
}

// NewHandler returns an http.Handler with the required endpoints
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	// Root endpoint - returns "Hello, world!"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})

	// Health endpoint - returns JSON {"status": "ok"}
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := HealthResponse{Status: "ok"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding health response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	return mux
}

func main() {
	const port = 8080

	handler := NewHandler()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	log.Printf("Server starting on port %d", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
