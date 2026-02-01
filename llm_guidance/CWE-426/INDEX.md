# CWE-426: Untrusted Search Path

## LLM Guidance

Untrusted Search Path vulnerabilities occur when applications search for resources (executables, libraries, DLLs) in directories that attackers can control, allowing malicious file injection. The core fix is to eliminate reliance on untrusted search paths and use absolute paths for all resource loading.

## Key Principles

- Always use absolute paths for executables and libraries
- Never rely on PATH, LD_LIBRARY_PATH, or similar environment variables
- Restrict search paths to trusted system directories only
- Verify integrity of loaded components when possible
- Remove current/relative directories from search order

## Remediation Steps

- Review flaw details to identify where resources are loaded from untrusted paths
- Identify resource types being loaded - executables, libraries, DLLs, shared objects, or config files
- Check how resources are located - relative paths, PATH variable, LD_LIBRARY_PATH, or system search order
- Determine if attackers can control any directories in the search path
- Replace all relative paths with absolute paths (e.g., `/usr/local/bin/tool` instead of `tool`)
- Remove dependency on environment variables for critical resource loading
