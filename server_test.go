package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yebis0942/kyotogo-2025-08-31-lt/gemini25flash"
	"github.com/yebis0942/kyotogo-2025-08-31-lt/gemini25pro"
	"github.com/yebis0942/kyotogo-2025-08-31-lt/human"
	"github.com/yebis0942/kyotogo-2025-08-31-lt/kimik2"
	"github.com/yebis0942/kyotogo-2025-08-31-lt/qwen25coder32binstruct"
)

func TestServer(t *testing.T) {
	handlers := map[string]func() http.Handler{
		"Human":                      human.NewHander,
		"Gemini 2.5 Flash":           gemini25flash.NewHandler,
		"Gemini 2.5 Pro":             gemini25pro.NewHandler,
		"Qwen2.5 Coder 32B Instruct": qwen25coder32binstruct.NewHandler,
		"Kimi-K2":                    kimik2.NewHandler,
	}

	tests := map[string]struct {
		path            string
		wantStatusCode  int
		assertBody      func(*testing.T, string)
		wantContentType string
	}{
		"GET /": {
			path:           "/",
			wantStatusCode: http.StatusOK,
			assertBody: func(t *testing.T, body string) {
				if body != "Hello, world!" {
					t.Errorf("got body %q, want %q", body, "Hello, world!")
				}
			},
			wantContentType: "text/plain; charset=utf-8",
		},
		"GET /health": {
			path:           "/health",
			wantStatusCode: http.StatusOK,
			assertBody: func(t *testing.T, body string) {
				var data map[string]interface{}

				if err := json.Unmarshal([]byte(body), &data); err != nil {
					t.Errorf("failed to unmarshal JSON: %v", err)
					return
				}

				if len(data) != 1 {
					t.Errorf("expected JSON object with a single key, got %d keys", len(data))
					return
				}

				status, ok := data["status"].(string)
				if !ok || status != "ok" {
					t.Errorf("expected status to be 'ok', got %v", data["status"])
					return
				}
			},
			wantContentType: "application/json",
		},
		"GET /notfound": {
			path:           "/notfound",
			wantStatusCode: http.StatusNotFound,
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
					if tc.assertBody != nil {
						tc.assertBody(t, rr.Body.String())
					}
					if tc.wantContentType != "" {
						if contentType := rr.Header().Get("Content-Type"); contentType != tc.wantContentType {
							t.Errorf("got Content-Type %q, want %q", contentType, tc.wantContentType)
						}
					}
				})
			}
		})
	}
}
