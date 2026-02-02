# CWE-295: Improper Certificate Validation - Python

## LLM Guidance

Improper certificate validation occurs when Python applications disable SSL/TLS certificate verification (e.g., `requests.get(url, verify=False)` or `urllib3` with `cert_reqs='CERT_NONE'`), enabling man-in-the-middle attacks where attackers intercept and modify encrypted traffic. The core fix is to enable certificate validation by using default settings or explicitly setting `verify=True`, allowing Python to validate certificates against trusted certificate authorities.

## Key Principles

- Always enable certificate validation - Never disable SSL/TLS verification in production code
- Use system certificate stores - Rely on Python's default CA bundle or OS-provided certificates
- Handle certificate errors properly - Log and investigate certificate failures rather than disabling validation
- Keep dependencies updated - Ensure `requests`, `urllib3`, and `certifi` libraries are current
- Use certificate pinning for critical connections - Pin specific certificates for high-security requirements

## Remediation Steps

- Remove all instances of `verify=False` from `requests` calls
- Remove `cert_reqs='CERT_NONE'` from `urllib3` configurations
- Use `verify=True` explicitly or rely on secure defaults
- Specify custom CA bundles with `verify='/path/to/ca-bundle.crt'` only when necessary
- Update certificate-related packages - `pip install --upgrade requests urllib3 certifi`
- Test connections to ensure certificate validation works in all environments

## Safe Pattern

```python
import requests

# Secure: Certificate validation enabled (default)
response = requests.get('https://api.example.com/data')

# Secure: Explicit verification
response = requests.get('https://api.example.com/data', verify=True)

# Secure: Custom CA bundle when needed
response = requests.get(
    'https://internal-api.company.com/data',
    verify='/etc/ssl/certs/company-ca-bundle.crt'
)
```
