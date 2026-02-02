# CWE-20: Improper Input Validation

## LLM Guidance

Improper Input Validation occurs when applications fail to enforce constraints on externally supplied data before use. While validation alone doesn't prevent all vulnerabilities, its absence allows unexpected or malformed values to reach security-sensitive logic, enabling SQL injection, XSS, command injection, and other attacks. Static analysis flags CWE-20 when code accepts external input without clear type, range, format, or semantic constraints.

## Key Principles

- CWE-20 is abstract - the correct fix depends on identifying the specific child CWE that matches how unvalidated input is actually used
- Address the specific security risk - remediation must target the vulnerability enabled by unconstrained input, not just add generic validation
- Trace data flow completely - follow input from entry point through all transformations to where it's consumed
- Defense-in-depth - combine input validation with context-specific output encoding, parameterized queries, and principle of least privilege

## Remediation Steps

- Identify input origin - HTTP parameter, header, file upload, API request, database result, or external system
- Trace to consumption point - determine where input is used - HTML output, SQL query, shell command, file path, LDAP query, XML parser, eval statement
- Classify the specific vulnerability - map usage to child CWE (SQL Injection, XSS, Path Traversal, Command Injection, etc.)
- Apply appropriate defense - use parameterized queries for SQL, context-aware encoding for HTML, allowlists for file paths, avoid shell execution
- Enforce input constraints - validate type, length, format, character set, and business logic rules at entry points
- Reject invalid input - fail securely by rejecting malformed data rather than attempting sanitization
