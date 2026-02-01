# CWE-93: CRLF Injection - Python

## LLM Guidance

CRLF Injection occurs when untrusted user input containing carriage return (`\r`) and line feed (`\n`) characters is used in HTTP headers, logs, or email headers without proper sanitization. Attackers exploit this to perform HTTP response splitting, header injection, log forgery, cache poisoning, and XSS attacks. Always strip or reject all newline characters (`\r`, `\n`, `\r\n`) and their URL-encoded equivalents (`%0d`, `%0a`) from user input before including in HTTP headers, logs, or protocol fields.

## Remediation Strategy

- Strip all CRLF characters (literal and encoded: `\r`, `\n`, `%0d`, `%0a`) from user input before use in headers or logs
- Validate input against strict allowlists using regex patterns (e.g., alphanumeric with specific symbols only)
- Use framework-provided sanitization methods (Flask's `make_response()`, Django's `HttpResponse()`) which provide built-in protection
- Implement structured logging (JSON format) to prevent log injection attacks
- Apply length limits to header values and logged content to prevent DoS attacks

## Remediation Steps

- Identify all locations where user input flows into HTTP headers, email headers, log messages, or response fields
- Implement `sanitize_header_value()` function that removes `\r`, `\n`, URL-encoded variants (`%0d`, `%0a`, `%0D`, `%0A`), and all control characters (`\x00-\x1f`, `\x7f`)
- Add input validation using strict allowlist regex patterns before sanitization (e.g., `^[a-zA-Z0-9._-]{3,50}$` for usernames)
- Apply length limits (200-254 chars for headers, appropriate limits for log entries) to prevent resource exhaustion
- Use validated/sanitized values in all header assignments, redirect calls, and logging statements
- Test with CRLF payloads - `value%0d%0aInjected-Header -%20malicious` and verify they're neutralized

## Minimal Safe Pattern

```python
from flask import Flask, request, redirect, abort
import re
from urllib.parse import urlparse

app = Flask(__name__)

def sanitize_url(url):
    if not url:
        return None
    # Remove CRLF (literal and encoded)
    clean_url = url.replace('\r', '').replace('\n', '')
    clean_url = clean_url.replace('%0d', '').replace('%0a', '')
    clean_url = clean_url.replace('%0D', '').replace('%0A', '')
    
    # Validate URL structure
    try:
        parsed = urlparse(clean_url)
        if parsed.scheme not in ['http', 'https', '']:
            return None
        return clean_url
    except:
        return None

@app.route('/redirect')
def secure_redirect():
    url = request.args.get('url', '')
    clean_url = sanitize_url(url)
    if not clean_url:
        abort(400, "Invalid URL")
    return redirect(clean_url)
```
