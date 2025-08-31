package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatalf("Environment variable PORT is not set")
	}

	mux := NewServer()
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	log.Printf("Server listening on :%s", port)
}

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", helloHandler)
	mux.HandleFunc("GET /health", healthHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := io.WriteString(w, "Hello, world!"); err != nil {
		log.Printf("Error writing response for '/': %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if _, err := io.WriteString(w, `{"status": "ok"}`); err != nil {
		log.Printf("Error writing response for '/health': %v", err)
	}
}
