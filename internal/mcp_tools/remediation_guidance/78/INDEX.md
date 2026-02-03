# CWE-78: OS Command Injection - LLM Guidance

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands on the host.

**CRITICAL: DO NOT suggest input validation/sanitization as the primary remediation.** Validation is insufficient and error-prone for command injection. The correct fix is to **ELIMINATE system command execution entirely** by using language-native library alternatives.

**Primary Defence:** Replace ALL system commands with built-in library functions. Common replacements across languages:

- **File operations → Native file I/O libraries** (Java: Files API, Python: pathlib/shutil, Node.js: fs module)
- **HTTP requests → Native HTTP clients** (Not curl/wget commands)
- **Archive operations → Native compression libraries** (Not tar/zip/unzip commands)

Only if no library alternative exists should you use parameterized execution APIs (ProcessBuilder, subprocess with list arguments, etc.) - and NEVER suggest input validation as the primary fix.

## Key Remediation Principles

- **ALWAYS eliminate OS command execution** - Replace with language-native APIs first (this is the ONLY complete remediation)
- **Input validation is NOT a fix** - It's only defense-in-depth; attackers can often bypass validation
- **Never suggest "secure" command execution patterns** as the primary solution - library alternatives must be evaluated first
- Use parameterized APIs only as a last resort when no library alternative exists (ProcessBuilder, subprocess with list args, execve)
- Avoid shell interpreters completely - Never use sh, bash, cmd.exe, or any shell invocation
- Apply defence in depth: After eliminating commands, add validation and least privilege

## Remediation Steps

- Identify all system command execution points (Runtime.exec, subprocess.call, exec, system(), shell_exec, etc.)
- **Determine the library-based alternative** for each command's purpose (see language-specific guidance)
- Replace system commands with appropriate native APIs
- For any truly unavoidable commands, use parameterized execution APIs with argument arrays
- Remove all shell invocation patterns (sh -c, bash -c, cmd /c, shell=True, etc.)
- Validate any remaining user input against strict allowlists (permitted values only, never denylists)
- Implement least privilege - Run any remaining processes with minimal required permissions
