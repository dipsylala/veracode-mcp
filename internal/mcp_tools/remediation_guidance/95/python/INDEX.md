# CWE-95: Eval Injection - Python

## LLM Guidance

CWE-95 occurs when untrusted input from HTTP requests, APIs, files, or external sources is passed to dynamic code execution functions like `eval()`, `exec()`, `compile()`, or `__import__()`. These functions execute arbitrary Python code with full application privileges, enabling attackers to access data, modify system state, or execute commands. The core fix is to eliminate dynamic code execution entirely and use safe alternatives like whitelisted function mappings, AST parsing, or data serialization.

## Key Principles

- Never use `eval()`, `exec()`, `compile()`, or `__import__()` with untrusted input
- Replace dynamic execution with static alternatives: dictionaries, match/case statements, or allowlists
- Use safe parsers like `ast.literal_eval()` for data structures or JSON for configuration
- Validate and sanitize all input with strict whitelists before any processing
- Apply least-privilege principles to limit damage if execution occurs

## Remediation Steps

- Identify all uses of `eval()`, `exec()`, `compile()`, and dynamic imports in the codebase
- Replace with safe alternatives - function dictionaries, `ast.literal_eval()`, or JSON parsing
- If dynamic execution is unavoidable, implement strict input validation with character and length limits
- Use sandboxing or restricted execution environments with minimal permissions
- Apply input whitelisting to allow only specific, predefined values or patterns
- Review and audit all external data sources feeding into the application

## Safe Pattern

```python
# UNSAFE: Using eval with user input
user_input = request.GET.get('expression')
result = eval(user_input)  # NEVER do this

# SAFE: Use a function mapping with whitelist
ALLOWED_OPERATIONS = {
    'add': lambda x, y: x + y,
    'subtract': lambda x, y: x - y,
    'multiply': lambda x, y: x * y
}

operation = request.GET.get('op')
if operation in ALLOWED_OPERATIONS:
    result = ALLOWED_OPERATIONS[operation](10, 5)
```
