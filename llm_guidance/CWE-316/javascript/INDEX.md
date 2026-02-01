# CWE-316: Cleartext Storage of Sensitive Information in Memory - JavaScript/Node.js

## LLM Guidance

Storing sensitive data (passwords, cryptographic keys, tokens) in memory as cleartext in JavaScript exposes it to memory dumps, debugging tools, and memory disclosure vulnerabilities. JavaScript strings are immutable in V8, making them persist in memory. Use `Buffer` for sensitive data, clear buffers explicitly with `fill(0)`, and avoid logging or concatenating sensitive values.

## Remediation Strategy

- Use mutable Buffers: Store credentials in `Buffer` objects instead of strings
- Explicit zeroing: Clear sensitive buffers with `.fill(0)` immediately after use
- Minimal lifetime: Keep sensitive data in memory only as long as necessary
- Avoid string operations: Never convert sensitive buffers to strings or log them
- Secure comparison: Use `crypto.timingSafeEqual()` for comparing sensitive values

## Remediation Steps

- Replace string variables holding passwords/keys with `Buffer.from()` or `Buffer.alloc()`
- Wrap sensitive operations in `try-finally` blocks with `.fill(0)` in finally
- Pass buffers directly to crypto functions without converting to strings
- Remove debug logging, console output, and error messages containing sensitive data
- Clear buffers before returning from functions handling credentials
- Use libraries like `libsodium-wrappers` that provide automatic memory protection

## Minimal Safe Pattern

```javascript
const crypto = require('crypto');

function authenticateUser(passwordBuffer) {
  const hashedPassword = Buffer.alloc(32);
  try {
    crypto.pbkdf2Sync(passwordBuffer, salt, 100000, 32, 'sha256').copy(hashedPassword);
    const isValid = crypto.timingSafeEqual(hashedPassword, storedHash);
    return isValid;
  } finally {
    hashedPassword.fill(0);
    passwordBuffer.fill(0); // Clear input
  }
}
```
