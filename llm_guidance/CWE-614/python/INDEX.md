# CWE-614: Sensitive Cookie Without 'Secure' Flag - Python

## LLM Guidance

Sensitive Cookie Without 'Secure' Flag occurs when cookies containing authentication tokens, session IDs, or other sensitive data are transmitted without the `Secure` flag, allowing them to be sent over unencrypted HTTP connections. This exposes cookies to interception through man-in-the-middle attacks and network sniffing. The fix requires setting `secure=True` on all sensitive cookies and enforcing HTTPS.

## Remediation Strategy

- Always set `secure=True` for cookies containing sensitive data (sessions, auth tokens, user identifiers)
- Enforce HTTPS site-wide and redirect HTTP requests to HTTPS
- Set `httponly=True` to prevent JavaScript access to sensitive cookies
- Use `samesite='Strict'` or `'Lax'` to prevent CSRF attacks
- Configure framework session settings to use secure defaults

## Remediation Steps

- Identify all locations where cookies are set in your application
- Add `secure=True` parameter to all sensitive cookie assignments
- Configure framework session management to enable secure cookies by default
- Enable HTTPS across your entire application infrastructure
- Set additional security flags - `httponly=True` and `samesite='Strict'/'Lax'`
- Test that cookies are not transmitted over HTTP connections

## Minimal Safe Pattern

```python
from flask import Flask, session

app = Flask(__name__)
app.config.update(
    SESSION_COOKIE_SECURE=True,
    SESSION_COOKIE_HTTPONLY=True,
    SESSION_COOKIE_SAMESITE='Strict'
)

response.set_cookie(
    'session_id',
    value=token,
    secure=True,
    httponly=True,
    samesite='Strict'
)
```
