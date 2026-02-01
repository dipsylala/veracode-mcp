# CWE-94: Code Injection - JavaScript/Node.js

## LLM Guidance

Code injection occurs when untrusted input flows into code execution functions like `eval()`, `Function()`, `setTimeout()`/`setInterval()` with strings, `vm.runInContext()`, or template engines, allowing attackers to execute arbitrary JavaScript. This grants full access to the application runtime, including file system, environment variables, and sensitive data.

## Key Principles

- Never pass user input to code evaluation functions (`eval`, `Function`, `vm` modules)
- Use safe alternatives: JSON.parse() for data, allow-lists for dynamic operations
- Sanitize template engine inputs and use auto-escaping modes
- Apply principle of least privilege to execution contexts
- Validate and restrict all dynamic code paths

## Remediation Steps

- Replace `eval()` with `JSON.parse()` for data parsing
- Convert `setTimeout(string)` to `setTimeout(function)` with callbacks
- Use allow-lists for dynamic property access instead of bracket notation with user input
- Configure template engines (EJS, Pug, Handlebars) with auto-escaping enabled
- If `vm` module is required, use isolated contexts with frozen globals
- Apply input validation at entry points before any processing

## Safe Pattern

```javascript
// UNSAFE: eval with user input
const result = eval(userInput);

// SAFE: Parse data, use allow-list for operations
const data = JSON.parse(userInput);
const allowedOps = { add: (a,b) => a+b, multiply: (a,b) => a*b };
const operation = allowedOps[data.operation];
if (operation) {
  const result = operation(data.a, data.b);
}

// SAFE: Function reference instead of string
setTimeout(() => handleTask(params), 1000);
```
