# CWE-183: Permissive List of Allowed Inputs - Python

## LLM Guidance

CWE-183 occurs when input validation uses overly permissive patterns that fail to match the entire input string, allowing attackers to inject malicious content before or after valid data. Use fully anchored regex patterns with `^` and `$` or `re.fullmatch()`, validate complex inputs with specialized libraries like `pathlib` and `ipaddress`, and enforce strict length limits.

## Key Principles

- Use `re.fullmatch()` or anchor patterns with `^...$` to ensure complete string matching
- Leverage Python's specialized validation libraries (`pathlib`, `ipaddress`, `urllib.parse`) instead of custom regex
- Validate against strict allowlists of permitted values using sets or enums
- Enforce input length limits before pattern matching
- Fail closed on validation errors with explicit rejection

## Remediation Steps

- Replace `re.search()` or `re.match()` with `re.fullmatch()` for complete validation
- Add anchors `^` and `$` to all existing regex patterns if not using `fullmatch()`
- Use `pathlib.Path.resolve()` to canonicalize and validate file paths against allowed directories
- Apply the `ipaddress` module for IP address validation instead of regex
- Check input length with `len()` before validation to prevent DoS
- Use sets or frozensets for exact string matching against allowlists

## Safe Pattern

```python
import re
from pathlib import Path

ALLOWED_USERNAMES = frozenset(['alice', 'bob', 'charlie'])
USERNAME_PATTERN = re.compile(r'^[a-z]{3,20}$')

def validate_username(username: str) -> bool:
    if len(username) > 20:  # Length check first
        return False
    return USERNAME_PATTERN.fullmatch(username) is not None

def validate_path(user_path: str, base_dir: Path) -> bool:
    try:
        resolved = Path(user_path).resolve()
        return resolved.is_relative_to(base_dir)
    except (ValueError, OSError):
        return False
```
