# CWE-94: Improper Control of Generation of Code (Code Injection) - Python

## LLM Guidance

Code injection in Python occurs when untrusted input is passed to code execution functions like `eval()`, `exec()`, `compile()`, or `__import__()`. This allows attackers to execute arbitrary Python code with full access to the application's runtime environment. Never use dynamic code execution with user input; use safe alternatives like literal evaluation, dictionaries for dispatch, or sandboxed parsers.

## Key Principles

- Never use `eval()`, `exec()`, or `compile()` with any user-controlled input
- Use `ast.literal_eval()` for safe evaluation of literals (strings, numbers, lists, dicts)
- Replace dynamic code execution with predefined function mappings or configuration
- Validate and sanitize all inputs before processing, using allowlists not denylists
- Use template engines with auto-escaping instead of string formatting for dynamic content

## Remediation Steps

- Replace `eval()` calls with `ast.literal_eval()` for parsing data structures
- Convert dynamic code patterns to dictionary-based function dispatch
- Use JSON parsing instead of evaluating Python code from external sources
- Implement strict input validation with type checking and range limits
- Remove or isolate any remaining exec/eval calls to sandboxed environments
- Audit all uses of `__import__()`, `compile()`, and `importlib` for user input

## Safe Pattern

```python
import ast
import json

# Safe: Use ast.literal_eval for literals
user_input = "[1, 2, 3]"
safe_data = ast.literal_eval(user_input)

# Safe: Use JSON for structured data
config = '{"action": "read", "limit": 10}'
params = json.loads(config)

# Safe: Use dictionary dispatch instead of eval
actions = {"read": read_func, "write": write_func}
action = actions.get(params["action"], default_func)
action()
```
