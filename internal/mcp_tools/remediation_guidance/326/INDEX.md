# CWE-326: Inadequate Encryption Strength

## LLM Guidance

Inadequate Encryption Strength occurs when cryptographic algorithms or key sizes are too weak to provide effective protection, allowing attackers to break encryption and access sensitive data. The core fix is to use strong, industry-standard algorithms with appropriate key sizes and ensure cryptographic strength is server-controlled, not determined by legacy compatibility or client input.

## Key Principles

- Never allow cryptographic strength to be determined by legacy compatibility or client input
- Cryptographic algorithms, protocols, and key sizes must be centrally defined and server-controlled
- Constrain all cryptographic operations to secure minimums based on current industry standards
- Replace weak algorithms (DES, 3DES, RC4, MD5, SHA-1) with strong alternatives (AES-256, SHA-256/SHA-3)

## Remediation Steps

- Review flaw details to identify where weak cryptographic algorithms or key sizes are used in your code
- Identify weak algorithms - DES, 3DES, RC4, MD5, SHA-1, AES-128 (when 256-bit is required)
- Verify minimum key sizes - RSA ≥ 2048 bits, AES ≥ 256 bits, ECC ≥ 256 bits
- Use AES-256-GCM for encryption (not DES, 3DES, RC4, or AES-128 for sensitive data)
- Use SHA-256 or SHA-3 for hashing (not MD5 or SHA-1)
- Implement centralized cryptographic configuration that enforces secure algorithm and key size minimums
