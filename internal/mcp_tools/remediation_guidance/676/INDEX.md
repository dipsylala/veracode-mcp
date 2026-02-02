# CWE-676: Use of Potentially Dangerous Function

## LLM Guidance

Dangerous functions (strcpy, gets, system, eval, exec) lack bounds checking, enable code execution, or have inherent security flaws. Using them creates buffer overflows, command injection, code injection, and other vulnerabilities that modern alternatives prevent.

## Key Principles

- Replace unsafe string functions with bounds-checked alternatives: strcpy → strncpy/snprintf/strlcpy, gets → fgets, sprintf → snprintf
- Replace command execution with parameterized alternatives: system/os.system → subprocess with argument arrays, shell=False
- Use secure crypto/random functions: rand() → RAND_bytes, /dev/random, secrets module
- Ban dangerous functions in security-sensitive code through linting rules, static analysis, and code review policies
- Add validation layers if dangerous functions cannot be replaced: strict input validation, bounds checking, allowlisting

## Remediation Steps

- Identify dangerous calls - Grep codebase for strcpy, gets, system, eval, exec, sprintf, strcat, scanf and similar functions
- Analyze data flow - Trace if user input or untrusted data reaches the dangerous function
- Replace systematically - Use strncpy + null terminator, snprintf, fgets, subprocess with list arguments, parameterized queries
- Validate inputs - Where replacement impossible, add strict length checks, character allowlisting, bounds verification before calling function
- Test with exploits - Try oversized inputs, command injection payloads (`;rm -rf /`), format strings to verify fixes block attacks
- Enforce prevention - Add linting rules (bandit, semgrep) to catch dangerous functions in CI/CD pipeline
