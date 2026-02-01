# CWE-331: Insufficient Entropy - PHP

## LLM Guidance

Insufficient entropy in PHP occurs when using `rand()`, `mt_rand()`, `uniqid()`, or other predictable functions for security-sensitive operations like generating tokens, encryption keys, IVs, or nonces. These functions produce predictable values that attackers can exploit. Always use `random_bytes()` or `random_int()` (PHP 7.0+) for cryptographic randomness.

## Remediation Strategy

- Replace all `rand()`, `mt_rand()`, `uniqid()`, and `srand()` calls with cryptographically secure alternatives in security contexts
- Use `random_bytes()` for generating random binary data (tokens, keys, IVs)
- Use `random_int()` for random integers within a specific range
- Never use predictable functions for password resets, session tokens, CSRF tokens, or encryption keys
- Ensure sufficient entropy length (minimum 16 bytes for tokens, 32 bytes for keys)

## Remediation Steps

- Identify all uses of `rand()`, `mt_rand()`, `uniqid()`, and similar functions in security-sensitive code
- Replace with `random_bytes()` for binary data or `random_int()` for integers
- Convert binary output to appropriate format using `bin2hex()` or `base64_encode()` if needed
- Ensure minimum entropy requirements (128+ bits for tokens, 256+ bits for encryption keys)
- Test token uniqueness and verify unpredictability in security reviews

## Minimal Safe Pattern

```php
// Generate secure random token (32 bytes = 256 bits)
$token = bin2hex(random_bytes(32));

// Generate secure random integer
$randomNumber = random_int(1000, 9999);

// Generate session token
$sessionToken = base64_encode(random_bytes(24));
```
