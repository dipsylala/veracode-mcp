# CWE-78: OS Command Injection - C\#

## LLM Guidance

**STOP. Before doing anything else: DO NOT execute system processes.**

OS Command Injection occurs when OS commands are executed with untrusted data. The ONLY correct remediation is to **ELIMINATE ALL calls to Process.Start(), ProcessStartInfo, and process execution entirely**.

## ⛔ FORBIDDEN APPROACHES

These are **NEVER acceptable remediations**:

- ❌ Input validation/sanitization
- ❌ Input filtering or allowlists
- ❌ Using `ProcessStartInfo.ArgumentList` with validated inputs
- ❌ Using `UseShellExecute = false`
- ❌ Escaping shell characters
- ❌ Any approach that still executes system processes

**Why?** Because .NET has framework class alternatives for virtually every operation. Process execution is almost never necessary.

## ✅ REQUIRED APPROACH: Replace with .NET Framework Classes

**Your task:** Find the .NET class that replaces the system command, then delete the process execution code entirely.

## Common Command → Class Replacements

**Use this table to find the replacement.** If the code executes ANY of these commands, replace them with the .NET class shown:

| System Command | .NET Class Alternative | Method |
| ---------------- | ------------------------- | -------- |
| `ping` | `System.Net.NetworkInformation.Ping` | `ping.SendPingAsync(host)` |
| `curl`, `wget` | `System.Net.Http.HttpClient` | `await httpClient.GetAsync(url)` |
| `copy`, `xcopy` | `System.IO.File` | `File.Copy(source, dest)` |
| `move`, `mv` | `System.IO.Directory` | `Directory.Move(source, dest)` |
| `del`, `rm` | `System.IO.File` | `File.Delete(file)` |
| `mkdir` | `System.IO.Directory` | `Directory.CreateDirectory(path)` |
| `unzip`, `tar -xf` | `System.IO.Compression.ZipFile` | `ZipFile.ExtractToDirectory(zip, dest)` |
| `zip`, `tar -cf` | `System.IO.Compression.ZipFile` | `ZipFile.CreateFromDirectory(src, zip)` |
| `tasklist`, `ps` | `System.Diagnostics.Process` | `Process.GetProcesses()` |
| `type`, `cat` | `System.IO.File` | `File.ReadAllText(file)` |
| `findstr`, `grep` | .NET String/LINQ | `str.Contains()`, `Regex.Match()` |

**If the command is not in this table:** Research the .NET class that provides the same functionality. There is almost certainly one.

## Example: Replacing ping with Ping class

```csharp
// ❌ WRONG: Executing ping command
var startInfo = new ProcessStartInfo 
{
    FileName = "cmd.exe",
    Arguments = $"/c ping {host}"
};
Process.Start(startInfo);

// ❌ STILL WRONG: Adding validation doesn't fix the root problem
if (Regex.IsMatch(host, @"^[a-zA-Z0-9.-]+$")) 
{
    var start = new ProcessStartInfo("ping");
    start.ArgumentList.Add(host);  // Still executing a process!
    Process.Start(start);
}

// ❌ STILL WRONG: Using ArgumentList with UseShellExecute=false is still executing a process
var psi = new ProcessStartInfo("ping") { UseShellExecute = false };
psi.ArgumentList.Add(host);  // NO! Don't execute processes!

// ✅ CORRECT: No process execution at all - use .NET class
using System.Net.NetworkInformation;

public async Task<bool> IsHostReachable(string host, int timeout = 5000)
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
```

**Notice:** The correct solution deletes all Process.Start() and ProcessStartInfo code completely.

## More Examples: File Operations

```csharp
// ❌ WRONG: Executing copy command
Process.Start("cmd.exe", $"/c copy {source} {dest}");

// ✅ CORRECT: Use File class
using System.IO;
File.Copy(source, dest, overwrite: true);
```

## More Examples: HTTP Requests

```csharp
// ❌ WRONG: Executing curl command
Process.Start("curl", url);

// ✅ CORRECT: Use HttpClient
using System.Net.Http;
using var httpClient = new HttpClient();
var response = await httpClient.GetAsync(url);
var content = await response.Content.ReadAsStringAsync();
```

---

## ⚠️ EXTREMELY RARE EXCEPTION: No .NET Class Exists

**STOP.** Before reading this section, ask yourself: "Have I thoroughly researched System.IO, System.Net, System.Net.Http, and System.IO.Compression namespaces?"

**If you're executing a standard command like ping, curl, copy, move, del, zip, tasklist, type, findstr - STOP. Go back to the replacement table above. These all have .NET class alternatives.**

Only proceed if:

1. The operation is highly specialized (e.g., hardware-specific tool, vendor CLI)
2. You've confirmed no .NET class exists
3. You cannot install a NuGet package that provides this functionality

In this extremely rare case, use ArgumentList with UseShellExecute = false:

```csharp
// ⚠️ LAST RESORT ONLY - Almost always wrong to use this
var startInfo = new ProcessStartInfo
{
    FileName = @"C:\tools\specialized-tool.exe",  // Absolute path
    UseShellExecute = false,  // NEVER use true
    RedirectStandardOutput = true
};
startInfo.ArgumentList.Add(validated_arg1);
startInfo.ArgumentList.Add(validated_arg2);

using var process = Process.Start(startInfo);
await process.WaitForExitAsync();

// Note: NO cmd.exe, NO bash, NO UseShellExecute = true
```

**Re-check:** Are you absolutely certain there's no .NET class? Process.Start() should appear in less than 1% of CWE-78 remediations.
