# CWE-78: OS Command Injection - Java

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into OS commands without proper validation, allowing attackers to execute arbitrary commands on the host system. The primary defense is to use Java native APIs (Files, HttpClient, etc.) instead of system commands to eliminate the vulnerability entirely. If system commands are unavoidable, use ProcessBuilder with argument arrays and never invoke a shell or concatenate strings.

## Key Principles

- Use Java native APIs instead of system calls (Files.copy, HttpClient, ProcessHandle) to eliminate command injection risk
- Use ProcessBuilder with argument arrays if commands are unavoidable—separate command from arguments, never use string concatenation or shell interpretation
- Apply allowlist input validation as defense-in-depth, never as sole protection
- Run with least privilege by configuring minimal OS permissions for processes
- Never use Runtime.exec() or shell invocation with any user-controlled input

## Remediation Steps

- Identify all Runtime.exec() and ProcessBuilder calls in the codebase
- Replace system commands with Java native APIs where possible (Files, HttpClient, etc.)
- Refactor unavoidable commands to use ProcessBuilder with separate command and argument array parameters
- Implement strict allowlist validation (permitted characters/values only) for any user input
- Configure process execution with minimal required OS permissions
- Remove all shell invocation (sh, cmd.exe, bash) and string concatenation patterns

## Safe Pattern

```java
// PREFERRED: Use Java native APIs
Path source = Paths.get(validatedUserPath);
Path dest = Paths.get("/backup/");
Files.copy(source, dest, StandardCopyOption.REPLACE_EXISTING);

// If unavoidable: ProcessBuilder with argument array
ProcessBuilder pb = new ProcessBuilder(
    "convert",           // command
    validatedInputFile,  // argument 1
    validatedOutputFile  // argument 2
);
Process process = pb.start();
```
