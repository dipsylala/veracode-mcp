# CWE-319: Cleartext Transmission of Sensitive Information

## LLM Guidance

Cleartext transmission occurs when sensitive data (passwords, tokens, PII, API keys, payment data) is sent over networks without encryption, making it readable to attackers. Network sniffing, man-in-the-middle attacks, and compromised infrastructure allow easy interception of unencrypted HTTP, WebSocket (ws://), FTP, or plaintext database traffic. The core fix: always use TLS/HTTPS and encrypted transport for all sensitive data.

## Key Principles

- Never transmit sensitive data over cleartext channels (HTTP, ws://, FTP, unencrypted databases)
- Require TLS 1.2+ for all communications, preferably TLS 1.3
- Use HTTPS for entire site, not just login/sensitive pages
- Implement certificate validation to prevent MITM attacks
- Apply encryption at transport layer and application layer when needed

## Remediation Steps

- Review flaw details to identify file, line number, and what sensitive data is transmitted in cleartext
- Trace data flow to determine protocol (HTTP vs HTTPS, ws -// vs wss -//, FTP vs SFTP)
- Migrate all endpoints to HTTPS/TLS-update URLs from http -// to https -// and ws -// to wss -//
- Configure TLS 1.2 minimum (disable TLS 1.0/1.1, SSL), prefer TLS 1.3
- Enable HSTS (HTTP Strict Transport Security) headers to enforce HTTPS
- Validate server certificates properly-don't disable certificate checks in production
