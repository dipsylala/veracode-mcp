# CWE-328 - Use of Weak Hash

## LLM Guidance

Weak cryptographic hashes (MD5, SHA-1) are broken and vulnerable to collision attacks, enabling attackers to forge signatures, create malicious files with same hash, and crack hashed passwords. Modern applications must use SHA-256+ for integrity and bcrypt/Argon2 for passwords.

## Key Principles

- Use purpose-appropriate hashing - bcrypt/Argon2 for passwords, SHA-256+ for integrity
- Never use fast hashes (MD5, SHA-1, plain SHA-256) for password storage or security-critical operations
- Apply key derivation functions with sufficient iteration counts (bcrypt cost 12+, PBKDF2 600k+ iterations)
- Upgrade legacy systems by rehashing on user login without forcing password resets
- Use SHA-256 or SHA-3 for file integrity, digital signatures, and non-password use cases

## Remediation Steps

- Identify weak hash usage - Review flaw details for file/line using MD5, SHA-1, or plain SHA-256; determine purpose (passwords, integrity, signatures)
- Replace password hashing - Migrate to bcrypt (cost 12+), Argon2id, or PBKDF2-HMAC-SHA256 (600k+ iterations); never use fast hashes for passwords
- Upgrade integrity checks - Replace MD5/SHA-1 with SHA-256, SHA-384, or SHA-3 for file verification and checksums
- Update cryptographic operations - Use SHA-256+ for HMACs, digital signatures, and key derivations
- Test thoroughly - Verify backward compatibility, test authentication flows, and validate integrity checks with new algorithms
- Deploy rehashing strategy - For legacy systems, rehash passwords during user login to migrate gradually without forcing resets
