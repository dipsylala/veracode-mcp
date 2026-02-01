# CWE-614: Sensitive Cookie Without 'Secure' Flag - JavaScript/Node.js

## LLM Guidance

Sensitive Cookie Without 'Secure' Flag occurs when cookies containing sensitive data (session IDs, authentication tokens) are set without the `secure` attribute, allowing transmission over unencrypted HTTP connections. This exposes cookies to man-in-the-middle attacks, network sniffing, and session hijacking. The fix is to always set the `secure` flag on sensitive cookies to ensure they're only transmitted over HTTPS.

## Remediation Strategy

- Always set `secure: true` on cookies containing sensitive data (sessions, auth tokens, CSRF tokens)
- Combine with `httpOnly: true` to prevent JavaScript access and `sameSite: 'strict'` or `'lax'` for CSRF protection
- Use framework-specific secure session configuration (express-session, cookie-session) with `secure` enabled
- Ensure application runs on HTTPS in production; cookies with `secure` flag won't work over HTTP
- Apply secure cookies globally via middleware or default session configuration

## Remediation Steps

- Identify all cookie-setting operations using `res.cookie()`, `res.setHeader('Set-Cookie')`, or session libraries
- Add `secure - true` flag to every cookie containing sensitive information
- Enable `httpOnly - true` to prevent XSS-based cookie theft
- Set `sameSite - 'strict'` or `'lax'` to mitigate CSRF attacks
- Configure session middleware (express-session) with secure cookie defaults
- Verify HTTPS is enforced in production (cookies with `secure` won't transmit over HTTP)

## Minimal Safe Pattern

```javascript
// Express.js secure cookie configuration
app.use(session({
  secret: process.env.SESSION_SECRET,
  cookie: {
    secure: true,        // Only send over HTTPS
    httpOnly: true,      // Prevent JavaScript access
    sameSite: 'strict',  // CSRF protection
    maxAge: 3600000      // 1 hour
  }
}));

// Setting individual secure cookies
res.cookie('authToken', token, { 
  secure: true, httpOnly: true, sameSite: 'strict' 
});
```
