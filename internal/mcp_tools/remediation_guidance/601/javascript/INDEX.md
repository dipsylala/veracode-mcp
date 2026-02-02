# CWE-601: Open Redirect - JavaScript/Node.js

## LLM Guidance

Open redirect vulnerabilities in Node.js occur when user-controlled input flows into redirect functions (`res.redirect()`, `res.writeHead()` Location headers, or `window.location`) without validation, enabling attackers to redirect users to malicious sites for phishing or credential theft. The core fix is to validate redirect destinations against an allowlist of trusted domains or use relative paths only.

**Primary Defence:** Validate all redirect URLs against an allowlist of allowed domains or restrict to relative paths only.

## Key Principles

- Implement allowlist validation for all redirect destinations before redirecting
- Use relative paths instead of absolute URLs when possible
- Reject or sanitize any redirect URL containing external domains
- Apply URL parsing to verify protocol and hostname match expected values

## Remediation Steps

- Parse the redirect URL using Node.js `URL` constructor to extract components
- Check if the URL is relative (no protocol/hostname) or matches allowlist
- For external redirects, compare hostname against approved domain list
- Reject invalid destinations with error or fallback to safe default page
- Log blocked redirect attempts for security monitoring

## Safe Pattern

```javascript
const ALLOWED_DOMAINS = ['example.com', 'app.example.com'];

function safeRedirect(res, redirectUrl, fallback = '/') {
  try {
    const url = new URL(redirectUrl, 'https://example.com');
    if (url.protocol === 'https:' && ALLOWED_DOMAINS.includes(url.hostname)) {
      return res.redirect(redirectUrl);
    }
  } catch {
    if (redirectUrl.startsWith('/')) return res.redirect(redirectUrl);
  }
  res.redirect(fallback);
}
```
