# CWE-83: XSS (Improper Neutralization)

## LLM Guidance

Cross-Site Scripting (XSS) via improper neutralization of script-related HTML occurs when applications fail to properly encode, escape, or sanitize user-controlled data before including it in HTML output. This allows attackers to inject malicious scripts (typically JavaScript) that execute in victims' browsers within the security context of the vulnerable application. XSS attacks enable session hijacking, credential theft, phishing, keylogging, defacement, malware distribution, and unauthorized actions on behalf of authenticated users.

## Remediation Strategy

- Never trust user input in HTML contexts—treat all user data as untrusted
- Apply context-appropriate output encoding based on where data is inserted (HTML body, attributes, JavaScript, URL, CSS)
- Use security-focused encoding libraries rather than building custom solutions
- Implement Content Security Policy (CSP) as defense-in-depth
- Validate and sanitize input at boundaries, but rely on output encoding as primary defense

## Remediation Steps

- Identify XSS type - Determine if vulnerability is Reflected (immediate response), Stored (persisted), DOM-based (client-side), or Mutation XSS (mXSS)
- Apply HTML entity encoding for user data in HTML body context (encode `<`, `>`, `&`, `"`, `'`)
- Use JavaScript encoding for data inserted into JavaScript contexts (escape quotes, backslashes, control characters)
- Apply URL encoding for data in URL parameters or href attributes
- Implement CSP headers to restrict script sources and prevent inline script execution
- Use auto-escaping template engines and framework-provided sanitization functions
