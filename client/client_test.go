package client

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoReturnsStructuredAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"Invalid parameter","code":400}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	_, err := client.do(context.Background(), http.MethodPost, "/api/test", nil, nil)
	if err == nil {
		t.Fatal("expected an error")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T", err)
	}

	if apiErr.Response.Error != "Invalid parameter" {
		t.Fatalf("expected parsed error message, got %q", apiErr.Response.Error)
	}

	if apiErr.Response.Code != 400 {
		t.Fatalf("expected parsed error code 400, got %d", apiErr.Response.Code)
	}

	if apiErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, apiErr.StatusCode)
	}
}
