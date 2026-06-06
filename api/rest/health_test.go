package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	testAPIID  = "test-api-id"
	testAPIKey = "0000000000000000000000000000000000000000000000000000000000000000"
)

func TestCheckHealth_UsesAuthenticatedPrincipalEndpoint(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != principalEndpoint {
			t.Errorf("expected path %s, got %s", principalEndpoint, r.URL.Path)
		}
		if auth := r.Header.Get("Authorization"); !strings.HasPrefix(auth, "VERACODE-HMAC-SHA-256 ") {
			t.Errorf("expected Veracode HMAC Authorization header, got %q", auth)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"user_name":"test-user"}`))
	}))
	defer server.Close()

	client := &Client{
		apiID:   testAPIID,
		apiKey:  testAPIKey,
		baseURL: server.URL + "/",
	}

	status, err := client.CheckHealth(context.Background())
	if err != nil {
		t.Fatalf("CheckHealth returned an error: %v", err)
	}
	if status == nil {
		t.Fatal("expected health status, got nil")
	}
	if !status.Available {
		t.Fatalf("expected API to be available, got message %q", status.Message)
	}
	if status.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status.StatusCode)
	}
	if !strings.Contains(status.Message, "authentication succeeded") {
		t.Errorf("expected authentication success message, got %q", status.Message)
	}
}

func TestCheckHealth_ReportsPrincipalEndpointFailure(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}))
	defer server.Close()

	client := &Client{
		apiID:   testAPIID,
		apiKey:  testAPIKey,
		baseURL: server.URL,
	}

	status, err := client.CheckHealth(context.Background())
	if err != nil {
		t.Fatalf("CheckHealth returned an error: %v", err)
	}
	if status == nil {
		t.Fatal("expected health status, got nil")
	}
	if status.Available {
		t.Fatal("expected API to be unavailable")
	}
	if status.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, status.StatusCode)
	}
	if !strings.Contains(status.Message, "status 401") {
		t.Errorf("expected HTTP status in message, got %q", status.Message)
	}
}

func TestCheckHealth_WithCancelledContext(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &Client{
		apiID:   testAPIID,
		apiKey:  testAPIKey,
		baseURL: server.URL,
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	status, err := client.CheckHealth(ctx)
	if err != nil {
		t.Fatalf("CheckHealth returned an error: %v", err)
	}
	if status == nil {
		t.Fatal("expected health status, got nil")
	}
	if status.Available {
		t.Fatal("expected API to be unavailable")
	}
	if status.StatusCode != 0 {
		t.Errorf("expected status code 0, got %d", status.StatusCode)
	}
	if !strings.Contains(status.Message, "context canceled") {
		t.Errorf("expected cancellation message, got %q", status.Message)
	}
}

func TestCheckHealthSimple(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		{name: "available", statusCode: http.StatusOK, want: true},
		{name: "unavailable", statusCode: http.StatusServiceUnavailable, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			client := &Client{
				apiID:   testAPIID,
				apiKey:  testAPIKey,
				baseURL: server.URL,
			}

			if got := client.CheckHealthSimple(context.Background()); got != tt.want {
				t.Errorf("CheckHealthSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}
