# CWE-117: Log Injection - LLM Guidance

## LLM Guidance

Log Injection occurs when untrusted user input is written to logs without proper validation or encoding, allowing attackers to forge log entries, hide malicious activity, or inject misleading information. This can compromise audit trails, inject false data into monitoring systems, or obscure security incidents.

## Key Principles

- Ideally, use structured JSON/ECS logging to separate data from structure
- If JSON/ECS logging is not possible, always encode untrusted input so it appears as data values, not log control characters
- Never concatenate user input directly into log messages
- Validate and sanitize all external data before logging
- Configure logging frameworks to auto-escape or encode fields

## Remediation Steps

- Review security findings to identify where untrusted data is written to logs
- Locate the source where untrusted data enters (HTTP parameters, headers, cookies, files, databases, network requests)
- Trace to the sink by finding the logging statement (`logger.info()`, `log.warn()`, `console.log()`, etc.)
- Check the data flow through each frame in the data path for missing encoding or validation
- Implement structured logging by configuring your logger to emit JSON/ECS format at the sink
- Use parameterized logging instead of string concatenation to ensure automatic encoding of user input
