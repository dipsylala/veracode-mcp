# CWE-331: Insufficient Entropy - JavaScript/Node.js

## LLM Guidance

Insufficient entropy occurs when using `Math.random()` for security-sensitive operations instead of cryptographically secure sources. `Math.random()` produces predictable values unsuitable for tokens, keys, or secrets. Always use `crypto.randomBytes()` (Node.js) or `crypto.getRandomValues()` (browser) for security-critical randomness.

**Primary Defence:** Use `crypto.randomBytes()` (Node.js) or `crypto.getRandomValues()` (browser) for all security-sensitive random value generation.

## Remediation Strategy

- Never use `Math.random()`, `Date.now()`, or predictable sources for security-sensitive values
- Always use cryptographically secure random number generators (CSPRNGs) from the `crypto` module
- Generate sufficient entropy: minimum 128 bits (16 bytes) for tokens, 256 bits (32 bytes) for keys
- Use appropriate encoding (hex, base64, base64url) for the use case
- Review third-party libraries to ensure they use secure random sources

## Remediation Steps

- Identify all `Math.random()` usage in security contexts (tokens, session IDs, passwords, keys)
- Replace with `crypto.randomBytes()` in Node.js or `crypto.getRandomValues()` in browsers
- Ensure sufficient byte length - 16+ bytes for tokens, 32+ bytes for cryptographic keys
- Convert to appropriate format using `.toString('hex')`, `.toString('base64')`, or base64url encoding
- Test that generated values are unpredictable and non-repeating
- Audit dependencies for insecure random usage

## Minimal Safe Pattern

```javascript
const crypto = require('crypto');

// Generate secure random token (Node.js)
function generateSecureToken() {
  return crypto.randomBytes(32).toString('hex'); // 64-char hex string
}

// Generate secure session ID
function generateSessionId() {
  return crypto.randomBytes(16).toString('base64'); // 24-char base64 string
}

const token = generateSecureToken();
const sessionId = generateSessionId();
```
