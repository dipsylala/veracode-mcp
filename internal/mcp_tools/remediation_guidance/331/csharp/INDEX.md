# CWE-331: Insufficient Entropy - C#

## LLM Guidance

Insufficient entropy in C# occurs when using `System.Random` or `Guid.NewGuid()` for security-sensitive operations instead of cryptographically secure alternatives. `Random` produces predictable sequences unsuitable for generating tokens, passwords, keys, or session identifiers. Always use `RandomNumberGenerator` (.NET 6+) or `RNGCryptoServiceProvider` (older .NET) for security-critical random values.

## Key Principles

- Replace all `System.Random` instances in security contexts with `RandomNumberGenerator`
- Never use `Guid.NewGuid()` for authentication tokens, session IDs, or cryptographic keys
- Use cryptographically secure methods for generating salts, nonces, IVs, and secrets
- Avoid seeding patterns that reduce entropy (e.g., time-based seeds with `Random`)

## Remediation Steps

- Identify all uses of `Random`, `Guid.NewGuid()`, or predictable value generation in security code
- Replace with `RandomNumberGenerator.GetBytes()` or `RandomNumberGenerator.Fill()`
- For random strings, encode secure bytes using Base64 or hex encoding
- Update token/session generation to use cryptographic RNG exclusively
- Review and test authentication, encryption, and access control flows

## Safe Pattern

```csharp
using System.Security.Cryptography;

// Generate cryptographically secure random bytes
byte[] randomBytes = new byte[32];
RandomNumberGenerator.Fill(randomBytes);

// Convert to token string
string secureToken = Convert.ToBase64String(randomBytes);

// For random integers in a range
int secureValue = RandomNumberGenerator.GetInt32(0, 100);
```
