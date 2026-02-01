# CWE-377: Insecure Temporary File - Java

## LLM Guidance

Insecure temporary file creation occurs when applications create files with predictable names, weak permissions, or in shared directories without proper safeguards, enabling symlink attacks and data tampering. Always use `Files.createTempFile()` with restrictive permissions and ensure proper cleanup.

## Key Principles

- Use `Files.createTempFile()` or `File.createTempFile()` instead of manual path construction
- Set permissions to owner-only (600) using `PosixFilePermissions` before writing sensitive data
- Ensure automatic cleanup with try-finally blocks or `deleteOnExit()`
- Never create temp files with predictable or hardcoded names

## Remediation Steps

- Replace manual file creation with `Files.createTempFile("prefix-", ".tmp")`
- Apply `PosixFilePermissions.fromString("rw-------")` for owner-only access
- Wrap operations in try-finally to guarantee deletion via `Files.deleteIfExists()`
- Validate temp directory permissions before use
- Consider in-memory alternatives for highly sensitive data

## Safe Pattern

```java
import java.nio.file.*;
import java.nio.file.attribute.*;

Path tempFile = Files.createTempFile("secure-", ".tmp",
    PosixFilePermissions.asFileAttribute(
        PosixFilePermissions.fromString("rw-------")));

try {
    Files.write(tempFile, sensitiveData);
    // Process file
} finally {
    Files.deleteIfExists(tempFile);
}
```
