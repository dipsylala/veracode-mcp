# CWE-916: Use of Password Hash With Insufficient Computational Effort

## LLM Guidance

Weak password hashing uses fast cryptographic functions (MD5, SHA-1, SHA-256) or poorly configured algorithms that attackers can brute-force when database dumps are compromised. Password hashing requires intentionally slow, computationally expensive algorithms designed specifically for password storage to resist offline cracking attacks. Use adaptive algorithms like Argon2id, bcrypt, or scrypt with work factors tuned to current hardware.

## Key Principles

- Use adaptive, purpose-built password hashing algorithms (Argon2id, bcrypt, scrypt)
- Never use fast general-purpose hashes (MD5, SHA-1, SHA-256) for passwords
- Tune work factors to current hardware (target 250-500ms per hash)
- Balance security (slow enough to resist brute force) with usability (fast enough for legitimate authentication)
- Implement password migration strategy when upgrading from weak algorithms

## Remediation Steps

- Identify the weak hashing algorithm in use (MD5, SHA-1, SHA-256, unsalted hash)
- Review flaw details for specific file, line number, and code pattern
- Trace password flow from user registration/login through hashing to storage
- Check database schema for password hash column and verify salt storage
- Replace weak algorithm with Argon2id (preferred), bcrypt, or scrypt
- Configure appropriate work factors and migrate existing password hashes
