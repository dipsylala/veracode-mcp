# CWE-326: Inadequate Encryption Strength - Python

## LLM Guidance

Inadequate Encryption Strength occurs when Python applications use weak cryptographic algorithms (DES, 3DES, RC4), insufficient key sizes (AES-128, RSA-1024), or deprecated ciphers that modern computing can break. The core fix is replacing weak algorithms with strong alternatives: AES-256-GCM for symmetric encryption, RSA-2048+ or Ed25519 for asymmetric operations, and SHA-256+ for hashing. Always use the `cryptography` library with secure defaults rather than implementing custom crypto or using deprecated libraries.

## Remediation Strategy

- Use AES-256-GCM for symmetric encryption, never DES/3DES/RC4
- Use RSA-2048+ or elliptic curve cryptography (Ed25519, secp256r1) for asymmetric operations
- Use SHA-256 or SHA-3 for hashing, never MD5 or SHA-1
- Leverage the `cryptography` library's high-level recipes with secure defaults
- Generate cryptographically secure random keys using `os.urandom()` or `secrets` module

## Remediation Steps

- Replace `PyCrypto` or `pycryptodome` with the actively maintained `cryptography` library
- Audit all encryption code for algorithm selection - replace DES/3DES with AES-256
- Verify key sizes meet current standards - 256-bit for AES, 2048-bit minimum for RSA
- Use authenticated encryption modes (GCM, ChaCha20-Poly1305) instead of CBC without HMAC
- Generate keys with `Fernet.generate_key()` or `os.urandom(32)` for 256-bit strength
- Review hash functions and replace MD5/SHA-1 with SHA-256 or newer

## Minimal Safe Pattern

```python
from cryptography.fernet import Fernet
import os

# Generate a strong 256-bit key
key = Fernet.generate_key()  # Or store securely
fernet = Fernet(key)

# Encrypt data (uses AES-128-CBC + HMAC-SHA256)
plaintext = b"Sensitive data"
ciphertext = fernet.encrypt(plaintext)

# Decrypt data
decrypted = fernet.decrypt(ciphertext)

# For AES-256-GCM directly
from cryptography.hazmat.primitives.ciphers.aead import AESGCM
key_256 = AESGCM.generate_key(bit_length=256)
aesgcm = AESGCM(key_256)
nonce = os.urandom(12)
ciphertext = aesgcm.encrypt(nonce, plaintext, None)
```
