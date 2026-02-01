# CWE-352: Cross-Site Request Forgery (CSRF) - JavaScript/Node.js

## LLM Guidance

CSRF vulnerabilities occur when state-changing endpoints don't verify that requests originated from the legitimate application, allowing attackers to trick users into executing unwanted actions. The core fix is implementing token-based verification where each form/request includes a secret token that the server validates. Use `csurf` middleware for Express or `@fastify/csrf-protection` for Fastify to automatically generate and validate CSRF tokens.

## Remediation Strategy

- Implement CSRF tokens for all state-changing operations (POST, PUT, DELETE, PATCH)
- Use SameSite cookie attribute (`SameSite=Lax` or `Strict`) as defense-in-depth
- Validate Origin/Referer headers for additional protection on critical endpoints
- Never rely solely on cookies for authentication without CSRF protection
- For REST APIs consumed by native apps, use token-based auth instead of cookies

## Remediation Steps

- Install CSRF middleware - `npm install csurf cookie-parser`
- Apply middleware globally or to protected routes requiring CSRF validation
- Generate and inject CSRF token into forms and AJAX request headers
- Configure client to send token in `_csrf` field (forms) or `X-CSRF-Token` header (AJAX)
- Set cookie SameSite attribute and verify implementation with security tests
- Handle CSRF errors gracefully with user-friendly messages

## Minimal Safe Pattern

```javascript
const csrf = require('csurf');
const cookieParser = require('cookie-parser');
const express = require('express');
const app = express();

app.use(cookieParser());
app.use(csrf({ cookie: { httpOnly: true, sameSite: 'strict' } }));

app.get('/form', (req, res) => {
  res.render('form', { csrfToken: req.csrfToken() });
});

app.post('/transfer', (req, res) => {
  // CSRF token automatically validated by middleware
  processTransfer(req.body);
});
```
