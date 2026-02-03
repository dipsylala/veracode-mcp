# Veracode Credentials Configuration

This package handles loading Veracode API credentials from multiple sources.

## Credential Sources (Priority Order)

1. **File-based configuration** (Preferred)
   - Location: `~/.veracode/veracode.yml`
   - Format:

     ```yaml
     api:
       key-id: 1*************1****
       key-secret: c*********************************************************a
       # override-api-base-url is optional - api is auto-detected from key-id prefix (vera01ei-* → EU, otherwise → US)
     ```

2. **Environment Variables** (Fallback)
   - `VERACODE_API_ID`
   - `VERACODE_API_KEY`
   - `VERACODE_OVERRIDE_API_BASE_URL` (Optional, auto-detected from API ID if not set)

## Usage

```go
import "github.com/dipsylala/veracodemcp-go/credentials"

// Get credentials from file or environment
apiID, apiSecret, baseURL, err := credentials.GetCredentials()
if err != nil {
    log.Fatal(err)
}
// baseURL defaults to "https://api.veracode.com" if not specified

// Get credentials with source information
apiID, apiSecret, baseURL, source, err := credentials.GetCredentialsWithFallback()
// source will be "file" or "env"
```

## Veracode API Regions

The package supports all Veracode regions with **automatic region detection** based on your API key ID.

### Automatic Region Detection

The package automatically selects the correct API region:

- **EU Region**: API key IDs beginning with `vera01ei-` → `https://api.veracode.eu`
- **US Region** (Default): All other API key IDs → `https://api.veracode.com`

This means you typically don't need to manually configure the region - just provide your credentials and the correct endpoint will be used automatically.

### Manual Region Override

You can override automatic detection by explicitly setting the base URL:

**Available Regions:**

- `https://api.veracode.com` - US Region (default)
- `https://api.veracode.eu` - European Region  
- `https://api.veracode.us` - US Federal Region

Configure the region in your `veracode.yml`:

```yaml
api:
  key-id: YOUR_API_KEY_ID
  key-secret: YOUR_API_KEY_SECRET
  override-api-base-url: https://api.veracode.eu  # Override auto-detection
```

Or via environment variable:

```bash
export VERACODE_OVERRIDE_API_BASE_URL="https://api.veracode.eu"
```

**Note**: Explicit configuration always takes precedence over automatic detection.

## Setting Up Credentials File

### Linux/macOS

```bash
mkdir -p ~/.veracode
cat > ~/.veracode/veracode.yml << EOF
api:
  key-id: YOUR_API_KEY_ID
  key-secret: YOUR_API_KEY_SECRET
EOF
chmod 600 ~/.veracode/veracode.yml
```

### Windows PowerShell

```powershell
New-Item -ItemType Directory -Path "$env:USERPROFILE\.veracode" -Force
@"
api:
  key-id: YOUR_API_KEY_ID
  key-secret: YOUR_API_KEY_SECRET
"@ | Out-File -FilePath "$env:USERPROFILE\.veracode\veracode.yml" -Encoding UTF8
```

## Getting Veracode API Credentials

1. Log in to the Veracode Platform
2. Go to **Settings → API Credentials**
3. Generate a new API key pair
4. Save the `key-id` and `key-secret` to your `veracode.yml` file

## Security Notes

- The credentials file should have restricted permissions (600 on Unix systems)
- Never commit `veracode.yml` to version control
- The API key secret is hex-encoded and used for HMAC signature generation
- Credentials are loaded once at client initialization and kept in memory
