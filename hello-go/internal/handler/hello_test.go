package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-fuego/fuego"

	"github.com/hello-bazel/hello-go/internal/handler"
)

func TestHelloWorld(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "default greeting",
			query:    "",
			expected: "Hello, World!",
		},
		{
			name:     "custom greeting",
			query:    "?to=Bazel",
			expected: "Hello, Bazel!",
		},
		{
			name:     "empty to parameter",
			query:    "?to=",
			expected: "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := fuego.NewServer()
			fuego.Get(s, "/", handler.HelloWorld)

			req := httptest.NewRequest(http.MethodGet, "/"+tt.query, nil)
			rec := httptest.NewRecorder()

			s.Mux.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
			}

			// Response body contains the expected greeting
			body := rec.Body.String()
			if body != tt.expected {
				t.Errorf("expected body %q, got %q", tt.expected, body)
			}
		})
	}
}
