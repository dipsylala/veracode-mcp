package api

import (
	"context"
	"testing"
	"time"
)

func TestLookupMCPVerademo(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping lookup test in short mode")
	}

	client, err := NewVeracodeClient()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	app, err := client.GetApplicationByName(ctx, "MCPVerademo")
	if err != nil {
		t.Fatalf("Failed to get app: %v", err)
	}

	if app == nil {
		t.Fatal("App not found")
	}

	t.Logf("MCPVerademo App ID: %s", *app.Guid)
	t.Logf("App Name: %s", *app.Profile.Name)
}
