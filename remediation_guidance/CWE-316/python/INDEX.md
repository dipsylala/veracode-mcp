# CWE-316: Cleartext Storage of Sensitive Information in Memory - Python

## LLM Guidance

Storing sensitive data (passwords, API keys, cryptographic keys) in memory as cleartext in Python exposes it to memory dumps, debuggers, and memory disclosure vulnerabilities. Python strings are immutable and persist in memory even after deletion, making secure handling challenging. Use `bytearray` for mutable secrets and explicitly zero them after use.

## Remediation Strategy

- Use mutable types (`bytearray`) instead of immutable strings for sensitive data
- Minimize the lifetime of secrets in memory—clear immediately after use
- Avoid operations that create copies of sensitive data (string concatenation, logging)
- Use secure input methods (`getpass`) and avoid printing/logging credentials
- Consider memory-locking libraries (`mlock`) for highly sensitive applications

## Remediation Steps

- Replace string-based credentials with `bytearray` for passwords and keys
- Implement explicit byte-by-byte zeroing before deallocation
- Use context managers or try-finally blocks to ensure cleanup occurs
- Avoid storing secrets in exception messages or stack traces
- Use `getpass.getpass()` for password input instead of `input()`
- Integrate libraries like `ctypes` with `mlock()` for critical data protection

## Minimal Safe Pattern

```python
import getpass

def authenticate():
    password = bytearray(getpass.getpass("Password: "), 'utf-8')
    try:
        # Use password for authentication
        result = verify_credentials(bytes(password))
        return result
    finally:
        # Clear sensitive data
        for i in range(len(password)):
            password[i] = 0
        del password
```
