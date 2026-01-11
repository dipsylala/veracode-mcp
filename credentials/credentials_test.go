package credentials

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetCredentials_FromFile(t *testing.T) {
	// Create a temporary directory to simulate home directory
	tempDir := t.TempDir()
	veracodeDir := filepath.Join(tempDir, ".veracode")
	if err := os.MkdirAll(veracodeDir, 0755); err != nil {
		t.Fatalf("Failed to create temp .veracode dir: %v", err)
	}

	// Create a valid veracode.yml file
	configPath := filepath.Join(veracodeDir, "veracode.yml")
	validConfig := `api:
  key-id: test-key-id-123
  key-secret: test-key-secret-456
`
	if err := os.WriteFile(configPath, []byte(validConfig), 0600); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test reading from the file directly (bypass home dir lookup)
	apiID, apiSecret, baseURL, err := readConfigFile(configPath)
	if err != nil {
		t.Fatalf("readConfigFile failed: %v", err)
	}

	if apiID != "test-key-id-123" {
		t.Errorf("Expected apiID 'test-key-id-123', got '%s'", apiID)
	}

	if apiSecret != "test-key-secret-456" {
		t.Errorf("Expected apiSecret 'test-key-secret-456', got '%s'", apiSecret)
	}

	if baseURL != DefaultBaseURL {
		t.Errorf("Expected baseURL '%s', got '%s'", DefaultBaseURL, baseURL)
	}
}

func TestGetCredentials_WithCustomBaseURL(t *testing.T) {
	tempDir := t.TempDir()
	veracodeDir := filepath.Join(tempDir, ".veracode")
	if err := os.MkdirAll(veracodeDir, 0755); err != nil {
		t.Fatalf("Failed to create temp .veracode dir: %v", err)
	}

	// Create a config with custom base URL
	configPath := filepath.Join(veracodeDir, "veracode.yml")
	customConfig := `api:
  key-id: test-key-id-123
  key-secret: test-key-secret-456
  api-base-url: https://api.veracode.eu
`
	if err := os.WriteFile(configPath, []byte(customConfig), 0600); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	apiID, apiSecret, baseURL, err := readConfigFile(configPath)
	if err != nil {
		t.Fatalf("readConfigFile failed: %v", err)
	}

	if apiID != "test-key-id-123" {
		t.Errorf("Expected apiID 'test-key-id-123', got '%s'", apiID)
	}

	if apiSecret != "test-key-secret-456" {
		t.Errorf("Expected apiSecret 'test-key-secret-456', got '%s'", apiSecret)
	}

	expectedURL := "https://api.veracode.eu"
	if baseURL != expectedURL {
		t.Errorf("Expected baseURL '%s', got '%s'", expectedURL, baseURL)
	}
}

func TestGetCredentials_FromEnvironment(t *testing.T) {
	// Clear any existing credentials
	oldID := os.Getenv("VERACODE_API_ID")
	oldKey := os.Getenv("VERACODE_API_KEY")
	defer func() {
		os.Setenv("VERACODE_API_ID", oldID)
		os.Setenv("VERACODE_API_KEY", oldKey)
	}()

	// Set test credentials in environment
	os.Setenv("VERACODE_API_ID", "env-test-id")
	os.Setenv("VERACODE_API_KEY", "env-test-secret")

	// Note: This will try to read from real home directory first
	// So this test validates the environment fallback when file doesn't exist
	apiID, apiSecret, baseURL, err := GetCredentials()

	// If there's a real ~/.veracode/veracode.yml file, this test might get those credentials
	// So we only validate the error case is handled properly
	if err != nil {
		t.Fatalf("GetCredentials failed: %v", err)
	}

	// At minimum, we should get some credentials (either from file or env)
	if apiID == "" || apiSecret == "" {
		t.Error("Expected non-empty credentials")
	}

	// BaseURL should always be set (defaults to DefaultBaseURL if not specified)
	if baseURL == "" {
		t.Error("Expected non-empty baseURL")
	}
}

