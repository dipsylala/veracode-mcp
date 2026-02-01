# CWE-80: Improper Neutralization of Script-Related HTML Tags (Basic XSS) - JavaScript

## LLM Guidance

Cross-Site Scripting (CWE-80) occurs when untrusted data is inserted into web pages without proper encoding, allowing attackers to inject malicious scripts that execute in victim browsers. In JavaScript applications, this happens when user input is directly inserted into the DOM via `innerHTML` or HTML attributes without sanitization. Use `textContent` or framework-provided escaping mechanisms to safely display user input.

## Key Principles

- Use `textContent` or `innerText` instead of `innerHTML` for displaying user-controlled data
- Leverage framework built-in escaping (React JSX, Vue templates, Angular bindings)
- Apply context-aware output encoding (HTML, JavaScript, URL, CSS contexts)
- Implement Content Security Policy (CSP) headers to restrict script execution
- Sanitize HTML content with libraries like DOMPurify when rich content is required

## Remediation Steps

- Identify all locations where user input is rendered to the DOM
- Replace `innerHTML`, `outerHTML`, and `document.write()` with safe alternatives
- Use `textContent` for plain text or framework escaping for dynamic content
- For HTML rendering requirements, sanitize with DOMPurify before insertion
- Set restrictive CSP headers (`script-src 'self'`) to block inline scripts
- Validate and encode data based on output context (HTML entity encoding, JavaScript escaping, URL encoding)

## Safe Pattern

```javascript
// UNSAFE: Direct innerHTML with user input
element.innerHTML = userInput;

// SAFE: Use textContent for plain text
element.textContent = userInput;

// SAFE: Framework escaping (React example)
return <div>{userInput}</div>;

// SAFE: Sanitize HTML if rich content needed
import DOMPurify from 'dompurify';
element.innerHTML = DOMPurify.sanitize(userInput);
```
