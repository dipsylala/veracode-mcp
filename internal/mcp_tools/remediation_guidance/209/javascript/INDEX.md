# CWE-209: Error Message Information Leak - JavaScript

## LLM Guidance

Error Message Information Leak occurs when JavaScript applications expose sensitive details like stack traces, database queries, file paths, or internal system information through API responses, error pages, or console output. The core fix is to sanitize error responses in production by logging detailed errors server-side while returning generic messages to clients. Node.js frameworks (Express, Fastify, Koa, Next.js) require proper error middleware configuration to prevent disclosure.

## Remediation Strategy

- Log detailed errors server-side with monitoring tools; never expose stack traces or internals to clients
- Return generic error messages in production (e.g., "An error occurred") while preserving specific errors for development
- Configure framework error handlers to distinguish between development and production environments
- Sanitize database errors to remove query details, table names, and schema information
- Disable debug mode and verbose logging in production deployments

## Remediation Steps

- Add environment-based error middleware that checks `NODE_ENV === 'production'`
- Replace detailed error responses with generic messages like "Internal server error"
- Configure logging to capture full errors server-side (Winston, Pino, Bunyan)
- Remove or disable client-side `console.error()` calls exposing sensitive data
- Set `NODE_ENV=production` in deployment environments
- Test error responses to verify no stack traces or internals leak

## Minimal Safe Pattern

```javascript
// Express production error handler
app.use((err, req, res, next) => {
  // Log full error server-side
  logger.error({ err, req: { method: req.method, url: req.url } });
  
  // Return sanitized response
  res.status(err.status || 500).json({
    error: process.env.NODE_ENV === 'production' 
      ? 'An error occurred' 
      : err.message
  });
});
```
