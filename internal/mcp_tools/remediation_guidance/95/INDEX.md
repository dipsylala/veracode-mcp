# CWE-95: Eval Injection - LLM Guidance

## LLM Guidance

Eval Injection occurs when untrusted input (from HTTP requests, external APIs, databases, files, or message queues) is passed to dynamic code execution functions like `eval()`, `exec()`, `Function()`, or `compile()`, allowing attackers to execute arbitrary code within the application context. The core fix is to never execute dynamically generated code from untrusted sources-eliminate eval/exec-style functions entirely and replace them with safe parsers or allowlisted interpreters.

## Key Principles

- Never use `eval()`, `exec()`, or similar dynamic code execution functions with untrusted data
- Replace dynamic code evaluation with safe alternatives: JSON parsers, template engines with auto-escaping, or expression evaluators with strict allowlists
- Treat all external data sources (user input, APIs, databases, files, configuration) as untrusted
- Use static analysis tools to detect and eliminate dangerous functions

## Actionable Steps

- Trace data flow: Identify where untrusted data enters (source), how it moves through the application, and where it reaches code execution functions (sink)
- Remove eval/exec functions: Refactor code to eliminate `eval()`, `exec()`, `Function()`, `compile()`, and similar dynamic execution entirely
- Use safe alternatives: Replace with JSON.parse() for data, template engines for rendering, or sandboxed expression evaluators with strict syntax allowlists
- Validate and sanitize: If dynamic evaluation is unavoidable, implement strict allowlists for permitted operations and reject any input that doesn't match
- Apply defence in depth: Run code in sandboxed environments with minimal privileges and monitor for suspicious execution patterns
- Conduct security review: Use static analysis tools and manual code review to find all instances of dangerous functions
