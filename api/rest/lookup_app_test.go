package rest

import (
	"context"
	"testing"
	"time"
)

func TestLookupMCPVerademo(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping lookup test in short mode")
	}

	client, err := NewClient()
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

	if app.Guid != nil {
		t.Logf("App GUID:    %s", *app.Guid)
	} else {
		t.Error("App GUID is nil")
	}

	if app.Profile != nil && app.Profile.Name != nil {
		t.Logf("App Name:    %s", *app.Profile.Name)
	}

	if app.Id != nil {
		t.Logf("App ID (numeric): %d", *app.Id)
	} else {
		t.Error("App numeric Id is nil — this field is required for --app-id")
	}

	if app.Oid != nil {
		t.Logf("App OID:     %d", *app.Oid)
	}
}
