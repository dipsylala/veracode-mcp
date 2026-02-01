# CWE-327: Use of a Broken or Risky Cryptographic Algorithm

## LLM Guidance

Weak or broken cryptographic algorithms fail to protect data confidentiality, integrity, and authenticity. Modern computing power can break algorithms like MD5, DES, and SHA-1 in seconds to minutes, allowing attackers to decrypt data, forge signatures, or crack passwords. Applications must use strong, modern cryptographic standards to ensure security.

## Key Principles

- Do not use broken or deprecated cryptography (MD5, DES, 3DES, SHA-1, RC4)
- Cryptographic algorithm selection must be centrally defined and server-controlled
- Use only algorithms that remain cryptographically sound with current computing power
- Separate concerns: use encryption for confidentiality, hashing for integrity, and password hashing for credential storage
- Configure TLS/SSL to use modern protocols (TLS 1.2+) and strong cipher suites

## Remediation Steps

- Identify the weak algorithm in use (DES, MD5, SHA-1, etc.) from flaw details including file and line number
- Determine the cryptographic purpose - encryption, hashing, password hashing, digital signatures, or TLS configuration
- Replace weak algorithms with approved alternatives - AES-256-GCM for encryption, SHA-256/SHA-3 for hashing, bcrypt/Argon2 for passwords
- Update TLS/SSL configurations to disable outdated protocols (SSL, TLS 1.0/1.1) and weak cipher suites
- Use vetted cryptographic libraries (NaCl, libsodium, OpenSSL 1.1+) rather than implementing custom cryptography
- Test thoroughly to ensure the replacement algorithm functions correctly without breaking existing functionality
