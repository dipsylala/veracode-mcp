# CWE-77: Command Injection - Python

## LLM Guidance

Command injection in Python occurs when applications construct system commands using untrusted input, allowing attackers to inject malicious commands through shell metacharacters.

**Primary Defence:** Use Python native libraries (pathlib, requests, shutil, etc.) instead of executing system commands to eliminate the vulnerability entirely. If system commands are unavoidable, use `subprocess.run()` with argument lists and `shell=False`.

## Key Principles

- **BEST:** Use Python native libraries (pathlib for files, requests for HTTP, shutil for file operations) instead of system commands
- **If commands unavoidable:** Use `subprocess.run()` with argument lists and `shell=False` (never `shell=True`)
- Avoid legacy functions like `os.system()`, `os.popen()`, and `commands` module entirely
- Validate and sanitize all user input against strict allowlists as defence-in-depth
- Implement least privilege principles for command execution contexts

## Remediation Steps

- **Replace system commands with Python native libraries** (pathlib, requests, shutil) to eliminate vulnerability
- If commands are unavoidable, replace `shell=True` with `shell=False` and convert command strings to argument lists
- Refactor `os.system()` calls to `subprocess.run()` with list arguments
- Implement input validation using allowlists for all user-controlled parameters as additional defence
- Use `shlex.quote()` only as a last resort when shell invocation is absolutely unavoidable
- Audit all subprocess and os module usage for untrusted input flow
- Add security testing to verify commands cannot be manipulated

## Safe Pattern

```python
import subprocess

# SAFE: Argument list with shell=False (default)
user_file = get_user_input()
result = subprocess.run(
    ['ls', '-l', user_file],
    capture_output=True,
    text=True,
    check=True
)

# For dynamic arguments, validate first
allowed_options = {'-l', '-a', '-h'}
if user_option in allowed_options:
    subprocess.run(['ls', user_option, user_file])
```
