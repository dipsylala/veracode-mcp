# CWE-601: URL Redirection to Untrusted Site (Open Redirect)

## LLM Guidance

Open redirect vulnerabilities occur when applications accept user-controllable input to determine redirect destinations without proper validation. Attackers exploit this by crafting malicious links from trusted domains that redirect victims to attacker-controlled sites. The core fix is to never allow untrusted input to control navigation targets—use server-defined destinations or strict allowlists only.

## Key Principles

- Never trust user input for redirect destinations
- Use server-defined redirect targets from an allowlist of approved URLs
- Validate redirect parameters against exact matches, not pattern matching
- Prefer indirect references (IDs/keys) over direct URL parameters
- Reject redirects to external domains by default

## Remediation Steps

- Identify the vulnerability - Review security findings for the file, line number, and parameters controlling redirects (`?next=`, `?redirect=`, `?returnUrl=`)
- Trace data flow - Follow how untrusted input flows to redirect functions (`Response.Redirect()`, `header("Location -")`, `sendRedirect()`)
- Implement allowlist validation - Create a server-side list of permitted redirect destinations and validate against exact matches
- Use indirect references - Replace direct URL parameters with lookup keys that map to server-defined destinations
- Validate strictly - Ensure redirects are relative paths or exact matches to allowlisted absolute URLs
- Reject external URLs - Block redirects to domains outside your application unless explicitly required and allowlisted
