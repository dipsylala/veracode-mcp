# CWE-522: Insufficiently Protected Credentials - JavaScript

## LLM Guidance

Insufficiently Protected Credentials in JavaScript/Node.js occurs when passwords, API keys, tokens, or secrets are stored in plaintext, hardcoded in source code, committed to version control, or transmitted insecurely. The core fix is to use strong one-way hashing (bcrypt, argon2) for passwords, environment variables for secrets, and encrypted transmission channels. Never store plaintext credentials or commit secrets to repositories.

## Key Principles

- Hash passwords with bcrypt or argon2 (never store plaintext or use weak algorithms like MD5/SHA1)
- Store secrets in environment variables or secure vaults (AWS Secrets Manager, Azure Key Vault, HashiCorp Vault)
- Use HTTPS/TLS for all credential transmission
- Add `.env` files to `.gitignore` and scan repositories for committed secrets
- Implement credential rotation policies and use short-lived tokens where possible

## Remediation Steps

- Replace plaintext password storage with bcrypt hashing (minimum 12 rounds)
- Move all hardcoded credentials to environment variables loaded via `dotenv` or platform secrets
- Add `.env`, `config.js`, and credential files to `.gitignore`
- Use `git-secrets` or `truffleHog` to scan repository history for leaked credentials
- Implement HTTPS for APIs and use secure cookie flags (`httpOnly`, `secure`, `sameSite`)
- Rotate any credentials that were previously exposed

## Safe Pattern

```javascript
const bcrypt = require('bcrypt');
require('dotenv').config();

// Hash password on registration
async function registerUser(plainPassword) {
  const hashedPassword = await bcrypt.hash(plainPassword, 12);
  // Store hashedPassword in database
}

// Verify password on login
async function loginUser(plainPassword, hashedPassword) {
  return await bcrypt.compare(plainPassword, hashedPassword);
}

// Load API keys from environment
const apiKey = process.env.API_KEY;
```
