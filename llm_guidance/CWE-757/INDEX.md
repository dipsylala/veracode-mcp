# CWE-757: Selection of Less-Secure Algorithm

## LLM Guidance

Using weak cryptographic algorithms (MD5, SHA-1, DES, RC4, weak RSA keys) instead of strong modern alternatives exposes data to collision attacks, brute force, and cryptanalysis, compromising confidentiality and integrity. The core fix is replacing legacy algorithms with modern, secure primitives appropriate for each use case, with algorithm selection server-controlled via an approved allowlist.

## Key Principles

- Use dedicated password hashing functions (bcrypt, Argon2, scrypt) instead of general hash functions for passwords
- Use SHA-256 or SHA-512 for file integrity checks and non-password hashing needs
- Use AES-256-GCM or ChaCha20-Poly1305 for symmetric encryption
- Use RSA-2048+ with SHA-256 or ECDSA for digital signatures
- Enforce server-controlled algorithm selection restricted to approved modern primitives

## Remediation Steps

- Identify weak algorithms in codebase (MD5, SHA-1, DES, RC4, weak RSA keys)
- Select appropriate replacement based on use case - passwords → bcrypt/Argon2, integrity → SHA-256+, encryption → AES-256-GCM, signatures → RSA-2048+/ECDSA
- Plan migration strategy for existing hashed or encrypted data
- Update code to use strong algorithms with proper configuration
- Test thoroughly including security validation and compatibility
- Monitor continuously for any continued use of weak algorithms
