package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	tests := map[string]struct {
		path            string
		wantStatusCode  int
		wantBody        string
		wantContentType string
	}{
		"GET /": {
			path:            "/",
			wantStatusCode:  http.StatusOK,
			wantBody:        "Hello, world!",
			wantContentType: "text/plain; charset=utf-8",
		},
		"GET /health": {
			path:            "/health",
			wantStatusCode:  http.StatusOK,
			wantBody:        `{"status": "ok"}`,
			wantContentType: "application/json",
		},
		"GET /notfound": {
			path:            "/notfound",
			wantStatusCode:  http.StatusNotFound,
			wantBody:        "404 page not found\n",
			wantContentType: "text/plain; charset=utf-8",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mux := NewServer()
			req, err := http.NewRequest("GET", tc.path, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if rr.Code != tc.wantStatusCode {
				t.Errorf("got status %d, want %d", rr.Code, tc.wantStatusCode)
			}
			if rr.Body.String() != tc.wantBody {
				t.Errorf("got body %q, want %q", rr.Body.String(), tc.wantBody)
			}
			if contentType := rr.Header().Get("Content-Type"); contentType != tc.wantContentType {
				t.Errorf("got Content-Type %q, want %q", contentType, tc.wantContentType)
			}
		})
	}
}
