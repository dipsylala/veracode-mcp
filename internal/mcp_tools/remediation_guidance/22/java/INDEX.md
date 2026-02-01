# CWE-22: Path Traversal - Java

## LLM Guidance

Path Traversal occurs when user input constructs file paths without validation, allowing attackers to use `../` sequences or absolute paths to access files outside the intended directory. This can expose sensitive files like `/etc/passwd` or `WEB-INF/web.xml`. The primary defense is using indirect reference mapping (map IDs to filenames) or validating with `Path.normalize()` and checking the canonical path stays within allowed directories.

## Key Principles

- Use indirect reference maps instead of accepting filenames directly from users
- Validate canonical paths remain within the intended base directory
- Reject paths containing traversal sequences (`../`, `..\\`) or null bytes
- Use allowlists for permitted file extensions and directories
- Avoid constructing paths from untrusted input when possible

## Remediation Steps

- Implement indirect object references (user provides ID, application maps to filename)
- Canonicalize user input with `File.getCanonicalPath()` or `Path.normalize()`
- Verify the resolved path starts with the intended base directory
- Reject requests with traversal sequences, absolute paths, or suspicious characters
- Apply allowlist validation for file extensions if direct input is unavoidable
- Use security manager or sandboxing to restrict file system access

## Safe Pattern

```java
public File getSecureFile(String userInput, String baseDir) throws IOException {
    File base = new File(baseDir).getCanonicalFile();
    File requested = new File(base, userInput).getCanonicalFile();
    
    if (!requested.getPath().startsWith(base.getPath())) {
        throw new SecurityException("Path traversal detected");
    }
    return requested;
}
```
