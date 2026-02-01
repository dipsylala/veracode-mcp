# CWE-77: Command Injection - C# / .NET

## LLM Guidance

Command injection in .NET occurs when applications construct system commands using untrusted input through Process class and related APIs, allowing attackers to execute arbitrary commands. The primary defense is avoiding shell execution by setting `UseShellExecute = false` and passing arguments as arrays instead of concatenating strings. Always validate user input against strict allowlists before using in command construction.

## Remediation Strategy

- Avoid shell execution entirely by setting `UseShellExecute = false`
- Pass arguments as arrays using `ArgumentList` collection, never concatenate strings
- Validate all user input against strict allowlists (e.g., alphanumeric patterns)
- Use absolute paths for executable files
- Apply least privilege principles to process execution

## Remediation Steps

- Set `ProcessStartInfo.UseShellExecute = false` to prevent shell interpretation
- Use `ProcessStartInfo.ArgumentList.Add()` to pass individual arguments safely
- Validate input against allowlists before use
- Specify full executable paths to prevent PATH injection
- Implement input sanitization for unavoidable dynamic values
- Log and monitor all system command executions

## Minimal Safe Pattern

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
