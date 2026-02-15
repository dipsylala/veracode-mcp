# CWE-78: OS Command Injection - Java

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands on the host. In Java, eliminate Runtime.exec() and ProcessBuilder calls entirely by using native Java libraries (java.nio.file.Files, java.net.http.HttpClient, java.util.zip, etc.) for file operations, HTTP requests, and archive handling.

## Key Principles

- Replace all Runtime.exec() and ProcessBuilder calls with Java standard library alternatives
- Use java.nio.file.Files for file operations (copy, move, delete) instead of system commands
- Use java.net.http.HttpClient or HttpURLConnection for HTTP requests instead of curl/wget
- Use java.util.zip for archive operations instead of tar/zip commands
- Never concatenate user input into command strings
- Only use ProcessBuilder as a last resort with validated argument lists (no shell invocation)

## Remediation Steps

- Locate command execution - Identify all Runtime.exec() and ProcessBuilder instances
- Determine the operation's purpose - Understand what the command is trying to accomplish
- Find the Java library alternative - Use Files API for file ops, HttpClient for HTTP, InetAddress for network checks
- Replace process execution - Delete Runtime.exec()/ProcessBuilder code and use the appropriate Java library
- For unavoidable commands - Use ProcessBuilder with separate arguments (never shell), validate all inputs
- Test thoroughly - Verify the Java library replacement provides the same functionality

## Safe Pattern

```java
// UNSAFE: Executing ping command
Runtime.getRuntime().exec("ping -c 1 " + host);

// UNSAFE: Even with ProcessBuilder
ProcessBuilder pb = new ProcessBuilder("ping", "-c1", host);
Process proc = pb.start();

// SAFE: Use InetAddress for reachability check
import java.net.InetAddress;

private boolean isHostReachable(String host, int timeout) throws IOException {
    return InetAddress.getByName(host).isReachable(timeout);
}

// SAFE: File copy with Files API
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;

Files.copy(Paths.get(source), Paths.get(dest), StandardCopyOption.REPLACE_EXISTING);

// SAFE: HTTP request with HttpClient
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

HttpClient client = HttpClient.newHttpClient();
HttpRequest request = HttpRequest.newBuilder().uri(URI.create(url)).build();
HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
```
