# CWE-256 - Plaintext Storage of a Password

## LLM Guidance

Storing passwords in plaintext (database, files, configuration, logs) exposes all user credentials if storage is compromised. Passwords must be hashed with strong algorithms (bcrypt, Argon2, PBKDF2) using salts, making them computationally infeasible to reverse even if the database is stolen.

## Key Principles

- Use strong, salted, slow hashing algorithms (Argon2, bcrypt, scrypt) exclusively
- Never store passwords in reversible form (plaintext, encoding, weak hashing, encryption)
- Implement proper work factors to resist brute-force attacks
- Use unique salts per password to prevent rainbow table attacks
- Never log, cache, or expose passwords in API responses

## Remediation Steps

- Check database schema - Identify VARCHAR password columns that should contain hashed values
- Review configuration files - Remove passwords from .properties, .env, and config files
- Search codebase for logging - Eliminate password variables from log statements
- Audit client-side storage - Remove passwords from cookies and session storage
- Review API responses - Ensure passwords are never returned in responses
- Replace plaintext storage - Implement bcrypt/Argon2 hashing with automatic salt generation
- Migrate existing passwords - Hash plaintext passwords on next user login, prompt password reset if needed
