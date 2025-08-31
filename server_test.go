package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/yebis0942/kyotogo-2025-08-31-lt/gemini25flash"
	"github.com/yebis0942/kyotogo-2025-08-31-lt/gemini25pro"
	"github.com/yebis0942/kyotogo-2025-08-31-lt/human"
	"github.com/yebis0942/kyotogo-2025-08-31-lt/kimik2"
	"github.com/yebis0942/kyotogo-2025-08-31-lt/qwen25coder32binstruct"
)

func TestServer(t *testing.T) {
	handlers := map[string]func() http.Handler{
		"Human":                     human.NewHander,
		"Gemini 2.5 Flash":          gemini25flash.NewHandler,
		"Gemini 2.5 Pro":            gemini25pro.NewHandler,
		"Qwen25 Coder 32B Instruct": qwen25coder32binstruct.NewHandler,
		"Kimi-K2":                   kimik2.NewHandler,
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
