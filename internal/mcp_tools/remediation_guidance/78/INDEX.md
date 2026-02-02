# CWE-78: OS Command Injection - LLM Guidance

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands on the host. Never execute OS commands constructed from untrusted input; eliminate shell execution entirely and use safe, parameterized system APIs when OS interaction is unavoidable.

## Key Remediation Principles

- Eliminate shell execution: Replace OS commands with native language APIs whenever possible
- Use parameterized APIs: When OS interaction is required, use safe APIs that prevent shell interpretation
- Validate strictly: If commands are unavoidable, use allowlists for permitted values only
- Avoid shell interpreters: Use direct execution functions that bypass shell parsing
- Apply defence in depth: Combine input validation, safe APIs, and least privilege

## Remediation Steps

- Trace the data path - Identify untrusted data sources, locate system execution sinks, and check for validation gaps between source and sink
- Replace with native APIs - Use language-native libraries for file operations, network requests, and system tasks instead of shell commands
- Use parameterized execution - When OS calls are necessary, use functions that accept argument arrays (e.g., `execve()`, `ProcessBuilder`) to prevent shell interpretation
- Apply strict allowlists - If user input influences commands, validate against a fixed set of permitted values-never use denylists
- Implement least privilege - Run processes with minimal permissions to limit damage from successful exploitation
- Remove special characters - If commands are unavoidable, escape or reject shell metacharacters (`;`, `|`, `&`, `$`, backticks)
