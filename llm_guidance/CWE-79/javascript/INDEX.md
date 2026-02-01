# CWE-79: Cross-Site Scripting (XSS) - JavaScript/Node.js

## LLM Guidance

Cross-Site Scripting (XSS) occurs when untrusted data is rendered in web pages without proper encoding, allowing attackers to inject malicious scripts. The primary defense is using framework auto-escaping (React JSX, Vue templates), `textContent` for DOM manipulation, or DOMPurify for rich HTML sanitization.

## Remediation Strategy

- Use framework built-in escaping mechanisms (React JSX, Vue templates, template engines with escaping enabled)
- Never use `innerHTML`, `dangerouslySetInnerHTML`, or `v-html` with untrusted data
- Sanitize rich HTML with DOMPurify before rendering
- Apply Content Security Policy (CSP) headers as defense-in-depth
- Encode data appropriately for context (HTML, JavaScript, URL)

## Remediation Steps

- Replace `innerHTML` with `textContent` or framework-safe rendering
- Enable auto-escaping in template engines (EJS, Pug, Handlebars)
- Install and use DOMPurify for sanitizing user-generated HTML
- Set CSP headers - `Content-Security-Policy - default-src 'self'; script-src 'self'`
- Validate and encode URL parameters before rendering
- Use parameterized queries for database operations to prevent injection

## Minimal Safe Pattern

```javascript
// Safe: Use textContent for plain text
const userInput = req.query.name;
element.textContent = userInput;

// Safe: React auto-escapes JSX expressions
return <div>{userInput}</div>;

// Safe: Sanitize rich HTML with DOMPurify
import DOMPurify from 'dompurify';
const clean = DOMPurify.sanitize(userHTML);
element.innerHTML = clean;

// Safe: Template engine with escaping
app.set('view options', { escape: true });
res.render('page', { userInput });
```
