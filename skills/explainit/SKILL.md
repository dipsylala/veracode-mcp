---
name: explainit
description: |
  Explains security vulnerabilities and flaws identified by Veracode

  - User asks to explain or understand a security vulnerability
  - User mentions "explainit", "explain", "describe", "what is", "tell me about"
  - User wants to understand a specific CVE, Flaw ID, or vulnerability type (XSS, SQL injection, path traversal, etc.)
  - User asks what a flaw means or why it is a risk

allowed-tools:
  - Read
  - Grep
  - mcp_veracode_remediation-guidance
license: Apache-2.0
compatibility: Requires Veracode MCP server connection and authenticated Veracode account. Supports SAST and SCA for all major package managers.
metadata:
  author: Dipsylala
  version: 1.0.0
---

# Security Flaw Explainer

Help developers understand security vulnerabilities identified by Veracode — what the flaw is, why it matters, and how it could be exploited — without overwhelming them with jargon.

**Core Principles**: Educate while you guide

## Identify the ID type

### Flaw IDs — first-party SAST findings

A flaw ID is a numeric identifier, optionally with a pipeline suffix:
- Plain numeric: `12345`
- With pipeline suffix: `12345-1`

### CVE / SCA IDs — third-party dependency findings

A third-party vulnerability identifier starts with `cve-` or `sid-` (case-insensitive), for example `CVE-2021-44228` or `sid-12345`.

## Retrieving flaw details

**For numeric flaw IDs only** (`12345` or `12345-1`), call `remediation_guidance` with the flaw ID. This single call retrieves all necessary details — do not call `finding_details` separately, as `remediation_guidance` already includes that information in its response.

**For CVE / SCA IDs** (`cve-` or `sid-` prefixes), do not call `remediation_guidance`. Instead, use any prior output from `run-sca-scan` or other scan results already in context to explain the vulnerability.

## Presenting the explanation

Use the retrieved flaw details to explain:

1. **What the vulnerability is and where it exists** — describe the flaw class (SQL injection, XSS, path traversal, etc.) in plain language and identify the affected file, function, or line number if available.

2. **Why it's a risk** — reference the CWE identifier and explain the real-world impact if exploited (data exposure, code execution, privilege escalation, denial of service, etc.).

3. **What makes this code exploitable** — identify the specific unsafe pattern, missing validation, or improper API usage that creates the vulnerability (e.g., "line 42 concatenates user input directly into the SQL query").

4. **How to address it** — summarize the general remediation approach (e.g., "use parameterized queries," "validate and sanitize input," "apply context-aware encoding"). Keep this brief and conceptual — detailed step-by-step fixes belong to the `/fixit` skill.

5. **Related best practices** — mention relevant secure coding patterns for the language/framework, OWASP guidelines, or security principles that prevent this class of vulnerability.

Keep explanations concise, educational, and grounded in the developer's context. Avoid unnecessary acronyms or academic references unless the user asks for deeper detail.