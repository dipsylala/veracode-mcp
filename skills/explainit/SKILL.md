---
name: explainit
description: |
  Provides remediation guidance and explanations for a number of languages and libraries

  - User asks to fix security vulnerabilities
  - User mentions "help", "explain", "describe"
  - User wants to understand a specific CVE,Flaw ID, or vulnerability type (XSS, SQL injection, path traversal, etc.)

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

# Secure Dependency Advisor

Help developers and AI agents make informed decisions when selecting open-source packages by evaluating security health, vulnerability history, and maintenance status.

## Parse the request

### Explaining flaws
The user will have received a list of flaws and wants to understand them better. The remediation_guidance tool will help the user when provided a CVE or flaw ID. It performs the appropriate calls to retrieve flaw details and will return guidance. No call to finding_details is necessary, as remediation_guidance will provide all necessary information in its response.

The key here is to help the user understand the nature of the flaw, its impact, and potential remediation steps without overwhelming them with technical jargon. Focus on clear, concise explanations that relate to the user's context and needs.