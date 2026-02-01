# CWE-73: External Control of File Name or Path - Python

## LLM Guidance

External control of file names or paths occurs when user-supplied input constructs file system paths without proper validation, enabling attackers to access unauthorized files through path traversal. Python's `open()`, `os.path`, and `pathlib` modules provide minimal built-in protection against these attacks. Use `Path.resolve()` with `relative_to()` validation to ensure canonicalized paths (with symlinks resolved) remain within intended directories.

## Remediation Strategy

- Canonicalize all paths using `Path.resolve()` to eliminate symlinks and relative components
- Validate resolved paths stay within the intended base directory using `relative_to()`
- Use allowlists for permitted filenames/extensions, never blocklists for dangerous patterns
- Leverage framework-provided safe functions like Flask's `send_from_directory()` with `safe_join()`
- Never directly concatenate user input into file paths without validation

## Remediation Steps

- Identify user input sources - request parameters (`request.args`, `request.form`), file uploads (`request.FILES`, `filename` attribute), URL path parameters, headers, and API payloads
- Trace data flow to file operations - `open()`, `Path().read_text()`, `os.remove()`, `shutil.copy()`, `send_file()`
- Wrap all file access with `Path.resolve()` followed by `.relative_to(base_dir)` validation in a try-except block
- Implement allowlist checks for filenames using regex patterns or approved extension lists
- Replace custom file serving with `send_from_directory()` (Flask) or equivalent framework functions
- Test with traversal payloads - `../../../etc/passwd`, `..%2F..%2F`, symlink attacks

## Minimal Safe Pattern

```python
from pathlib import Path

BASE_DIR = Path("/var/app/uploads").resolve()

def safe_read_file(user_filename):
    try:
        requested_path = (BASE_DIR / user_filename).resolve()
        requested_path.relative_to(BASE_DIR)  # Raises ValueError if outside
        return requested_path.read_text()
    except (ValueError, FileNotFoundError):
        raise PermissionError("Invalid file path")
```
