# CWE-78: OS Command Injection - Python

## LLM Guidance

OS Command Injection occurs when applications incorporate untrusted data into OS commands without proper validation, allowing attackers to execute arbitrary commands on the host system. The primary defense is using Python native libraries (pathlib, requests, etc.) instead of system commands. If system calls are unavoidable, use `subprocess.run()` with argument lists and `shell=False`.

## Key Principles

- Use Python native libraries instead of executing system commands to eliminate vulnerability entirely
- If subprocess is required, use `subprocess.run()` with argument lists and `shell=False` (never `shell=True`)
- Implement allowlist-based input validation as defense in depth, never as sole protection
- Apply least privilege principles and resource limits alongside other defenses
- Never concatenate user input into command strings

## Remediation Steps

- Replace system commands with Python native libraries (pathlib for files, requests for HTTP)
- If subprocess is unavoidable, pass arguments as a list and explicitly set `shell=False`
- Validate all input using strict allowlists of permitted characters
- Drop privileges before command execution and enforce resource limits
- Avoid `os.system()`, `shell=True`, and any string-based command construction

## Safe Pattern

```python
import subprocess

# SAFE: Argument list with shell=False
user_file = get_user_input()
result = subprocess.run(
    ['ls', '-l', user_file],
    shell=False,
    capture_output=True,
    text=True
)

# BETTER: Use Python native library
from pathlib import Path
file_path = Path(user_file)
if file_path.exists():
    stat = file_path.stat()
```
