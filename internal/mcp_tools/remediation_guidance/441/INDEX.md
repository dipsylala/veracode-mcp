# CWE-441: Unintended Proxy or Intermediary ('Confused Deputy')

## LLM Guidance

Unintended proxy vulnerabilities occur when applications forward requests to arbitrary destinations based on user input, acting as open proxies that enable SSRF attacks, bypassing firewalls, port scanning internal networks, and accessing cloud metadata services. The core issue is allowing intermediaries to become confused deputies that perform actions on behalf of untrusted callers without proper authentication or authorization constraints.

## Key Principles

- Never trust user-supplied URLs or destinations - treat all external input as hostile when determining where to send requests
- Implement strict allowlists - only permit connections to explicitly approved domains, IPs, and ports
- Authenticate and authorize the true caller - verify who is making the request and what they're permitted to access
- Constrain delegated actions - limit what the proxy can do even for authorized requests
- Apply defence in depth - combine input validation, network controls, and runtime restrictions

## Remediation Steps

- Examine data_paths in scan results to identify where user input controls request destinations
- Locate proxy endpoints including URL fetching, HTTP forwarding, webhook handlers, and redirect logic
- Implement destination allowlisting with explicit approved domains and reject all others by default
- Parse and validate URLs before use - check scheme, hostname, port, and IP ranges
- Block internal/private IP ranges including localhost, RFC1918, link-local, and cloud metadata endpoints (169.254.169.254)
- Disable automatic redirect following or validate redirect targets against the same allowlist
