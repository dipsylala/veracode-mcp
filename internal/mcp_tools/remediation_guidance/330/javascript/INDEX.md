# CWE-330: Use of Insufficiently Random Values - JavaScript/Node.js

## LLM Guidance

`Math.random()` is predictable and unsuitable for security-sensitive operations like generating session tokens, CSRF tokens, API keys, or cryptographic material. For security purposes, always use `crypto.randomBytes()` (Node.js) or `crypto.getRandomValues()` (browser), which provide cryptographically secure random values.

## Key Principles

- Replace all `Math.random()` calls in security contexts with cryptographic random generators
- Use minimum 32 bytes of entropy for tokens and keys
- Implement token expiration to limit exposure window
- Store tokens securely with proper validation mechanisms
- Use URL-safe encoding (base64url) for tokens in URLs

## Remediation Steps

- Identify all `Math.random()` usage for tokens, keys, or secrets
- Replace with `crypto.randomBytes(32)` for Node.js or `crypto.getRandomValues()` for browsers
- Convert random bytes to appropriate format (base64url for tokens, hex for keys)
- Add expiration timestamps to all tokens (typically 1-24 hours)
- Store tokens with hashing in database for validation
- Test token uniqueness and unpredictability

## Safe Pattern

```javascript
const crypto = require('crypto');

function generateSecureToken(userId) {
    // Generate 32 bytes of cryptographically secure random data
    const token = crypto.randomBytes(32).toString('base64url');
    const expiry = Date.now() + (3600 * 1000); // 1 hour
    
    // Store hashed token with expiry
    const hashedToken = crypto.createHash('sha256').update(token).digest('hex');
    storeToken(userId, hashedToken, expiry);
    
    return token;
}
```
