# CWE-377: Insecure Temporary File - C#

## LLM Guidance

Insecure temporary file creation occurs when applications create files with predictable names, insecure permissions, or without proper cleanup mechanisms. Use .NET's `Path.GetTempFileName()` for unpredictable filenames, `FileOptions.DeleteOnClose` for automatic cleanup, and set appropriate ACLs to restrict file access to the current user only.

## Key Principles

- Use `Path.GetTempFileName()` instead of constructing custom temporary filenames
- Enable `FileOptions.DeleteOnClose` to ensure automatic cleanup when file handles close
- Apply restrictive ACLs limiting access to the current user and SYSTEM
- Implement deterministic disposal patterns using `using` statements
- Validate temporary directory paths before creating files

## Remediation Steps

- Replace manual filename construction with `Path.GetTempFileName()` for unpredictable names
- Add `FileOptions.DeleteOnClose` flag to `FileStream` constructor for automatic deletion
- Configure `FileSecurity` with ACLs restricting access to current user
- Wrap file operations in `using` statements to ensure cleanup on exceptions
- Use `FileOptions.Encrypted` when handling sensitive data in temporary files
- Validate `Path.GetTempPath()` output points to expected secure location

## Safe Pattern

```csharp
string tempFile = Path.GetTempFileName();
try
{
    using (FileStream fs = new FileStream(
        tempFile,
        FileMode.Create,
        FileAccess.ReadWrite,
        FileShare.None,
        4096,
        FileOptions.DeleteOnClose))
    {
        // Write temporary data securely
        byte[] data = GetDataToWrite();
        fs.Write(data, 0, data.Length);
    } // Automatically deleted on close
}
catch { if (File.Exists(tempFile)) File.Delete(tempFile); }
```
