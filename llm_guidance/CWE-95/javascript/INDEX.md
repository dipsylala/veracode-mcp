# CWE-95: Eval Injection - JavaScript

## LLM Guidance

JavaScript eval injection occurs when untrusted input flows into dynamic code execution functions like `eval()`, `Function()`, or `setTimeout()`/`setInterval()` with strings. Attackers can execute arbitrary code, access sensitive data, or compromise the application. Never pass user-controlled data to these functions; use safer alternatives like object property access, JSON parsing, or predefined function mappings.

## Remediation Strategy

- Eliminate dynamic code execution: Replace `eval()` and `Function()` with static alternatives
- Use allowlists: Map user input to predefined functions or values rather than executing strings
- Validate strictly: If dynamic execution is unavoidable, validate input against narrow allowlists
- Parse safely: Use `JSON.parse()` for data, never `eval()` for JSON
- Sanitize module paths: Validate and restrict dynamic `require()` or `import()` calls

## Remediation Steps

- Search codebase for `eval()`, `Function()`, `setTimeout/setInterval` with string args, and dynamic `require()`
- Replace `eval(json)` with `JSON.parse(json)`
- Convert dynamic property access using bracket notation - `obj[userInput]` instead of `eval('obj.' + userInput)`
- Refactor string-based function calls to object mappings or switch statements
- For unavoidable cases, validate input against strict allowlists before execution
- Enable CSP headers with `unsafe-eval` disabled to prevent runtime eval

## Minimal Safe Pattern

```javascript
// UNSAFE: eval with user input
const result = eval(userInput);

// SAFE: Use object mapping
const allowedFunctions = {
  add: (a, b) => a + b,
  multiply: (a, b) => a * b
};
const operation = allowedFunctions[userInput];
if (operation) {
  const result = operation(x, y);
}
```
