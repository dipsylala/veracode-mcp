# CWE-77: Command Injection

## LLM Guidance

Command injection occurs when applications construct system commands using untrusted input without proper sanitization, allowing attackers to inject shell metacharacters and execute arbitrary commands with application privileges. The core fix is to avoid shell execution entirely and use parameterized process APIs where input remains a single argument that cannot alter command structure.

## Key Principles

- Never allow untrusted input to influence shell command structure
- Avoid shell execution entirely; use direct process invocation instead
- Use parameterized process APIs where input stays as a single argument
- Apply strict allowlisting for required parameters
- Run commands with least privilege necessary

## Remediation Steps

- Identify all sources of untrusted data (user input, external files, databases, network requests)
- Trace data flow to command execution functions (system(), exec(), Runtime.exec(), Process.Start())
- Check if shell is invoked (string form vs array form, shell=True flags)
- Replace shell invocation with array/list form direct process execution
- Implement strict allowlisting for any required command parameters
- Apply principle of least privilege to execution contexts
