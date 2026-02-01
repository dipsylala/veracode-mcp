# CWE-94: Code Injection (Improper Control of Code Generation)

## LLM Guidance

Code injection occurs when applications dynamically generate and execute code using untrusted input, allowing attackers to inject arbitrary code that executes within the application's runtime with full access to internals, variables, functions, database connections, and secrets. Unlike command injection that executes OS commands, code injection executes in the application's programming language context. Common vulnerable patterns include `eval()`, `exec()`, `Function()`, `compile()`, `ScriptEngine`, and unsafe template rendering.

## Key Principles

- Never execute dynamically generated code derived from untrusted input
- Remove eval/dynamic compilation functions or strictly sandbox with allowlists
- Use static code paths and predefined logic instead of dynamic execution
- Treat all user input as untrusted regardless of source (HTTP parameters, databases, APIs)
- Apply principle of least privilege to execution contexts

## Remediation Steps

- Trace data flow from source (HTTP parameters, form inputs, file uploads, API requests, database fields) to sink (eval(), exec(), Function(), compile(), ScriptEngine, unsafe templates)
- Review scan results for specific file paths, line numbers, and variable names where code execution occurs
- Replace dynamic code execution with safer alternatives - lookup tables, predefined functions, switch statements, or configuration-driven logic
- If dynamic execution is unavoidable, implement strict allowlists that validate input against known-safe values only
- Validate and sanitize all inputs with strict type checking before any processing
- Use sandboxed execution environments with restricted access to system resources
