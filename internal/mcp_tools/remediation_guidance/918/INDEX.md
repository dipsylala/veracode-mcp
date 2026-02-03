# CWE-918 - Server-Side Request Forgery (SSRF)

## LLM Guidance

Server-Side Request Forgery (SSRF) occurs when an application fetches remote resources based on user-supplied URLs without proper validation, allowing attackers to force the server to make requests to arbitrary destinations including internal services and cloud metadata endpoints. The vulnerability exploits the server's trusted network position. Never allow untrusted input to determine outbound request destinations.

## Key Principles

- Use URL allowlists - Maintain explicit allowlists of permitted domains/IPs and validate full URLs (scheme, host, port, path) against them-reject anything not on the list
- Block private IP ranges - Prevent access to internal networks (10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16), loopback (127.0.0.0/8), link-local (169.254.0.0/16), and metadata endpoints
- Perform DNS validation - Resolve URLs and check that resolved IPs don't point to internal resources
- Never use denylists - Attackers will bypass them; allowlists are the only effective approach
- Enforce network egress controls - Limit outbound connections at the infrastructure level

## Remediation Steps

- Locate SSRF vulnerabilities by tracing untrusted data flow from input sources to HTTP request functions (`requests.get()`, `fetch()`, `curl_exec()`)
- Implement URL allowlists as primary defence - validate full URLs against permitted destinations before making requests
- Add IP validation by resolving URLs and blocking private/internal IP ranges
- Use safe URL parsing libraries to prevent bypasses via encoding or URL manipulation
- Apply defence-in-depth with network segmentation and egress filtering
- Test with various bypass techniques (DNS rebinding, redirects, alternate encodings)
