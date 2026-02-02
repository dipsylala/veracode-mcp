# CWE-252: Unchecked Return Value - LLM Guidance

## LLM Guidance

Unchecked return values occur when code ignores error indicators from security-critical functions (setuid, chroot, malloc, read, write), assuming operations succeeded when they may have failed. This leads to continued execution with wrong privileges, uninitialized memory, or invalid state. Always check return values and handle failures securely.

## Key Principles

- Check all return values from security-critical functions (privilege changes, memory allocation, file I/O, cryptographic operations)
- Fail securely when operations don't succeed-terminate, log errors, or revert to safe state
- Never assume success for functions that can fail (setuid, malloc, chroot, read/write)
- Validate both return codes and side effects (e.g., verify privileges actually changed)
- Use compiler warnings and static analysis to detect unchecked returns

## Remediation Steps

- Locate unchecked calls - Find setuid/setgid, malloc/calloc, file operations (open/read/write), chroot/chdir, and crypto functions without return value checks
- Add explicit checks - Wrap critical calls with `if (func() != 0)` or `if (ptr == NULL)` checks
- Handle failures securely - Log the error, clean up resources, and terminate or return error codes-never continue with invalid state
- Verify side effects - After privilege changes, confirm actual UID/GID with getuid/getgid
- Use defensive patterns - Initialize pointers to NULL, check errno after failures, validate file descriptors before use
- Enable compiler warnings - Use `-Wall -Wextra` (GCC/Clang) or `/W4` (MSVC) and treat unused return value warnings as errors
