# CWE-77: Command Injection - Java

## LLM Guidance

Command injection occurs when applications construct system commands using untrusted input without proper sanitization, allowing attackers to inject shell metacharacters and execute arbitrary commands.

**Primary Defence:** Use Java native APIs (Files, HttpClient, ProcessHandle, etc.) instead of executing system commands to eliminate the vulnerability entirely. If system commands are unavoidable, use `ProcessBuilder` with argument arrays and never invoke a shell.

## Key Principles

- **BEST:** Use Java native APIs (Files, HttpClient, ProcessHandle) instead of system commands to eliminate command injection risk
- **If commands unavoidable:** Use `ProcessBuilder` with argument arrays and never invoke shell interpreters
- Avoid shell interpreters (`sh -c`, `cmd /c`) that enable metacharacter injection
- Never concatenate user input directly into command strings
- Validate all user input against strict allowlists as defence-in-depth
- Apply principle of least privilege to command execution processes

## Remediation Steps

- **Replace system commands with Java native APIs** (Files, HttpClient, etc.) to eliminate vulnerability
- If commands are unavoidable, replace `Runtime.exec(String)` calls with `ProcessBuilder` using separate arguments
- Remove all shell interpreter invocations (`sh -c`, `cmd /c`) from command construction
- Implement allowlist validation for any user-supplied input as additional defence
- Use hardcoded command paths and restrict argument values to known-safe options
- Apply runtime security policies to limit command execution capabilities

## Safe Pattern

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
