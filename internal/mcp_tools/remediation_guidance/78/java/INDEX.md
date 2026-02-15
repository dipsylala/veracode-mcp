# CWE-78: OS Command Injection - Java

## LLM Guidance

**STOP. Before doing anything else: DO NOT execute system processes.**

OS Command Injection occurs when OS commands are executed with untrusted data. The ONLY correct remediation is to **ELIMINATE ALL calls to Runtime.exec(), ProcessBuilder, and process execution entirely**.

## ⛔ FORBIDDEN APPROACHES

These are **NEVER acceptable remediations**:

- ❌ Input validation/sanitization
- ❌ Input filtering or allowlists
- ❌ Using ProcessBuilder with "safe" arguments
- ❌ Escaping special characters
- ❌ Any approach that still executes system processes

**Why?** Because Java has native library alternatives for virtually every operation. Process execution is almost never necessary.

## ✅ REQUIRED APPROACH: Replace with Java Libraries

**Your task:** Find the Java library that replaces the system command, then delete the process execution code entirely.

## Common Command → Library Replacements

**Use this table to find the replacement.** If the code executes ANY of these commands, replace them with the Java library shown:

| System Command | Java Library Alternative | Method |
| ---------------- | ------------------------- | -------- |
| `ping` | `java.net.InetAddress` | `InetAddress.getByName(host).isReachable(timeout)` |
| `curl`, `wget` | `java.net.http.HttpClient` | `HttpClient.send(request, handler)` |
| `cp`, `copy` | `java.nio.file.Files` | `Files.copy(source, dest)` |
| `mv`, `move` | `java.nio.file.Files` | `Files.move(source, dest)` |
| `rm`, `del` | `java.nio.file.Files` | `Files.delete(path)` |
| `mkdir` | `java.nio.file.Files` | `Files.createDirectory(path)` |
| `unzip`, `tar -xf` | `java.util.zip.ZipFile` | `ZipFile.extractAll()` |
| `zip`, `tar -cf` | `java.util.zip.ZipOutputStream` | `ZipOutputStream.putNextEntry()` |
| `ps`, `tasklist` | `java.lang.ProcessHandle` | `ProcessHandle.allProcesses()` |
| `cat`, `type` | `java.nio.file.Files` | `Files.readString(path)` |
| `grep`, `findstr` | Java String/Stream API | `str.contains()`, `Pattern.matcher()` |

**If the command is not in this table:** Research the Java library that provides the same functionality. There is almost certainly one.

## Example: Replacing ping with InetAddress

```java
// ❌ WRONG: Executing ping command
proc = Runtime.getRuntime().exec(new String[] { "bash", "-c", "ping -c1 " + host });

// ❌ STILL WRONG: Adding validation doesn't fix the root problem
if (host.matches("[a-zA-Z0-9.-]+")) {
    proc = Runtime.getRuntime().exec("ping -c1 " + host);  // Still executing a process!
}

// ❌ STILL WRONG: Using ProcessBuilder is still executing a process
ProcessBuilder pb = new ProcessBuilder("ping", "-c1", host);  // NO! Don't execute processes!

// ✅ CORRECT: No process execution at all - use Java library
import java.net.InetAddress;

private boolean isHostReachable(String host) {
    try {
        return InetAddress.getByName(host).isReachable(5000);  // Pure Java, no subprocess
    } catch (IOException e) {
        logger.error("Connectivity check failed for: " + host, e);
        return false;
    }
}
```

**Notice:** The correct solution deletes all Runtime.exec() and ProcessBuilder code completely.

## More Examples: File Operations

```java
// ❌ WRONG: Executing cp command
Runtime.getRuntime().exec("cp " + source + " " + dest);

// ✅ CORRECT: Use Files API
import java.nio.file.Files;
import java.nio.file.Paths;

Files.copy(Paths.get(source), Paths.get(dest), StandardCopyOption.REPLACE_EXISTING);
```

## More Examples: HTTP Requests

```java
// ❌ WRONG: Executing curl command  
Runtime.getRuntime().exec("curl " + url);

// ✅ CORRECT: Use HttpClient
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

HttpClient client = HttpClient.newHttpClient();
HttpRequest request = HttpRequest.newBuilder()
    .uri(URI.create(url))
    .build();
String response = client.send(request, HttpResponse.BodyHandlers.ofString()).body();
```

---

## ⚠️ EXTREMELY RARE EXCEPTION: No Java Library Exists

**STOP.** Before reading this section, ask yourself: "Have I thoroughly researched Java libraries, Apache Commons, and other standard libraries?"

**If you're executing a standard command like ping, curl, cp, mv, tar, zip, ps, cat, grep - STOP. Go back to the replacement table above. These all have Java library alternatives.**

Only proceed if:

1. The operation is highly specialized (e.g., hardware-specific tool, vendor CLI)
2. You've confirmed no Java library exists
3. You cannot install a Java library that provides this functionality

In this extremely rare case, use argument arrays (never shell invocation):

```java
// ⚠️ LAST RESORT ONLY - Almost always wrong to use this
ProcessBuilder pb = new ProcessBuilder(
    "/usr/bin/specialized-tool",  // Absolute path to avoid PATH manipulation
    validatedArg1,                 // Each argument separately
    validatedArg2
);
pb.redirectErrorStream(true);
Process process = pb.start();

// Note: NO shell (no "sh", "bash", "cmd"), NO string concatenation
```

**Re-check:** Are you absolutely certain there's no Java library? ProcessBuilder should appear in less than 1% of CWE-78 remediations.
