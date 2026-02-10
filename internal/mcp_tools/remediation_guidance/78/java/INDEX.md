# CWE-78: OS Command Injection - Java

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into OS commands without proper validation, allowing attackers to execute arbitrary commands on the host system.

**CRITICAL: DO NOT suggest input validation/sanitization as the primary fix.** The correct remediation is to **ELIMINATE the system command entirely** by using Java's built-in APIs. Input validation alone is insufficient and error-prone.

**Primary Defence:** Replace ALL system command calls with Java native library alternatives. Common replacements:

- **file operations → java.nio.file.Files API**
- **network requests → HttpClient or HttpURLConnection**
- **archive operations → java.util.zip or Apache Commons Compress**

Only if no library alternative exists, then use ProcessBuilder with separate argument arrays (never Runtime.exec() or shell invocation).

## Key Principles

- **ALWAYS prefer library-based solutions over system commands** - this eliminates injection risk entirely
- For file operations, use **java.nio.file.Files** API instead of system commands like cp, mv, chmod
- For archives, use **java.util.zip** or **Apache Commons Compress** instead of tar/zip commands
- **Input validation is NOT a remediation** - it's only defense-in-depth after eliminating command execution
- If commands are truly unavoidable, use ProcessBuilder with argument arrays - never concatenate strings or invoke shells (sh, bash, cmd.exe)
- Run with least privilege by configuring minimal OS permissions for processes

## Remediation Steps

- Identify all Runtime.exec(), ProcessBuilder, and exec family calls
- **First, determine the library-based replacement** (see mapping below) - do not default to "secure" command execution
- Replace system commands with appropriate Java APIs
- For any remaining unavoidable commands, refactor to use ProcessBuilder with separate argument array parameters (command and arguments separated)
- Remove all shell invocation patterns (Runtime.exec("sh -c"), Runtime.exec("bash -c"), Runtime.exec("cmd /c"))
- Configure minimal OS permissions for any remaining process executions

## Remediation Pattern

**The transformation pattern:** System Command → Java Library API

This applies to ALL system commands - if you find Runtime.exec() or ProcessBuilder executing a command, there is almost always a Java library alternative.

```java
// ❌ VULNERABLE: System command execution with user input
proc = Runtime.getRuntime().exec(new String[] { "bash", "-c", "ping -c1 " + host });

// ❌ STILL BAD: Validation doesn't eliminate injection risk
if (host.matches("[a-zA-Z0-9.-]+")) {  // Insufficient protection!
    proc = Runtime.getRuntime().exec("ping -c1 " + host);
}

// ✅ CORRECT: Replace command with Java library
import java.net.InetAddress;

private boolean checkConnectivity(String host) {
    try {
        return InetAddress.getByName(host).isReachable(5000);
    } catch (IOException e) {
        logger.error("Connectivity check failed: " + host, e);
        return false;
    }
}
```

## Common Command → Library Replacements

Apply the same pattern for ANY system command:

```java
// File operations → java.nio.file.Files
Files.copy(sourcePath, destPath);           // instead of: cp, copy
Files.move(sourcePath, destPath);           // instead of: mv, move
Files.delete(filePath);                     // instead of: rm, del

// Network → HttpClient, InetAddress, Socket
HttpClient.send(request, handler);          // instead of: curl, wget
InetAddress.getByName(host).isReachable();  // instead of: ping
Socket.connect(address, timeout);           // instead of: telnet, nc

// Archives → java.util.zip, java.util.jar
ZipFile.extractAll(zipPath, destDir);       // instead of: unzip, tar -xf
ZipOutputStream.putNextEntry(entry);        // instead of: zip, tar -cf

// Process management → ProcessHandle
ProcessHandle.current().pid();              // instead of: ps, tasklist
ProcessHandle.allProcesses();               // instead of: ps aux, tasklist
```

## When No Library Exists (Rare)

Only use ProcessBuilder if there is genuinely no Java library for the operation:

```java
// ⚠️ LAST RESORT: Use argument array, never shell invocation
ProcessBuilder pb = new ProcessBuilder(
    "/usr/bin/imagemagick-convert",  // absolute path
    validatedInput,
    validatedOutput
);
Process process = pb.start();  // No shell, no string concatenation
```
