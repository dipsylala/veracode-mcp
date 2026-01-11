package credentials

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// VeracodeConfig represents the structure of the veracode.yml file
type VeracodeConfig struct {
	API struct {
		KeyID     string `yaml:"key-id"`
		KeySecret string `yaml:"key-secret"`
		BaseURL   string `yaml:"api-base-url,omitempty"`
	} `yaml:"api"`
}

const (
	// DefaultBaseURL is the default Veracode API base URL (US region)
	DefaultBaseURL = "https://api.veracode.com"
	// EUBaseURL is the Veracode API base URL for European region
	EUBaseURL = "https://api.veracode.eu"
	// EUKeyPrefix is the prefix for European region API keys
	EUKeyPrefix = "vera01ei-"
)

// GetCredentials retrieves Veracode API credentials from ~/.veracode/veracode.yml
// Falls back to environment variables VERACODE_API_ID and VERACODE_API_KEY if file doesn't exist
// Returns apiID, apiSecret, baseURL, and error. baseURL will be DefaultBaseURL if not specified.
func GetCredentials() (apiID, apiSecret, baseURL string, err error) {
	// Try to read from ~/.veracode/veracode.yml first
	homeDir, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(homeDir, ".veracode", "veracode.yml")
		apiID, apiSecret, baseURL, err = readConfigFile(configPath)
		if err == nil && apiID != "" && apiSecret != "" {

			return apiID, apiSecret, baseURL, nil
		}
	}

	// Fall back to environment variables
	apiID = os.Getenv("VERACODE_API_ID")
	apiSecret = os.Getenv("VERACODE_API_KEY")
	baseURL = os.Getenv("VERACODE_API_BASE_URL")

	if baseURL == "" {
		baseURL = detectRegionFromKeyID(apiID)
	}

	if apiID == "" || apiSecret == "" {
		return "", "", "", fmt.Errorf("Veracode credentials not found. " +
			"Please create ~/.veracode/veracode.yml with key-id and key-secret" +
			"or set VERACODE_API_ID and VERACODE_API_KEY environment variables")
	}

	return apiID, apiSecret, baseURL, nil
}

// readConfigFile reads and parses the veracode.yml configuration file
func readConfigFile(path string) (string, string, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", "", err
	}

	var config VeracodeConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return "", "", "", fmt.Errorf("failed to parse %s: %w", path, err)
	}

	baseURL := config.API.BaseURL
	if baseURL == "" {
		// Auto-detect region from API key ID if no explicit base URL is configured
		baseURL = detectRegionFromKeyID(config.API.KeyID)
	}

	return config.API.KeyID, config.API.KeySecret, baseURL, nil
}

// detectRegionFromKeyID automatically detects the region based on API key ID prefix
// Returns the appropriate base URL for the region
func detectRegionFromKeyID(apiID string) string {
	if len(apiID) >= len(EUKeyPrefix) && apiID[:len(EUKeyPrefix)] == EUKeyPrefix {
		return EUBaseURL
	}
	return DefaultBaseURL
}

// GetCredentialsWithFallback retrieves credentials with custom fallback logic
// Returns apiID, apiSecret, baseURL, source ("file" or "env"), and error
func GetCredentialsWithFallback() (apiID, apiSecret, baseURL, source string, err error) {
	// Try to read from ~/.veracode/veracode.yml first
	homeDir, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(homeDir, ".veracode", "veracode.yml")
		apiID, apiSecret, baseURL, err = readConfigFile(configPath)
		if err == nil && apiID != "" && apiSecret != "" {

			return apiID, apiSecret, baseURL, "file", nil
		}
	}

	// Fall back to environment variables
	apiID = os.Getenv("VERACODE_API_ID")
	apiSecret = os.Getenv("VERACODE_API_KEY")
	baseURL = os.Getenv("VERACODE_API_BASE_URL")

	if baseURL == "" {
		baseURL = detectRegionFromKeyID(apiID)
	}

	if apiID == "" || apiSecret == "" {
		return "", "", "", "", fmt.Errorf("Veracode credentials not found. " +
			"Please create ~/.veracode/veracode.yml with key-id and key-secret " +
			"or set VERACODE_API_ID and VERACODE_API_KEY environment variables")
	}

	return apiID, apiSecret, baseURL, "env", nil
}
