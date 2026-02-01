# CWE-338: Use of Cryptographically Weak PRNG - PHP

## LLM Guidance

Use of Cryptographically Weak PRNG in PHP occurs when developers use non-cryptographic functions like `rand()`, `mt_rand()`, or `uniqid()` for security-sensitive operations such as generating tokens, keys, or passwords. These functions produce predictable values that attackers can exploit to compromise sessions, guess tokens, or break encryption. **Primary Defence:** Use `random_bytes()` or `random_int()` for all security-critical random value generation.

## Remediation Strategy

- Replace all `rand()`, `mt_rand()`, `uniqid()`, and `srand()` calls in security contexts with cryptographically secure alternatives
- Use `random_bytes()` for generating random binary data (tokens, keys, salts)
- Use `random_int()` for random integers within a specific range
- Never seed or predict CSPRNG output
- Ensure sufficient entropy (minimum 16 bytes for tokens, 32+ bytes for keys)

## Remediation Steps

- Identify all uses of weak PRNGs in authentication, session management, cryptography, and token generation
- Replace `rand()`/`mt_rand()` with `random_int()` for integer values
- Replace `uniqid()` with `bin2hex(random_bytes())` for unique identifiers
- Ensure proper encoding when converting binary random data (use `bin2hex()` or `base64_encode()`)
- Review and test all changes to verify randomness quality
- Remove any custom seeding logic for PRNGs

## Minimal Safe Pattern

```php
// Generate cryptographically secure token
function generateSecureToken($length = 32) {
    return bin2hex(random_bytes($length));
}

// Generate random integer in range
function generateSecurePin() {
    return random_int(1000, 9999);
}

// Usage
$sessionToken = generateSecureToken(32); // 64-char hex string
$resetToken = bin2hex(random_bytes(16)); // 32-char hex string
$otp = random_int(100000, 999999); // 6-digit number
```
