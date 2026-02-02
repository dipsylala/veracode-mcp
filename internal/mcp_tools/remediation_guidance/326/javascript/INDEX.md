# CWE-326: Inadequate Encryption Strength - JavaScript/Node.js

## LLM Guidance

Inadequate Encryption Strength in JavaScript/Node.js applications occurs when developers use weak cryptographic algorithms (MD5, DES, RC4), insufficient key sizes, or deprecated ciphers that fail to protect sensitive data against modern attacks. The Node.js `crypto` module provides both secure and insecure options-always use AES-256-GCM, ChaCha20-Poly1305, or modern algorithms with proper key derivation (PBKDF2, scrypt, Argon2).

## Key Principles

- Use AES-256-GCM or ChaCha20-Poly1305 for symmetric encryption with authenticated encryption modes
- Generate keys with cryptographically secure random sources (`crypto.randomBytes()`) at minimum 256-bit length
- Derive keys from passwords using PBKDF2 (100,000+ iterations), scrypt, or Argon2
- Avoid deprecated algorithms: DES, 3DES, RC4, MD5, SHA1, AES-ECB mode

## Remediation Steps

- Replace weak ciphers (DES, RC4, AES-128) with AES-256-GCM or ChaCha20-Poly1305
- Generate 256-bit keys using `crypto.randomBytes(32)` or derive from passwords with `crypto.pbkdf2()` (100,000+ iterations)
- Use authenticated encryption modes (GCM, CCM) that provide both confidentiality and integrity
- Generate unique IVs/nonces per encryption operation using `crypto.randomBytes(12)` for GCM
- Store algorithm, IV, salt, and ciphertext together; never hardcode encryption keys

## Safe Pattern

```javascript
const crypto = require('crypto');

function encryptData(plaintext, password) {
  const salt = crypto.randomBytes(16);
  const key = crypto.pbkdf2Sync(password, salt, 100000, 32, 'sha256');
  const iv = crypto.randomBytes(12);
  const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);
  
  const encrypted = Buffer.concat([cipher.update(plaintext, 'utf8'), cipher.final()]);
  const authTag = cipher.getAuthTag();
  
  return Buffer.concat([salt, iv, authTag, encrypted]);
}
```
