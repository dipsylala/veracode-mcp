# CWE-384: Session Fixation

## LLM Guidance

Session Fixation occurs when an application allows an attacker to set or reuse a session identifier for another user, enabling the attacker to hijack the victim's session after authentication. The core fix is to regenerate session identifiers whenever authentication or privilege level changes, ensuring that pre-authentication sessions never remain valid post-authentication.

## Key Principles

- Regenerate session IDs immediately after authentication or privilege escalation
- Never accept session identifiers from URL parameters or untrusted sources
- Invalidate old session identifiers to prevent reuse
- Use framework-native session regeneration functions (session.regenerate(), session_regenerate_id())
- Bind sessions to additional user context for defense-in-depth

## Remediation Steps

- Review authentication flows (login handlers, OAuth callbacks, SSO) to identify where session IDs are not regenerated
- Call session.regenerate() or framework equivalent immediately after successful authentication
- Invalidate the previous session ID to prevent attackers from reusing it
- Reject session IDs passed via URL parameters or query strings
- Verify session ID changes before and after login during testing
- Bind sessions to additional context (IP address, user agent) as a secondary control
