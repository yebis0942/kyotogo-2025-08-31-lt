package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	handlers := map[string]func() http.Handler{
		"NewServer":                 NewServer,
		"NewServer_gemini_25_flash": NewServer_gemini_25_flash,
		"NewServer_gemini_25_pro":   NewServer_gemini_25_pro,
	}

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

	for handlerName, handlerFunc := range handlers {
		t.Run(handlerName, func(t *testing.T) {
			for testName, tc := range tests {
				t.Run(testName, func(t *testing.T) {
					mux := handlerFunc()
					req, err := http.NewRequest("GET", tc.path, nil)
					if err != nil {
						t.Fatalf("Failed to create request: %v", err)
					}
					rr := httptest.NewRecorder()
					mux.ServeHTTP(rr, req)

					if rr.Code != tc.wantStatusCode {
						t.Errorf("got status %d, want %d", rr.Code, tc.wantStatusCode)
					}
					if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(tc.wantBody) {
						t.Errorf("got body %q, want %q", rr.Body.String(), tc.wantBody)
					}
					if contentType := rr.Header().Get("Content-Type"); contentType != tc.wantContentType {
						t.Errorf("got Content-Type %q, want %q", contentType, tc.wantContentType)
					}
				})
			}
		})
	}
}
