# CWE-331: Insufficient Entropy - Python

## LLM Guidance

Insufficient entropy occurs when using Python's `random` module instead of `secrets` or `os.urandom()` for security-sensitive operations like generating tokens, encryption keys, IVs, or nonces. The `random` module uses pseudorandom number generation designed for statistical purposes, producing predictable values that attackers can exploit. Always use the `secrets` module (Python 3.6+) for cryptographic operations.

## Remediation Strategy

- Use `secrets` module exclusively for all cryptographic random value generation
- Never use `random` module for tokens, passwords, keys, IVs, nonces, or security decisions
- Ensure sufficient length for generated values (minimum 32 bytes for tokens)
- Use `os.urandom()` as fallback for Python versions before 3.6
- Apply cryptographically secure randomness at all security boundaries

## Remediation Steps

- Replace `random.Random()`, `random.randint()`, `random.choice()` with `secrets` equivalents
- Use `secrets.token_bytes(32)` for binary tokens and keys
- Use `secrets.token_hex(32)` for hexadecimal tokens
- Use `secrets.token_urlsafe(32)` for URL-safe tokens
- Use `secrets.choice()` for selecting random elements securely
- Verify token length meets security requirements (≥32 bytes recommended)

## Minimal Safe Pattern

```python
import secrets

# Generate secure token (32 bytes = 256 bits)
token = secrets.token_urlsafe(32)

# Generate secure random bytes for keys
encryption_key = secrets.token_bytes(32)

# Secure random selection
allowed_chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
secure_code = ''.join(secrets.choice(allowed_chars) for _ in range(8))

# Compare tokens safely
if secrets.compare_digest(user_token, stored_token):
    # Grant access
```
