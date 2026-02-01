# CWE-798: Hard-coded Credentials - Python

## LLM Guidance

Hard-coded credentials (passwords, API keys, database credentials, encryption keys) in Python source code create critical security vulnerabilities by exposing secrets in version control and making rotation impossible. Always externalize credentials using environment variables, configuration files outside version control, or dedicated secrets managers. For production systems, use cloud-native solutions like AWS Secrets Manager or Azure Key Vault.

## Key Principles

- Never commit credentials to version control; use `.gitignore` for local config files
- Separate configuration from code using environment variables or external config files
- Use secrets managers for production deployments with automatic rotation
- Apply principle of least privilege to all credentials
- Implement secure defaults and fail securely when credentials are missing

## Remediation Steps

- Identify all hard-coded credentials in source code using grep/scanning tools
- Replace hard-coded values with `os.getenv()` calls with no defaults for secrets
- Store credentials in environment variables or `.env` files (add `.env` to `.gitignore`)
- For production, migrate to secrets managers (AWS Secrets Manager, Azure Key Vault, HashiCorp Vault)
- Rotate all exposed credentials immediately
- Implement validation to ensure required credentials are present at startup

## Safe Pattern

```python
import os
from dotenv import load_dotenv

load_dotenv()  # Load from .env file for local dev

DB_PASSWORD = os.getenv("DB_PASSWORD")
API_KEY = os.getenv("API_KEY")

if not DB_PASSWORD or not API_KEY:
    raise ValueError("Missing required credentials")

# Use credentials
connection = connect(password=DB_PASSWORD)
```
