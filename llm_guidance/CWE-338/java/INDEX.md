# CWE-338: Use of Cryptographically Weak Pseudo-Random Number Generator (PRNG) - Java

## LLM Guidance

Java's `java.util.Random` and `Math.random()` are predictable and unsuitable for security operations like generating tokens, keys, or passwords. Attackers can predict outputs and compromise security-sensitive values. Always use `java.security.SecureRandom` for cryptographic purposes.

## Remediation Strategy

- Replace all `java.util.Random` and `Math.random()` with `SecureRandom` in security-sensitive contexts
- Use strong algorithm providers (e.g., "NativePRNGNonBlocking", "SHA1PRNG") when available
- Initialize `SecureRandom` once and reuse the instance to avoid performance overhead
- Never seed `SecureRandom` with predictable values (timestamps, constants)
- Ensure sufficient entropy by relying on OS-level random sources

## Remediation Steps

- Identify all random number generation in security-sensitive code (tokens, keys, salts, IVs, nonces)
- Replace `new Random()` with `new SecureRandom()` or `SecureRandom.getInstanceStrong()`
- Replace `Math.random()` calls with `SecureRandom.nextDouble()` or equivalent methods
- Remove any manual seeding with `setSeed()` unless using truly random entropy
- Validate that the default provider offers cryptographic strength for your platform
- Test for performance impact and optimize instance reuse if needed

## Minimal Safe Pattern

```java
import java.security.SecureRandom;

public class SecureTokenGenerator {
    private static final SecureRandom secureRandom = new SecureRandom();
    
    public static String generateToken() {
        byte[] randomBytes = new byte[32];
        secureRandom.nextBytes(randomBytes);
        return bytesToHex(randomBytes);
    }
    
    private static String bytesToHex(byte[] bytes) {
        StringBuilder sb = new StringBuilder();
        for (byte b : bytes) {
            sb.append(String.format("%02x", b));
        }
        return sb.toString();
    }
}
```
