# CWE-78: OS Command Injection - C\#

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands.

**CRITICAL: DO NOT suggest input validation as the primary fix.** The correct remediation is to **ELIMINATE Process.Start() calls** by using .NET framework classes. Input validation alone is insufficient.

**Primary Defence:** Replace ALL system commands with .NET framework classes. Common replacements:

- **file operations → File, Directory classes** instead of system commands
- **HTTP requests → HttpClient** instead of curl/wget commands
- **archive operations → ZipFile, ZipArchive** instead of tar/unzip commands

Only if no .NET class exists, use `ProcessStartInfo.ArgumentList` with `UseShellExecute = false`.

## Key Principles

- **ALWAYS prefer .NET framework classes over Process.Start()** - eliminates injection risk entirely
- For file operations, use **File** and **Directory** classes instead of system commands
- For archives, use **System.IO.Compression.ZipFile** instead of system commands
- **Input validation is NOT a remediation** - it's only defense-in-depth after eliminating Process.Start()
- If Process is unavoidable, use `ProcessStartInfo.ArgumentList` (.NET 6+) with `UseShellExecute = false`
- Never use `UseShellExecute = true` or string-based command construction

## Remediation Steps

- Identify all Process.Start(), Process.Run, and ProcessStartInfo usages
- **First, determine the .NET framework class alternative** - do not default to "secure" Process patterns
- Replace system commands with .NET framework classes
- For any truly unavoidable commands, refactor to use `ProcessStartInfo.ArgumentList` with `UseShellExecute = false`
- Remove all `UseShellExecute = true` usage and cmd.exe/bash invocations
- Apply strict allowlist validation only as defense-in-depth

## Remediation Pattern

**The transformation pattern:** System Command → .NET Framework Class

This applies to ALL system commands - if you find Process.Start() executing a command, there is almost always a .NET class alternative.

```csharp
using System.Net.NetworkInformation;
using System.Net.Sockets;

// ❌ VULNERABLE: System command execution with user input
var startInfo = new ProcessStartInfo
{
    FileName = "cmd.exe",
    Arguments = $"/c ping {host}"
};
Process.Start(startInfo);

// ❌ STILL BAD: ArgumentList with validation doesn't eliminate injection risk
if (Regex.IsMatch(host, @"^[a-zA-Z0-9.-]+$"))  // Insufficient!
{
    var start = new ProcessStartInfo("ping");
    start.ArgumentList.Add(host);
    Process.Start(start);
}

// ✅ CORRECT: Replace command with .NET class
public async Task<bool> CheckConnectivity(string host, int timeout = 5000)
{
    try {
        using var ping = new Ping();
        var reply = await ping.SendPingAsync(host, timeout);
        return reply.Status == IPStatus.Success;
    } catch (PingException) {
        return false;
    }
}
```

## Common Command → Class Replacements

Apply the same pattern for ANY system command:

```csharp
// File operations → File, Directory classes
File.Copy(source, dest);                // instead of: copy, xcopy
Directory.Move(source, dest);           // instead of: move, mv
File.Delete(file);                      // instead of: del, rm

// Network → HttpClient, Ping, TcpClient
await httpClient.GetAsync(url);         // instead of: curl, wget
await ping.SendPingAsync(host);         // instead of: ping
await tcpClient.ConnectAsync(host, port); // instead of: telnet

// Archives → ZipFile, ZipArchive
ZipFile.ExtractToDirectory(zip, dest);  // instead of: unzip, tar -xf  
ZipFile.CreateFromDirectory(src, zip);  // instead of: zip, tar -cf

// Process info → Process class
Process.GetProcesses();                 // instead of: tasklist, ps
Process.GetCurrentProcess().Id;         // instead of: echo $$
```

## When No Class Exists (Rare)

Only use Process if there is genuinely no .NET class for the operation:

```csharp
// ⚠️ LAST RESORT: Use ArgumentList with UseShellExecute = false
var startInfo = new ProcessStartInfo
{
    FileName = @"C:\tools\imagemagick.exe",
    UseShellExecute = false  // NEVER use true
};
startInfo.ArgumentList.Add(validatedInput);
startInfo.ArgumentList.Add(validatedOutput);
using var process = Process.Start(startInfo);
```
