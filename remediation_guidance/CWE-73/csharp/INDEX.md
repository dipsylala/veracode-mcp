# CWE-73: External Control of File Name or Path - C\#

## LLM Guidance

External control of file names or paths occurs when user-supplied input constructs file system paths without validation in C#/.NET applications. The .NET `System.IO` namespace provides minimal built-in protection against path traversal attacks. Use `Path.GetFullPath()` with `StartsWith()` validation to ensure resolved paths remain within intended base directories.

## Key Principles

- Always validate user-supplied file paths against an allowed base directory
- Canonicalize paths using `Path.GetFullPath()` to resolve traversal sequences (`../`, `..\\`)
- Use allowlists for file extensions and names when possible
- Never trust `IFormFile.FileName` or any user-controlled path input directly
- Implement defense-in-depth with filesystem permissions

## Remediation Steps

- Identify sources - Find untrusted input from `Request.Query`, `Request.Form`, `IFormFile.FileName`, route parameters, headers, or deserialized objects
- Trace to sinks - Locate file operations using `File.*()` methods, `FileStream`, `StreamReader/Writer`, `FileInfo`, or `Path.Combine()`
- Define base directory - Establish an allowed root directory for file operations
- Canonicalize paths - Use `Path.GetFullPath()` to resolve the full absolute path
- Validate containment - Verify the resolved path starts with the allowed base directory using `StartsWith()`
- Implement allowlists - Validate file extensions and names against approved patterns

## Safe Pattern

```csharp
string baseDirectory = Path.GetFullPath("/safe/files/");
string userInput = Request.Query["filename"];

// Combine and canonicalize
string requestedPath = Path.Combine(baseDirectory, userInput);
string fullPath = Path.GetFullPath(requestedPath);

// Validate path stays within base directory
if (!fullPath.StartsWith(baseDirectory, StringComparison.OrdinalIgnoreCase))
{
    throw new UnauthorizedAccessException("Invalid file path");
}

// Safe to proceed
byte[] content = File.ReadAllBytes(fullPath);
```
