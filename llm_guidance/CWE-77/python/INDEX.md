# CWE-77: Command Injection - Python

## LLM Guidance

Command injection in Python occurs when applications construct system commands using untrusted input, allowing attackers to inject malicious commands through shell metacharacters. The primary defense is using `subprocess.run()` or `subprocess.Popen()` with argument lists (`shell=False`) instead of shell strings, never using `os.system()` or commands with `shell=True` on untrusted input.

## Remediation Strategy

- Use parameterized command execution with argument lists instead of shell string interpolation
- Keep `shell=False` (default) in subprocess calls to prevent shell interpretation
- Validate and sanitize all user input against strict allowlists before use in commands
- Avoid legacy functions like `os.system()`, `os.popen()`, and `commands` module
- Implement least privilege principles for command execution contexts

## Remediation Steps

- Replace `shell=True` with `shell=False` and convert command strings to argument lists
- Refactor `os.system()` calls to `subprocess.run()` with list arguments
- Implement input validation using allowlists for all user-controlled parameters
- Use `shlex.quote()` only as a last resort when shell invocation is unavoidable
- Audit all subprocess and os module usage for untrusted input flow
- Add security testing to verify commands cannot be manipulated

## Minimal Safe Pattern

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
