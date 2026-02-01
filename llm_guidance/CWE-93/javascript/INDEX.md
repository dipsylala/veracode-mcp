# CWE-93: CRLF Injection - JavaScript

## LLM Guidance

CRLF Injection occurs when untrusted input containing carriage return (`\r`) and line feed (`\n`) characters is used in HTTP headers or protocol fields without validation, enabling HTTP response splitting, header injection, cache poisoning, and XSS attacks. The primary defense is to strip or reject newline characters from all user input before using it in headers or protocol-sensitive contexts.

## Remediation Strategy

- Remove or reject `\r` and `\n` characters from user input before use in HTTP headers
- Use framework-provided header-setting methods that automatically sanitize input
- Validate input against allowlists rather than attempting to sanitize dangerous characters
- Encode or escape special characters when headers must include dynamic content
- Never concatenate raw user input directly into HTTP response headers

## Remediation Steps

- Identify all locations where user input flows into HTTP headers or response fields
- Replace manual header construction with framework methods (e.g., Express `res.set()`, `res.setHeader()`)
- Add input validation to strip `\r`, `\n`, and URL-encoded variants (`%0D`, `%0A`)
- Use allowlist validation for headers with predictable values (e.g., redirect URLs)
- Test with payloads containing CRLF sequences to verify sanitization effectiveness
- Review logging and error handling to ensure headers aren't logged with unsanitized input

## Minimal Safe Pattern

```javascript
// Safe header setting using Express framework methods
app.get('/redirect', (req, res) => {
  const url = req.query.url;
  
  // Validate and sanitize input
  const sanitized = url.replace(/[\r\n]/g, '');
  
  // Use framework method (automatically prevents CRLF)
  res.setHeader('Location', sanitized);
  res.status(302).send();
});
```
