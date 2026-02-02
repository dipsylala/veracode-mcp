# CWE-377: Insecure Temporary File - Python

## LLM Guidance

Insecure temporary file creation occurs when applications create files with predictable names, insecure permissions, or in shared directories without proper protection. Python's `tempfile` module provides secure alternatives that generate unpredictable names with restricted permissions and automatic cleanup. Always use `tempfile.NamedTemporaryFile()` or `tempfile.mkstemp()` instead of manually creating files in `/tmp` or similar directories.

## Key Principles

- Use Python's `tempfile` module exclusively for temporary file operations
- Never hardcode temporary file names or use predictable patterns
- Ensure file permissions are restrictive (mode 0o600) to prevent unauthorized access
- Implement automatic cleanup using context managers or `delete=True` parameter
- Avoid creating temporary files in world-writable directories like `/tmp` without proper protections

## Remediation Steps

- Replace manual file creation with `tempfile.NamedTemporaryFile()` or `tempfile.mkstemp()`
- Use context managers (`with` statements) to ensure automatic cleanup
- Set `delete=True` for auto-removal or explicitly handle cleanup in exception handlers
- Verify file permissions are restrictive (default 0o600 is secure)
- For temporary directories, use `tempfile.TemporaryDirectory()` with context managers
- Avoid passing temporary file paths to untrusted code or processes

## Safe Pattern

```python
import tempfile

# Secure temporary file with auto-cleanup
with tempfile.NamedTemporaryFile(mode='w+', delete=True, suffix='.txt') as tmp:
    tmp.write("sensitive data")
    tmp.flush()
    # File automatically deleted when context exits

# Alternative: Manual control with secure creation
fd, path = tempfile.mkstemp(suffix='.txt')
try:
    with os.fdopen(fd, 'w') as tmp:
        tmp.write("sensitive data")
finally:
    os.unlink(path)  # Ensure cleanup
```
