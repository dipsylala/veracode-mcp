# CWE-329: Generation of Predictable IV with CBC Mode

## LLM Guidance

CBC mode requires a random, unpredictable Initialization Vector (IV) for each encryption operation. Using static, sequential, or predictable IVs enables plaintext recovery through IV manipulation attacks, completely breaking encryption security even with the correct key.

## Key Principles

- Never reuse IVs when encrypting with CBC mode and the same key
- Always generate IVs using cryptographically secure random number generators
- Generate a fresh, unpredictable IV for each encryption operation
- Store or transmit the IV alongside the ciphertext (IV does not need to be secret)
- Ensure IV size matches the cipher block size (typically 16 bytes for AES)

## Remediation Steps

- Locate predictable IV usage - Review flaw details to identify the specific file, line number, and code pattern using CBC mode with static, sequential, or reused IVs
- Identify current IV generation - Determine if IV is hardcoded, derived from a counter/timestamp, or reused across encryptions
- Use cryptographically secure randomness - Replace with `SecureRandom` (Java), `os.urandom()` (Python), or `crypto.randomBytes()` (Node.js)
- Generate fresh IV per encryption - Create a new random IV immediately before each encryption operation
- Store IV with ciphertext - Prepend or append the IV to the encrypted output for use during decryption
- Verify IV uniqueness - Ensure no IV is ever reused with the same encryption key across different messages
