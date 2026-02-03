# CWE-77: Command Injection

## LLM Guidance

Command injection occurs when applications construct system commands using untrusted input without proper sanitization, allowing attackers to inject shell metacharacters and execute arbitrary commands with application privileges. The primary fix is to eliminate the vulnerability entirely by using native language libraries instead of executing system commands. If system commands are unavoidable, use parameterized process APIs where input remains a single argument that cannot alter command structure.

## Key Principles

- **BEST:** Use native language libraries for file operations, network requests, and system tasks to eliminate command execution entirely
- **If commands unavoidable:** Use parameterized process APIs where input stays as a single argument and cannot alter command structure
- Never allow untrusted input to influence shell command structure
- Avoid shell interpreters entirely; use direct process invocation with argument arrays
- Apply strict allowlisting for required parameters as defence-in-depth
- Run commands with least privilege necessary

## Remediation Steps

- Identify all sources of untrusted data (user input, external files, databases, network requests)
- Trace data flow to command execution functions (system(), exec(), Runtime.exec(), Process.Start())
- **Replace system commands with native language APIs** (file libraries, HTTP clients, etc.) to eliminate the vulnerability
- If commands are unavoidable, check if shell is invoked (string form vs array form, shell=True flags)
- Replace shell invocation with array/list form direct process execution with no shell interpreter
- Implement strict allowlisting for any required command parameters as additional defence
- Apply principle of least privilege to execution contexts
