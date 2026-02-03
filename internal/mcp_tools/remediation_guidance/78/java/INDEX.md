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

## Safe Pattern

```java
// ❌ VULNERABLE: System command with user input
proc = Runtime.getRuntime().exec(new String[] { "bash", "-c", "ping -c1 " + host });

// ❌ STILL BAD: Validation doesn't fix command injection
if (host.matches("[a-zA-Z0-9.-]+")) {  // Insufficient protection!
    proc = Runtime.getRuntime().exec("ping -c1 " + host);
}

// ✅ CORRECT: Use Java library instead of system command
import java.net.InetAddress;

private boolean isHostReachable(String host) {
    try {
        InetAddress address = InetAddress.getByName(host);
        // Test connectivity with 5 second timeout
        return address.isReachable(5000);
    } catch (UnknownHostException e) {
        logger.error("Unknown host: " + host, e);
        return false;
    } catch (IOException e) {
        logger.error("Network error checking host: " + host, e);
        return false;
    }
}

// Alternative: Socket-based connectivity test
private boolean isPortReachable(String host, int port, int timeoutMs) {
    try (Socket socket = new Socket()) {
        socket.connect(new InetSocketAddress(host, port), timeoutMs);
        return true;
    } catch (IOException e) {
        return false;
    }
}

// ✅ OTHER LIBRARY-BASED SOLUTIONS:

// File copy: Use Files API, not cp command
Path source = Paths.get(validatedPath);
Path dest = Paths.get("/backup/data.txt");
Files.copy(source, dest, StandardCopyOption.REPLACE_EXISTING);

// HTTP request: Use HttpClient, not curl/wget
HttpClient client = HttpClient.newHttpClient();
HttpRequest request = HttpRequest.newBuilder()
    .uri(URI.create(validatedUrl))
    .build();
HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

// Archive extraction: Use ZipFile, not unzip command
try (ZipFile zipFile = new ZipFile(validatedZipPath)) {
    Enumeration<? extends ZipEntry> entries = zipFile.entries();
    while (entries.hasMoreElements()) {
        ZipEntry entry = entries.nextElement();
        Path entryPath = Paths.get("/extract/" + entry.getName());
        Files.copy(zipFile.getInputStream(entry), entryPath);
    }
}

// ⚠️ LAST RESORT: If no library alternative exists
// Use ProcessBuilder with separated arguments (never shell invocation)
ProcessBuilder pb = new ProcessBuilder(
    "/usr/bin/convert",  // absolute path to command
    validatedInputFile,   // argument 1
    validatedOutputFile   // argument 2
);
pb.redirectErrorStream(true);
Process process = pb.start();
```
