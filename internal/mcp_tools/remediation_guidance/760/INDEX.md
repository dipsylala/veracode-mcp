# CWE-760: Use of One-Way Hash Without Salt

## LLM Guidance

Hashing passwords without unique salts enables rainbow table attacks where attackers precompute hashes for common passwords and instantly crack unsalted hashes, compromising all accounts using the same password. The core fix is using modern password hashing algorithms with automatic salt generation.

## Key Principles

- Use modern password hashing functions (bcrypt, Argon2, scrypt, or PBKDF2 with 600,000+ iterations) that automatically generate unique random salts for each password
- Never use fast hash functions (MD5, SHA-1, SHA-256) for passwords as they are extremely fast to crack
- Password hashing algorithms should be computationally expensive with adjustable work factors
- Store algorithm, cost parameters, salt, and hash together

## Remediation Steps

- Identify current hashing method in your codebase
- Choose appropriate password hashing function (prefer Argon2 or bcrypt)
- Implement password hashing with automatic salt generation
- Plan migration strategy for existing password hashes
- Implement dual-read verification during migration
- Test with sample passwords and monitor migration progress
