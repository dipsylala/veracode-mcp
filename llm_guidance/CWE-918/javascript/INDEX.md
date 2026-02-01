# CWE-918: Server-Side Request Forgery (SSRF) - JavaScript/Node.js

## LLM Guidance

SSRF in Node.js occurs when applications fetch remote resources using user-supplied URLs without validation, enabling attackers to access internal services, cloud metadata endpoints, and bypass firewalls. The primary defense is validating URLs against an allowlist of permitted domains before making any HTTP requests.

## Remediation Strategy

- Allowlist domains: Only permit requests to explicitly approved domains/hosts
- Block private networks: Reject private IP ranges (10.x, 172.16-31.x, 192.168.x, 127.x, localhost)
- Disable redirects: Prevent attackers from bypassing validation via HTTP redirects
- Parse and validate: Use `URL` constructor to parse and validate scheme, hostname, and port

## Remediation Steps

- Create an allowlist of permitted domains/hosts for external requests
- Parse user input with `new URL()` and validate hostname against allowlist
- Reject private IP addresses and localhost addresses
- Disable automatic redirect following in HTTP clients
- Validate resolved IPs before connecting (DNS rebinding protection)
- Use network-level controls to restrict outbound connections

## Minimal Safe Pattern

```javascript
const ALLOWED_HOSTS = ['api.trusted-service.com', 'cdn.example.com'];

async function safeFetch(userUrl) {
  const url = new URL(userUrl); // Throws on invalid URL
  
  if (!ALLOWED_HOSTS.includes(url.hostname)) {
    throw new Error('Domain not allowed');
  }
  
  if (url.protocol !== 'https:') {
    throw new Error('Only HTTPS allowed');
  }
  
  return fetch(url.href, { redirect: 'manual' });
}
```
