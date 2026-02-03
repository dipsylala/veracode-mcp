# CWE-470: Unsafe Reflection

## LLM Guidance

Unsafe reflection occurs when applications use untrusted input to select classes, methods, or code via reflection APIs (Class.forName, eval, import), enabling arbitrary code execution and complete application compromise. Never allow untrusted input to directly select classes/types/methods for execution.

## Key Principles

- Allowlist over blacklist: Define explicit permitted classes/methods; reject all others
- Indirect mapping: Map user input to safe identifiers, not class/method names directly
- Avoid reflection: Use polymorphism, factory patterns, or strategy patterns instead
- Input validation: If reflection required, validate against strict allowlist before use
- Least privilege: Restrict reflection to minimum required classes/methods

## Remediation Steps

- Locate vulnerability - Review scan data_paths to find where untrusted input flows to `Class.forName()`, `getMethod()`, `eval()`, `ScriptEngine.eval()`, or dynamic imports
- Map to allowlist - Create `Map<String, Class<?>>` mapping safe identifiers to permitted classes; reject unmapped input
- Use factory pattern - Replace reflection with factory that returns instances based on allowlisted types
- Validate strictly - If reflection unavoidable, validate input against allowlist before `Class.forName()` or `getMethod()`
- Remove eval - Replace `eval()` and script engines with type-safe alternatives
- Test coverage - Verify malicious class names (e.g., `java.lang.Runtime`, `ProcessBuilder`) are rejected
