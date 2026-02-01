# CWE-330: Use of Insufficiently Random Values - PHP

## LLM Guidance

Weak random number generation in PHP occurs when insecure functions like `rand()`, `mt_rand()`, or `uniqid()` are used for security-sensitive operations such as generating session tokens, password reset tokens, API keys, or CSRF tokens. These functions use predictable pseudo-random number generators that attackers can exploit. Always use cryptographically secure functions `random_bytes()` and `random_int()` (PHP 7.0+) for any security-related randomness.

## Key Principles

- Replace all instances of `rand()`, `mt_rand()`, `uniqid()`, and `lcg_value()` with `random_int()` or `random_bytes()` in security contexts
- Use `random_bytes()` for generating tokens, keys, and binary random data
- Use `random_int()` for random integers within a specific range
- Never use predictable PRNGs for authentication, authorization, cryptography, or session management
- Verify token length is sufficient (at least 16-32 bytes for security tokens)

## Remediation Steps

- Identify all uses of weak random functions in security-sensitive code paths
- Replace token generation with `bin2hex(random_bytes($length))` or `base64_encode(random_bytes($length))`
- Replace random number generation with `random_int($min, $max)`
- Ensure minimum token length of 32 characters (16 bytes) for session/CSRF tokens
- Test that tokens are unpredictable and unique across multiple generations
- Review and update any cryptographic operations to use secure randomness

## Safe Pattern

```php
// Generate secure session token
function generateSecureToken($length = 32) {
    return bin2hex(random_bytes($length / 2));
}

// Generate secure random integer
function generateSecurePin() {
    return random_int(100000, 999999);
}

// Usage
$sessionToken = generateSecureToken(32);  // 32-char hex token
$resetToken = bin2hex(random_bytes(32));  // 64-char hex token
$csrfToken = base64_encode(random_bytes(24)); // 32-char base64 token
```
