# CWE-316: Cleartext Storage of Sensitive Information in Memory - Java

## LLM Guidance

Storing sensitive data (passwords, cryptographic keys, tokens) in memory as cleartext exposes it to heap dumps, debuggers, and memory disclosure vulnerabilities. Java strings are immutable and persist in the string pool, making them particularly dangerous. Use `char[]` for passwords, clear arrays explicitly with `Arrays.fill()`, and avoid string operations on credentials.

## Key Principles

- Use mutable data structures: Prefer `char[]` or `byte[]` over `String` for sensitive data
- Clear immediately after use: Zero out arrays in `finally` blocks to ensure cleanup
- Minimize lifetime: Process and discard sensitive data as quickly as possible
- Avoid string conversions: Never call `new String(charArray)` or similar on credentials
- Use secure APIs: Leverage `javax.crypto.SecretKey`, `java.security.KeyStore`, and `Destroyable` interfaces

## Remediation Steps

- Replace `String password` parameters with `char[] password`
- Add `Arrays.fill(password, '\0')` in `finally` blocks after processing
- Remove any `.toString()`, string concatenation, or logging of sensitive values
- Use `SecureString` or `GuardedString` wrappers where available
- Implement `AutoCloseable` or `Destroyable` for credential holder classes
- Review heap dump and debugging configurations to prevent memory exposure

## Safe Pattern

```java
char[] password = null;
try {
    password = getPasswordFromUser();
    byte[] hash = hashPassword(password);
    authenticate(hash);
} finally {
    if (password != null) {
        Arrays.fill(password, '\0');
    }
}
```
