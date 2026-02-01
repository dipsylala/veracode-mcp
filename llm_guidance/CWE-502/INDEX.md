# CWE-502: Insecure Deserialization

## LLM Guidance

Insecure Deserialization occurs when applications deserialize untrusted data without validation, allowing attackers to manipulate serialized objects to execute arbitrary code, modify logic, or access unauthorized data. Formats like Java ObjectInputStream, Python pickle, PHP serialize(), and .NET BinaryFormatter can instantiate arbitrary classes during deserialization. Never allow untrusted data to be deserialized into executable objects; enforce integrity and type safety before object creation.

## Key Principles

- Replace native deserialization with safe data formats (JSON, XML with schema validation)
- Implement cryptographic integrity checks (HMAC signatures) on all serialized data
- Enforce strict type whitelisting and class instantiation controls
- Isolate deserialization operations in sandboxed, low-privilege environments
- Apply defense-in-depth: validation, monitoring, and runtime restrictions

## Remediation Steps

- Identify the deserialization call location, serialized data source, and format used
- Trace complete data flow from origin to deserialization operation
- Verify if attacker-controlled data reaches deserialization without integrity checks
- Replace unsafe formats with JSON/XML and implement schema validation
- Add HMAC signature verification before any deserialization attempts
- Apply type whitelisting to restrict instantiable classes