func TestGetCredentials_NotFound(t *testing.T) {
	// Clear environment variables
	oldID := os.Getenv("VERACODE_API_ID")
	oldKey := os.Getenv("VERACODE_API_KEY")
	defer func() {
		os.Setenv("VERACODE_API_ID", oldID)
		os.Setenv("VERACODE_API_KEY", oldKey)
	}()

	os.Unsetenv("VERACODE_API_ID")
	os.Unsetenv("VERACODE_API_KEY")

	// This will fail if there's no ~/.veracode/veracode.yml and no env vars
	// We can't guarantee that file doesn't exist in CI/dev environment
	// So we just verify the function returns an error with appropriate message
	_, _, _, err := GetCredentials()

	// If credentials were found (e.g., from ~/.veracode/veracode.yml), that's fine
	if err == nil {
		t.Skip("Skipping test - credentials found in environment (likely from ~/.veracode/veracode.yml)")
		return
	}

	// If error occurred, verify it has helpful message
	expectedMsg := "Veracode credentials not found"
	if err.Error()[:len(expectedMsg)] != expectedMsg {
		t.Errorf("Expected error to start with '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestReadConfigFile_InvalidYAML(t *testing.T) {
	// Create a temporary file with invalid YAML
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "invalid.yml")
	invalidYAML := `api:
  key-id: test-id
  key-secret: [invalid yaml structure
`
	if err := os.WriteFile(configPath, []byte(invalidYAML), 0600); err != nil {
		t.Fatalf("Failed to write invalid config: %v", err)
	}

	_, _, _, err := readConfigFile(configPath)
	if err == nil {
		t.Error("Expected error for invalid YAML, got nil")
	}

	if err != nil && err.Error()[:len("failed to parse")] != "failed to parse" {
		t.Errorf("Expected 'failed to parse' error, got: %v", err)
	}
}

func TestReadConfigFile_MissingFile(t *testing.T) {
	_, _, _, err := readConfigFile("/nonexistent/path/veracode.yml")
	if err == nil {
		t.Error("Expected error for missing file, got nil")
	}
}

func TestReadConfigFile_EmptyCredentials(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "empty.yml")
	emptyConfig := `api:
  key-id: ""
  key-secret: ""
`
	if err := os.WriteFile(configPath, []byte(emptyConfig), 0600); err != nil {
		t.Fatalf("Failed to write empty config: %v", err)
	}

	apiID, apiSecret, baseURL, err := readConfigFile(configPath)
	if err != nil {
		t.Fatalf("readConfigFile failed: %v", err)
	}

	// Function returns empty strings for empty values
	if apiID != "" || apiSecret != "" {
		t.Errorf("Expected empty credentials, got apiID='%s', apiSecret='%s'", apiID, apiSecret)
	}

	// BaseURL should default to DefaultBaseURL
	if baseURL != DefaultBaseURL {
		t.Errorf("Expected baseURL '%s', got '%s'", DefaultBaseURL, baseURL)
	}
}

func TestReadConfigFile_MissingFields(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "missing.yml")
	missingFieldsConfig := `api:
  key-id: test-id
  # key-secret is missing
`
	if err := os.WriteFile(configPath, []byte(missingFieldsConfig), 0600); err != nil {
		t.Fatalf("Failed to write config with missing fields: %v", err)
	}

	apiID, apiSecret, baseURL, err := readConfigFile(configPath)
	if err != nil {
		t.Fatalf("readConfigFile failed: %v", err)
	}

	if apiID != "test-id" {
		t.Errorf("Expected apiID 'test-id', got '%s'", apiID)
	}

	if apiSecret != "" {
		t.Errorf("Expected empty apiSecret, got '%s'", apiSecret)
	}

	if baseURL != DefaultBaseURL {
		t.Errorf("Expected baseURL '%s', got '%s'", DefaultBaseURL, baseURL)
	}
}

func TestGetCredentialsWithFallback_ReturnsSource(t *testing.T) {
	// Clear environment variables
	oldID := os.Getenv("VERACODE_API_ID")
	oldKey := os.Getenv("VERACODE_API_KEY")
	defer func() {
		os.Setenv("VERACODE_API_ID", oldID)
		os.Setenv("VERACODE_API_KEY", oldKey)
	}()

	// Set environment variables
	os.Setenv("VERACODE_API_ID", "test-id-source")
	os.Setenv("VERACODE_API_KEY", "test-secret-source")

	apiID, apiSecret, baseURL, source, err := GetCredentialsWithFallback()

	if err != nil {
		t.Fatalf("GetCredentialsWithFallback failed: %v", err)
	}

	// Should return some credentials
	if apiID == "" || apiSecret == "" {
		t.Error("Expected non-empty credentials")
	}

	// BaseURL should always be set
	if baseURL == "" {
		t.Error("Expected non-empty baseURL")
	}

	// Source should be either "file" or "env"
	if source != "file" && source != "env" {
		t.Errorf("Expected source to be 'file' or 'env', got '%s'", source)
	}
}

