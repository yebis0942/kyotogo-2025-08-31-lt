package main

import (
	"encoding/json"
	"net/http"
)

// healthResponse represents the JSON response for the /health endpoint.
type healthResponse struct {
	Status string `json:"status"`
}

// helloHandler handles requests to the root endpoint (/).
func qwen_helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, world!"))
}

// healthHandler handles requests to the /health endpoint.
func qwen_healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}
	response := healthResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// newMux sets up a new ServeMux with the defined handlers.
func NewServer_qwen25_Coder_32B_instruct() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", qwen_helloHandler)
	mux.HandleFunc("/health", qwen_healthHandler)
	return mux
}
