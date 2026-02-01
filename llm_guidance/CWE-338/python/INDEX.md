# CWE-338: Use of Cryptographically Weak Pseudo-Random Number Generator (PRNG) - Python

## LLM Guidance

Use of Cryptographically Weak PRNG occurs when developers use Python's `random` module (Mersenne Twister algorithm) for security-sensitive operations like generating tokens, passwords, or encryption keys. This module is predictable and allows attackers to forecast generated values. **Primary Defence:** Use `secrets` module (Python 3.6+) for all security-sensitive randomness.

## Remediation Strategy

- Replace `random` module with `secrets` module for tokens, passwords, keys, and security identifiers
- Use `secrets.token_bytes()`, `secrets.token_hex()`, or `secrets.token_urlsafe()` for cryptographic randomness
- Reserve `random` module only for non-security contexts like simulations, games, or testing
- Use `os.urandom()` as fallback for Python < 3.6 or lower-level randomness needs

## Remediation Steps

- Identify all `import random` statements and audit usage for security contexts
- Replace `random.randint()`, `random.choice()`, `random.random()` with `secrets` equivalents
- For session tokens - use `secrets.token_urlsafe(32)` (generates 32-byte URL-safe string)
- For numeric secrets - use `secrets.randbelow(n)` instead of `random.randint(0, n-1)`
- For password generation - use `secrets.choice()` with character sets
- Verify changes don't affect non-security code that legitimately uses `random`

## Minimal Safe Pattern

```python
import secrets
import string

# Generate secure session token
session_token = secrets.token_urlsafe(32)

# Generate secure API key
api_key = secrets.token_hex(32)

# Generate secure password
alphabet = string.ascii_letters + string.digits + string.punctuation
password = ''.join(secrets.choice(alphabet) for _ in range(16))

# Secure random number (0 to 99)
secure_number = secrets.randbelow(100)
```