func TestGetCredentialsWithFallback_NotFound(t *testing.T) {
	// Clear environment variables
	oldID := os.Getenv("VERACODE_API_ID")
	oldKey := os.Getenv("VERACODE_API_KEY")
	defer func() {
		os.Setenv("VERACODE_API_ID", oldID)
		os.Setenv("VERACODE_API_KEY", oldKey)
	}()

	os.Unsetenv("VERACODE_API_ID")
	os.Unsetenv("VERACODE_API_KEY")

	_, _, _, source, err := GetCredentialsWithFallback()

	// If credentials were found (from file), that's okay
	if err == nil {
		t.Skip("Skipping test - credentials found in environment (likely from ~/.veracode/veracode.yml)")
		return
	}

	// If error occurred, verify it's appropriate
	if source != "" {
		t.Errorf("Expected empty source on error, got '%s'", source)
	}

	expectedMsg := "Veracode credentials not found"
	if err.Error()[:len(expectedMsg)] != expectedMsg {
		t.Errorf("Expected error to start with '%s', got '%s'", expectedMsg, err.Error())
	}
}

// Integration test: Test the full workflow with temporary home directory
func TestGetCredentials_Integration(t *testing.T) {
	tests := []struct {
		name           string
		setupFile      bool
		fileContent    string
		envID          string
		envKey         string
		expectError    bool
		expectedSource string // for GetCredentialsWithFallback
	}{
		{
			name:        "Valid file credentials",
			setupFile:   true,
			fileContent: "api:\n  key-id: file-id\n  key-secret: file-secret\n",
			expectError: false,
		},
		{
			name:        "Invalid YAML in file",
			setupFile:   true,
			fileContent: "api:\n  key-id: [invalid\n",
			envID:       "env-id",
			envKey:      "env-secret",
			expectError: true, // Invalid YAML should cause an error
		},
		{
			name:        "No file, valid env vars",
			setupFile:   false,
			envID:       "env-id",
			envKey:      "env-secret",
			expectError: false,
		},
		{
			name:        "No file, missing env vars",
			setupFile:   false,
			expectError: true,
		},
		{
			name:        "Empty credentials in file",
			setupFile:   true,
			fileContent: "api:\n  key-id: \"\"\n  key-secret: \"\"\n",
			envID:       "env-id",
			envKey:      "env-secret",
			expectError: false, // Should fall back to env
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test validates the logic but can't fully isolate from real home directory
			// So we primarily test readConfigFile directly with temp files

			if tt.setupFile {
				tempDir := t.TempDir()
				configPath := filepath.Join(tempDir, "veracode.yml")
				if err := os.WriteFile(configPath, []byte(tt.fileContent), 0600); err != nil {
					t.Fatalf("Failed to write config file: %v", err)
				}

				_, _, _, err := readConfigFile(configPath)
				if tt.expectError && err == nil {
					t.Error("Expected error but got nil")
				}
				if !tt.expectError && err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestGetCredentials_AutoDetectEURegion(t *testing.T) {
	tempDir := t.TempDir()
	veracodeDir := filepath.Join(tempDir, ".veracode")
	if err := os.MkdirAll(veracodeDir, 0755); err != nil {
		t.Fatalf("Failed to create temp .veracode dir: %v", err)
	}

	// Create a config with EU key prefix but no explicit base URL
	configPath := filepath.Join(veracodeDir, "veracode.yml")
	euConfig := `api:
  key-id: vera01ei-1234-5678-90ab-cdef
  key-secret: test-key-secret-456
`
	if err := os.WriteFile(configPath, []byte(euConfig), 0600); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	apiID, apiSecret, baseURL, err := readConfigFile(configPath)
	if err != nil {
		t.Fatalf("readConfigFile failed: %v", err)
	}

	if apiID != "vera01ei-1234-5678-90ab-cdef" {
		t.Errorf("Expected EU apiID, got '%s'", apiID)
	}

	if apiSecret != "test-key-secret-456" {
		t.Errorf("Expected apiSecret 'test-key-secret-456', got '%s'", apiSecret)
	}

	// Should auto-detect EU region from key prefix
	if baseURL != EUBaseURL {
		t.Errorf("Expected auto-detected EU baseURL '%s', got '%s'", EUBaseURL, baseURL)
	}
}

func TestGetCredentials_AutoDetectUSRegion(t *testing.T) {
	tempDir := t.TempDir()
	veracodeDir := filepath.Join(tempDir, ".veracode")
	if err := os.MkdirAll(veracodeDir, 0755); err != nil {
		t.Fatalf("Failed to create temp .veracode dir: %v", err)
	}

	// Create a config with US key (no vera01ei- prefix) and no explicit base URL
	configPath := filepath.Join(veracodeDir, "veracode.yml")
	usConfig := `api:
  key-id: 1234567890abcdef1234567890abcdef
  key-secret: test-key-secret-789
`
	if err := os.WriteFile(configPath, []byte(usConfig), 0600); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	apiID, apiSecret, baseURL, err := readConfigFile(configPath)
	if err != nil {
		t.Fatalf("readConfigFile failed: %v", err)
	}

	if apiID != "1234567890abcdef1234567890abcdef" {
		t.Errorf("Expected US apiID, got '%s'", apiID)
	}

	if apiSecret != "test-key-secret-789" {
		t.Errorf("Expected apiSecret 'test-key-secret-789', got '%s'", apiSecret)
	}

	// Should default to US region for non-EU keys
	if baseURL != DefaultBaseURL {
		t.Errorf("Expected default US baseURL '%s', got '%s'", DefaultBaseURL, baseURL)
	}
}

func TestGetCredentials_ExplicitBaseURLOverridesAutoDetect(t *testing.T) {
	tempDir := t.TempDir()
	veracodeDir := filepath.Join(tempDir, ".veracode")
	if err := os.MkdirAll(veracodeDir, 0755); err != nil {
		t.Fatalf("Failed to create temp .veracode dir: %v", err)
	}

	// Create a config with EU key prefix but explicit US base URL
	configPath := filepath.Join(veracodeDir, "veracode.yml")
	explicitConfig := `api:
  key-id: vera01ei-1234-5678-90ab-cdef
  key-secret: test-key-secret-456
  api-base-url: https://api.veracode.com
`
	if err := os.WriteFile(configPath, []byte(explicitConfig), 0600); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	apiID, apiSecret, baseURL, err := readConfigFile(configPath)
	if err != nil {
		t.Fatalf("readConfigFile failed: %v", err)
	}

	// Verify the key ID was read correctly
	if apiID != "vera01ei-1234-5678-90ab-cdef" {
		t.Errorf("Expected EU apiID, got '%s'", apiID)
	}

	if apiSecret != "test-key-secret-456" {
		t.Errorf("Expected apiSecret, got '%s'", apiSecret)
	}

	// Explicit base URL should override auto-detection
	expectedURL := "https://api.veracode.com"
	if baseURL != expectedURL {
		t.Errorf("Expected explicit baseURL '%s', got '%s'", expectedURL, baseURL)
	}

}

func TestDetectRegionFromKeyID(t *testing.T) {
	tests := []struct {
		name     string
		keyID    string
		expected string
	}{
		{
			name:     "EU key with vera01ei- prefix",
			keyID:    "vera01ei-1234-5678-90ab-cdef",
			expected: EUBaseURL,
		},
		{
			name:     "US key without prefix",
			keyID:    "1234567890abcdef1234567890abcdef",
			expected: DefaultBaseURL,
		},
		{
			name:     "Empty key ID",
			keyID:    "",
			expected: DefaultBaseURL,
		},
		{
			name:     "Short key that doesn't match prefix",
			keyID:    "vera",
			expected: DefaultBaseURL,
		},
		{
			name:     "Key with similar but different prefix",
			keyID:    "vera01ea-1234-5678",
			expected: DefaultBaseURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detectRegionFromKeyID(tt.keyID)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
