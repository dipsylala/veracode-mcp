# CWE-780: Use of RSA Without OAEP

## LLM Guidance

Using RSA encryption without OAEP (Optimal Asymmetric Encryption Padding) enables padding oracle attacks, chosen ciphertext attacks, and message malleability. OAEP adds randomness and integrity checks, making RSA encryption secure against modern attacks. Always use RSA-OAEP instead of raw RSA or PKCS#1 v1.5 padding.

## Key Principles

- Always specify OAEP padding when using RSA encryption with SHA-256 or better (not SHA-1) and MGF1 mask generation function
- Use hybrid encryption for large data: generate random AES-256 key, encrypt data with AES-GCM/ChaCha20-Poly1305, encrypt symmetric key with RSA-OAEP
- Never use "default" RSA without explicitly specifying OAEP padding mode
- Choose appropriate key sizes with minimum 2048-bit RSA keys for current security requirements

## Remediation Steps

- Locate RSA encryption calls in codebase and identify padding mode used
- Update cipher configuration to explicitly specify OAEP padding (e.g., `RSA/ECB/OAEPWithSHA-256AndMGF1Padding`)
- Verify hash algorithms use SHA-256 or better for OAEP and MGF1
- Implement hybrid encryption for data larger than key size minus padding overhead
- Test encryption/decryption end-to-end to ensure compatibility
- Scan for other RSA usage patterns that may need similar fixes
