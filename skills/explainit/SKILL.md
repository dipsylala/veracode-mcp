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

**Core Principle**: Call `remediation_guidance` to obtain authoritative flaw details, then explain the result in clear, contextual language suited to the developer's level.

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

1. **What the flaw is** — the vulnerability class (e.g. SQL injection, path traversal) in plain terms.
2. **Why it is a risk** — the potential impact if exploited (data exposure, code execution, privilege escalation, etc.).
3. **Where it exists** — the affected file, function, or component if the information is available.
4. **How to address it** — a brief, non-prescriptive summary of the remediation approach (full fix guidance belongs to the `/fixit` skill).

Keep explanations concise and relate them to the developer's context. Avoid unnecessary acronyms or academic references unless the user asks for deeper detail.