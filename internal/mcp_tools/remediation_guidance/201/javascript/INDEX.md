# CWE-201: Insertion of Sensitive Information Into Sent Data - JavaScript

## LLM Guidance

CWE-201 occurs when JavaScript/Node.js applications expose sensitive data (passwords, tokens, API keys, stack traces, internal paths) in HTTP responses, error messages, logs, or client-side code. The core fix is to sanitize all outbound data by filtering sensitive fields, using allowlists for response properties, and implementing proper error handling that hides internal details from end users.

## Key Principles

- Sanitize error objects and responses by explicitly selecting only non-sensitive fields for transmission
- Use environment variables for secrets and never embed them in client-side JavaScript bundles
- Implement centralized error handling middleware that returns generic error messages to clients
- Apply response filtering to remove sensitive fields like passwords, tokens, and internal IDs before sending
- Log detailed errors server-side only; send sanitized messages to clients

## Remediation Steps

- Review all API response handlers and remove sensitive fields using allowlists or field exclusion
- Configure error handling middleware to catch exceptions and return generic error messages
- Audit client-side code (React/Vue components) to ensure no secrets are embedded in bundles
- Use `.env` files with tools like `dotenv` and never commit secrets to version control
- Implement response serializers/transformers that explicitly define allowed fields
- Add logging sanitization to strip sensitive data before writing to logs or external services

## Safe Pattern

```javascript
// Express error handler - sanitize errors before sending
app.use((err, req, res, next) => {
  // Log full error server-side
  console.error(err);
  
  // Send generic message to client
  res.status(err.status || 500).json({
    error: 'An error occurred',
    requestId: req.id
  });
});

// Sanitize user object before sending
const sanitizeUser = (user) => ({
  id: user.id,
  email: user.email,
  name: user.name
  // Exclude: password, token, internalId
});
```
