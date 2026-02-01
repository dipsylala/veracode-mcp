# CWE-326 - Inadequate Encryption Strength - C\#

## LLM Guidance

Inadequate Encryption Strength occurs when C# applications use weak algorithms (DES, TripleDES, MD5, SHA1), short keys, or unauthenticated cipher modes (ECB, CBC without MAC). Modern .NET (6+) provides strong primitives like `AesGcm`, `HMACSHA256`, and `RandomNumberGenerator`, but legacy APIs and outdated examples still introduce weak cryptography. The core fix - migrate to AES-256-GCM for encryption and SHA-256/SHA-512 for hashing.

## Key Principles

- Use authenticated encryption - AES-GCM or ChaCha20-Poly1305 prevent tampering
- Generate cryptographic keys properly - `RandomNumberGenerator`, never `new Random()` or hardcoded values
- Enforce minimum key sizes - AES-256 (32 bytes), RSA-2048+, PBKDF2 ≥100k iterations
- Avoid deprecated algorithms - No DES, TripleDES, MD5, SHA1, RC2, or ECB mode
- Separate keys per purpose - Don't reuse encryption keys for signing/authentication

## Remediation Steps

- Search for weak algorithms - Grep for `DES`, `TripleDES`, `MD5`, `SHA1`, `RC2`, `RSA.Create(1024)`
- Replace symmetric encryption with `AesGcm` using 32-byte keys from `RandomNumberGenerator`
- Upgrade hashing from MD5/SHA1 to `SHA256` or `SHA512`
- Fix key derivation - Use `Rfc2898DeriveBytes` with ≥100k iterations and random salts
- Add authentication if using CBC - apply HMAC-SHA256 (encrypt-then-MAC pattern)
- Validate key storage - Ensure keys are in secure vaults (Azure Key Vault, DPAPI), not config files

## Safe Pattern

```csharp
using System.Security.Cryptography;

byte[] key = RandomNumberGenerator.GetBytes(32);
byte[] nonce = RandomNumberGenerator.GetBytes(AesGcm.NonceByteSizes.MaxSize);
byte[] plaintext = Encoding.UTF8.GetBytes("sensitive data");
byte[] ciphertext = new byte[plaintext.Length];
byte[] tag = new byte[AesGcm.TagByteSizes.MaxSize];

using var aes = new AesGcm(key, AesGcm.TagByteSizes.MaxSize);
aes.Encrypt(nonce, plaintext, ciphertext, tag);
// Store - nonce + ciphertext + tag
```
