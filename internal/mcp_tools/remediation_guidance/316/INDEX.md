# CWE-316 - Cleartext Storage of Sensitive Information in Memory

## LLM Guidance

Storing sensitive data (passwords, keys, tokens, PII) in memory as cleartext exposes it to memory dumps, swap files, debuggers, and memory disclosure vulnerabilities. The core fix is to minimize the lifetime of sensitive data in memory, clear it immediately after use, and prevent it from being written to disk.

## Key Principles

- Avoid storing sensitive data in cleartext memory whenever possible
- Minimize the lifetime and exposure of sensitive data in memory
- Explicitly zero/overwrite memory containing secrets after use
- Use secure memory types and APIs designed for sensitive data
- Prevent sensitive data from being swapped to disk or captured in dumps

## Remediation Steps

- Identify where sensitive data enters memory - locate files, line numbers, and code patterns storing passwords, keys, tokens, or PII
- Zero memory immediately after use - overwrite arrays/buffers with zeros (`Arrays.fill(password, '\0')`, `memset_s()`, `SecureZeroMemory()`)
- Use secure types - prefer `SecureString`, `char[]` over `String`, or platform-specific secure memory APIs
- Minimize data lifetime - load sensitive data only when needed, clear it in finally blocks or defer statements
- Prevent swapping - use memory locking APIs (`mlock()`, `VirtualLock()`) for highly sensitive data
- Avoid logging or serializing variables containing secrets
