# CWE-601: Open Redirect - Python

## LLM Guidance

Open redirect vulnerabilities occur when user-controlled input determines redirect destinations without validation, enabling phishing and credential theft. The core fix is validating that redirect URLs are either relative paths (starting with `/` but not `//`) or match an allowlist of trusted domains using `urlparse()`. For Flask, Django, and FastAPI, use framework-specific validators and never directly pass user input to redirect functions.

## Remediation Strategy

- Validate redirect URLs are relative paths using `urlparse().netloc == ''` to ensure no external domain
- Maintain an allowlist of trusted domains for absolute URLs and reject all others
- Sanitize user input by checking that paths start with `/` but not `//` to prevent protocol-relative URLs
- Use framework-specific safe redirect utilities (Django's `url_has_allowed_host_and_scheme()`, Flask's `is_safe_url()`)

## Remediation Steps

- Parse the redirect URL using `urllib.parse.urlparse()` to extract scheme, netloc, and path components
- Check if `netloc` is empty (relative URL) or matches your allowlist of trusted domains
- Reject URLs with schemes other than `http`/`https` or protocol-relative URLs (`//example.com`)
- For relative paths, ensure they start with `/` but not `//`
- Use framework validators - Django's `url_has_allowed_host_and_scheme()` or implement custom validation
- Set a default safe redirect (e.g., `/dashboard`) when validation fails

## Minimal Safe Pattern

```python
from urllib.parse import urlparse

def safe_redirect(url, allowed_hosts=['example.com']):
    parsed = urlparse(url)
    
    # Allow relative URLs without domain
    if not parsed.netloc:
        return url if url.startswith('/') and not url.startswith('//') else '/'
    
    # Allow absolute URLs only from trusted hosts
    if parsed.netloc in allowed_hosts and parsed.scheme in ['http', 'https']:
        return url
    
    return '/'  # Default safe fallback
```
