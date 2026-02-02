# CWE-502: Insecure Deserialization - Java

## LLM Guidance

Insecure deserialization occurs when untrusted data is used to create objects, potentially allowing attackers to execute arbitrary code, manipulate application logic, or achieve denial of service. Java's native serialization is particularly dangerous because it can invoke methods during deserialization. **Primary Defence:** Use JSON (Jackson, Gson) instead of Java serialization, or if Java serialization is required, implement `ObjectInputFilter` (Java 9+) to whitelist allowed classes.

## Key Principles

- Prefer data-only formats: Replace Java serialization with JSON, Protocol Buffers, or other data-only formats that don't execute code during deserialization
- Whitelist classes explicitly: If Java serialization is unavoidable, use `ObjectInputFilter` to allow only specific, known-safe classes
- Never trust serialized data: Treat all serialized input as untrusted, even from seemingly secure sources
- Apply defense in depth: Combine multiple controls including input validation, least privilege, and monitoring

## Remediation Steps

- Replace `ObjectInputStream` with JSON parsers like Jackson or Gson for data transfer
- Implement `ObjectInputFilter` to whitelist allowed classes by package/class name
- Validate and sanitize all input before deserialization
- Mark classes as `final` or use sealed class hierarchies to prevent subclass exploitation
- Update dependencies regularly to patch known deserialization gadgets
- Monitor and log all deserialization activity for anomaly detection

## Safe Pattern

```java
// ObjectInputFilter whitelist approach (Java 9+)
ObjectInputFilter filter = ObjectInputFilter.Config.createFilter(
    "com.example.SafeClass;com.example.SafeDTO;!*"
);

try (ObjectInputStream ois = new ObjectInputStream(inputStream)) {
    ois.setObjectInputFilter(filter);
    Object obj = ois.readObject();
    
    if (obj instanceof SafeClass) {
        SafeClass safeObj = (SafeClass) obj;
        // Process safely
    }
}
```
