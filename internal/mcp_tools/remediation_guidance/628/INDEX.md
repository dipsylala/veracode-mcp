# CWE-628: Function Call with Incorrectly Specified Arguments

## LLM Guidance

Incorrect function arguments (wrong type, wrong order, wrong count, null when required) cause undefined behavior, security check bypass, buffer overflows, null pointer dereferences, and logic errors, often due to API misuse or type confusion. Call functions with correct argument types, order, and count to prevent memory and state corruption.

## Key Remediation Principles

- Enforce strict type checking and argument validation at function boundaries
- Use API wrappers or interfaces that prevent argument misuse
- Validate argument order for functions with multiple same-type parameters
- Reject null arguments for security-critical functions unless explicitly allowed
- Apply compiler warnings and static analysis to catch mismatches early

## Remediation Steps

- Locate vulnerable calls - Search for functions with multiple same-type parameters, buffer operations (memcpy, strncpy), and authentication/authorization functions
- Verify argument order - Check dest/src positioning, user/password order, and size/length parameter placement
- Add type safety - Use strongly-typed wrappers, enums instead of integers, and null-safe types
- Validate inputs - Check for null, verify types match expectations, and ensure counts align with buffer sizes
- Enable compile-time checks - Turn on strict compiler warnings, use static analyzers, and apply linting rules for function calls
- Test edge cases - Verify behavior with null, wrong types, swapped arguments, and boundary values
