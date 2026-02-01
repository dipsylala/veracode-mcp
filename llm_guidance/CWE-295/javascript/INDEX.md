# CWE-295: Improper Certificate Validation - JavaScript/Node.js

## LLM Guidance

Improper certificate validation in JavaScript/Node.js applications allows man-in-the-middle (MITM) attacks where attackers intercept and modify HTTPS communications. This occurs when applications disable TLS/SSL certificate verification, fail to validate certificate chains, don't check hostname matching, or accept self-signed certificates in production. The fix is to always enable strict certificate validation and never disable `rejectUnauthorized` in production code.

## Remediation Strategy

- Always keep `rejectUnauthorized: true` (the default) for HTTPS requests in production
- Use proper certificate authorities (CA) and avoid self-signed certificates in production
- Validate certificate hostname matches the target server
- Pin certificates only when necessary and maintain proper update procedures
- Use updated Node.js versions with current TLS/SSL libraries

## Remediation Steps

- Remove any `rejectUnauthorized - false` or `NODE_TLS_REJECT_UNAUTHORIZED='0'` settings
- Use built-in `https` module with default settings or libraries that respect certificate validation
- Configure custom CAs using `ca` option instead of disabling validation
- Implement certificate pinning only if required, using tools like `node-http-mitm-proxy`
- Test HTTPS connections in staging with valid certificates
- Enable strict transport security headers (HSTS) to enforce HTTPS

## Minimal Safe Pattern

```javascript
const https = require('https');

// Secure HTTPS request with proper certificate validation
const options = {
  hostname: 'api.example.com',
  port: 443,
  path: '/data',
  method: 'GET',
  // rejectUnauthorized defaults to true - don't override it
};

https.request(options, (res) => {
  res.on('data', (d) => process.stdout.write(d));
}).end();
```
