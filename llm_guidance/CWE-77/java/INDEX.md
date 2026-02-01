# CWE-77: Command Injection - Java

## LLM Guidance

Command injection occurs when applications construct system commands using untrusted input without proper sanitization, allowing attackers to inject shell metacharacters and execute arbitrary commands. The primary defense is using `ProcessBuilder` with separate arguments instead of concatenating strings with `Runtime.exec()`, avoiding shell invocation entirely.

## Remediation Strategy

- Use `ProcessBuilder` with argument arrays to prevent shell interpretation
- Avoid shell interpreters (`sh -c`, `cmd /c`) that enable metacharacter injection
- Validate all user input against strict allowlists before command construction
- Never concatenate user input directly into command strings
- Apply principle of least privilege to command execution processes

## Remediation Steps

- Replace `Runtime.exec(String)` calls with `ProcessBuilder` using separate arguments
- Remove shell interpreter invocations from command construction
- Implement allowlist validation for any user-supplied input
- Escape or reject special characters if shell invocation is unavoidable
- Use hardcoded command paths and restrict argument values to known-safe options
- Apply runtime security policies to limit command execution capabilities

## Minimal Safe Pattern

```java
// Safe: ProcessBuilder with separate arguments
public void processFile(String userFilename) {
    // Validate against allowlist
    if (!userFilename.matches("^[a-zA-Z0-9_-]+\\.txt$")) {
        throw new IllegalArgumentException("Invalid filename");
    }
    
    ProcessBuilder pb = new ProcessBuilder(
        "/usr/bin/convert",
        userFilename,
        "output.pdf"
    );
    pb.start();
}
```
