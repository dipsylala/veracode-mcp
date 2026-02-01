# CWE-331: Insufficient Entropy - Java

## LLM Guidance

Insufficient entropy in Java occurs when using `java.util.Random` or `Math.random()` instead of `java.security.SecureRandom` for cryptographic operations. The `Random` class produces predictable sequences that can be reproduced if the seed is known, making it unsuitable for security-sensitive operations like generating session tokens, encryption keys, or initialization vectors. **Use `java.security.SecureRandom` for all security-sensitive random value generation.**

## Remediation Strategy

- Always use `SecureRandom` for cryptographic purposes; never use `Random` or `Math.random()` for security-sensitive operations
- Prefer `SecureRandom.getInstanceStrong()` for maximum entropy when performance is not critical
- Generate sufficient entropy: minimum 128 bits (16 bytes) for tokens/IVs, 256 bits (32 bytes) for encryption keys
- Avoid manual seeding unless using a cryptographically strong entropy source
- Properly encode random bytes (Base64, hex) for safe transmission and storage

## Remediation Steps

- Replace all instances of `new Random()` with `new SecureRandom()` in security contexts
- Replace `Math.random()` calls with `SecureRandom.nextDouble()` or equivalent methods
- Use `SecureRandom.getInstanceStrong()` for critical operations like key generation
- Generate adequate bytes - `byte[] token = new byte[32]; secureRandom.nextBytes(token);`
- Encode output for use - `Base64.getUrlEncoder().withoutPadding().encodeToString(token)`
- Review all random value generation in authentication, encryption, and token creation code

## Minimal Safe Pattern

```java
import java.security.SecureRandom;
import java.util.Base64;

public class SecureTokenGenerator {
    private static final SecureRandom secureRandom = new SecureRandom();
    
    public static String generateSessionToken() {
        byte[] randomBytes = new byte[32]; // 256 bits
        secureRandom.nextBytes(randomBytes);
        return Base64.getUrlEncoder()
                     .withoutPadding()
                     .encodeToString(randomBytes);
    }
}
```
