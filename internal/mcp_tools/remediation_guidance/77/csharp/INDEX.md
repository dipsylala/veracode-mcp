# CWE-77: Command Injection - C# / .NET

## LLM Guidance

Command injection in .NET occurs when applications construct system commands using untrusted input through Process class and related APIs, allowing attackers to execute arbitrary commands.

**Primary Defence:** Use .NET native APIs (File, Directory, HttpClient, etc.) instead of executing system commands to eliminate the vulnerability entirely. If system commands are unavoidable, set `UseShellExecute = false` and use `ProcessStartInfo.ArgumentList` with argument arrays.

## Key Principles

- **BEST:** Use .NET native APIs (File, Directory, HttpClient) instead of Process.Start() to eliminate command injection risk
- **If commands unavoidable:** Set `UseShellExecute = false` and use `ProcessStartInfo.ArgumentList` with argument arrays
- Never concatenate strings or use shell execution with untrusted input
- Validate all user input against strict allowlists as defence-in-depth
- Use absolute paths for executable files to prevent PATH injection
- Apply least privilege principles to process execution

## Remediation Steps

- **Replace system commands with .NET native APIs** (File.Copy, Directory.Delete, HttpClient, etc.) to eliminate vulnerability
- If commands are unavoidable, set `ProcessStartInfo.UseShellExecute = false` to prevent shell interpretation
- Use `ProcessStartInfo.ArgumentList.Add()` (.NET 6+) to pass individual arguments safely
- Validate input against allowlists as additional defence
- Specify full executable paths to prevent PATH injection
- Log and monitor all system command executions

## Safe Pattern

```csharp
var startInfo = new ProcessStartInfo
{
    FileName = "/usr/bin/tool",
    UseShellExecute = false,
    RedirectStandardOutput = true
};

// Add arguments individually - never concatenate
startInfo.ArgumentList.Add("--option");
startInfo.ArgumentList.Add(validatedInput);

using var process = Process.Start(startInfo);
```
