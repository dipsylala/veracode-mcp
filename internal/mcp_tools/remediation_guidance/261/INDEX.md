# CWE-261 - Weak Encoding for Password

## LLM Guidance

Weak encoding schemes (Base64, XOR, ROT13, URL encoding) are not cryptography and provide zero security. Encoded passwords are trivially reversible and offer no protection. Use strong cryptographic password hashing algorithms (bcrypt, Argon2, PBKDF2) instead of encoding.

## Key Principles

- Replace all encoding with proper cryptographic password hashing
- Use adaptive hashing algorithms with appropriate work factors
- Store only password hashes, never plaintext or encoded passwords
- Implement secure password comparison using constant-time functions
- Use password salting to prevent rainbow table attacks

## Remediation Steps

- Identify weak encoding patterns - Search for `btoa()`, `atob()`, `base64_encode()`, `base64.b64encode()` in password handling code
- Check for XOR operations - Find XOR operations applied to password strings or static key obfuscation
- Review database schemas - Look for VARCHAR/TEXT password columns indicating reversible storage
- Locate reversible operations - Find URL encoding, hex encoding, or other reversible transformations on passwords
- Replace with strong hashing - Implement bcrypt, Argon2, or scrypt with appropriate cost parameters
- Update password comparison logic - Use hash verification functions instead of decoded string comparisons
