package kimik2

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Create the handler using the function provided below.
	handler := NewHandler()

	const addr = ":8080"
	log.Printf("Server starting on http://localhost%s ...", addr)

	// http.ListenAndServe always returns a non-nil error unless
	// http.ErrServerClosed, which is returned after Shutdown.
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// NewHandler returns an http.Handler that serves:
//   - GET /         -> "Hello, world!"
//   - GET /health   -> {"status":"ok"}
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only allow the exact path "/" so "/foo" doesn't match.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello, world!"))
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	return mux
}
