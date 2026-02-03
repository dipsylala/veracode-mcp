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

## Safe Pattern

```csharp
using System.Diagnostics;
using System.Net;
using System.Net.NetworkInformation;
using System.Net.Sockets;

// ❌ VULNERABLE: System command with user input
var startInfo = new ProcessStartInfo
{
    FileName = "cmd.exe",
    Arguments = $"/c ping {host}",
    UseShellExecute = false
};
Process.Start(startInfo);

// ❌ STILL BAD: ArgumentList with validation doesn't fix the root problem
if (Regex.IsMatch(host, @"^[a-zA-Z0-9.-]+$"))  // Insufficient protection!
{
    var start = new ProcessStartInfo("ping");
    start.ArgumentList.Add(host);
    Process.Start(start);
}

// ✅ CORRECT: Use .NET Ping class instead of ping command
using System.Net.NetworkInformation;

public async Task<bool> IsHostReachableAsync(string host, int timeout = 5000)
{
    try
    {
        using var ping = new Ping();
        var reply = await ping.SendPingAsync(host, timeout);
        return reply.Status == IPStatus.Success;
    }
    catch (PingException ex)
    {
        Console.WriteLine($"Ping failed: {ex.Message}");
        return false;
    }
}

// Alternative: TcpClient for port connectivity test
public async Task<bool> IsPortReachableAsync(string host, int port, int timeoutMs = 5000)
{
    try
    {
        using var client = new TcpClient();
        using var cts = new CancellationTokenSource(timeoutMs);
        await client.ConnectAsync(host, port, cts.Token);
        return true;
    }
    catch (Exception ex)
    {
        Console.WriteLine($"Connection failed: {ex.Message}");
        return false;
    }
}

// ✅ OTHER FRAMEWORK-BASED SOLUTIONS:

// File operations: Use File/Directory classes, not system commands
File.Copy(userSourcePath, @"C:\backup\data.txt", overwrite: true);
Directory.Move(oldPath, newPath);
File.Delete(filePath);

// HTTP requests: Use HttpClient, not curl/wget commands
using var httpClient = new HttpClient { Timeout = TimeSpan.FromSeconds(10) };
var response = await httpClient.GetAsync(validatedUrl);
var content = await response.Content.ReadAsStringAsync();

// Archive extraction: Use ZipFile, not unzip command
using System.IO.Compression;

ZipFile.ExtractToDirectory(validatedZipPath, @"C:\extract\");

// Archive creation: Use ZipArchive
using (var archive = ZipFile.Open(@"C:\backup\archive.zip", ZipArchiveMode.Create))
{
    archive.CreateEntryFromFile(validatedFilePath, Path.GetFileName(validatedFilePath));
}

// ⚠️ LAST RESORT: If no .NET framework class exists
// Use ProcessStartInfo.ArgumentList with UseShellExecute = false
var startInfo = new ProcessStartInfo
{
    FileName = @"C:\tools\convert.exe",  // Absolute path
    UseShellExecute = false,              // NEVER use true
    RedirectStandardOutput = true,
    RedirectStandardError = true
};
startInfo.ArgumentList.Add(validatedInput);
startInfo.ArgumentList.Add(validatedOutput);

using var process = Process.Start(startInfo);
await process.WaitForExitAsync();
```
