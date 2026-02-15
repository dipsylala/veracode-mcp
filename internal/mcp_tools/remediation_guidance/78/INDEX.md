# CWE-78: OS Command Injection

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands on the host. The primary remediation is to eliminate system command execution entirely by using language-native library alternatives (file I/O APIs, HTTP clients, compression libraries). Only if no library exists should parameterized execution APIs be considered.

## Key Principles

- Eliminate OS command execution completely - Replace with built-in library functions as the primary defence
- Never concatenate user input into command strings
- Avoid shell interpreters (sh, bash, cmd.exe) completely
- Use parameterized APIs only as a last resort when no native library exists (ProcessBuilder, subprocess with list arguments)
- Input validation is insufficient - It's only effective as a secondary defence layer
- Apply least privilege to any remaining process execution

## Remediation Steps

- Identify all command execution points (Runtime.exec, subprocess.call, exec, system(), shell_exec, Process.Start, etc.)
- Determine the native library alternative for each command's purpose (file operations → File I/O APIs, HTTP requests → HTTP clients, etc.)
- Replace system commands with appropriate language-native APIs
- For truly unavoidable commands, use parameterized execution APIs with separate argument arrays (never shell invocation)
- Remove all shell patterns and string concatenation in command construction
- Add input validation as a secondary defence layer using strict allowlists
- Apply least privilege principles to any remaining process execution
