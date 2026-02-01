# CWE-22: Path Traversal - C\#

## LLM Guidance

Path Traversal occurs when user-supplied input constructs file paths without validation, allowing attackers to use `../` sequences or absolute paths to access files outside intended directories. The core fix is to use indirect reference mapping (map IDs to filenames) or validate paths with `Path.GetFullPath()` and ensure they remain within the allowed base directory.

## Key Principles

- Never directly concatenate user input into file paths
- Use allowlists for filenames, not denylists for patterns
- Canonicalize paths with `Path.GetFullPath()` before validation
- Always verify resolved paths start with the intended base directory
- Prefer indirect references (database IDs mapped to filenames)

## Remediation Steps

- Identify all user inputs that influence file operations (reads, writes, includes)
- Replace direct path construction with safe methods using `Path.Combine()`
- Implement base directory validation after canonicalizing with `Path.GetFullPath()`
- Strip or reject path traversal sequences (`..`, absolute paths) from user input
- Use allowlist validation for permitted filenames or extensions
- Test with payloads - `../`, `..\\`, absolute paths, encoded variants

## Safe Pattern

```csharp
public string GetSafeFilePath(string userInput, string baseDirectory)
{
    // Remove any path traversal attempts
    string fileName = Path.GetFileName(userInput);
    
    // Combine with base directory
    string fullPath = Path.GetFullPath(Path.Combine(baseDirectory, fileName));
    
    // Verify path stays within base directory
    if (!fullPath.StartsWith(Path.GetFullPath(baseDirectory), StringComparison.OrdinalIgnoreCase))
        throw new UnauthorizedAccessException("Invalid path");
    
    return fullPath;
}
```
