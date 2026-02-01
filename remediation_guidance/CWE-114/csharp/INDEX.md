# CWE-114: Process Control - C\#

## LLM Guidance

In C#, CWE-114 vulnerabilities occur when loading DLLs or executing processes without proper validation. Attackers exploit weak library loading through DLL hijacking, DllImport path manipulation, and Process.Start() command injection. .NET applications are vulnerable to both native DLL loading (P/Invoke) and managed assembly loading.

**Primary Defence:** Use `SetDllDirectory()` and `LoadLibraryEx` with `LOAD_LIBRARY_SEARCH_SYSTEM32` flag, validate all paths against allowlists, disable current directory DLL search, and use full absolute paths for Process.Start().

## Remediation Strategy

- Restrict DLL search paths using `SetDllDirectory("")` to remove current directory from search order
- Validate all file paths against strict allowlists before loading assemblies or executing processes
- Use absolute paths with known-safe directories (e.g., System32, application directory)
- Apply least privilege principles when spawning child processes
- Never concatenate user input directly into process arguments or DLL paths

## Remediation Steps

- Call `SetDllDirectory("")` at application startup to disable current directory DLL loading
- Use `LoadLibraryEx` with `LOAD_LIBRARY_SEARCH_SYSTEM32` flag for system DLLs
- Validate DLL/executable paths against allowlist before P/Invoke or Process.Start()
- Use ProcessStartInfo with `UseShellExecute = false` and escape arguments properly
- Sign assemblies and enable strong name verification for managed code
- Implement file integrity checks (digital signatures) before loading external DLLs

## Minimal Safe Pattern

```csharp
// Safe process execution with validated path
public void ExecuteProcess(string userInput) {
    var allowedCommands = new[] { "notepad.exe", "calc.exe" };
    var command = Path.GetFileName(userInput);
    
    if (!allowedCommands.Contains(command))
        throw new SecurityException("Invalid command");
    
    var psi = new ProcessStartInfo {
        FileName = Path.Combine(Environment.SystemDirectory, command),
        UseShellExecute = false,
        CreateNoWindow = true
    };
    Process.Start(psi);
}
```
