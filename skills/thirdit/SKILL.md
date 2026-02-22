---
name: thirdit
description: |
  Packages and starts a scan of the codebase to identify vulnerabilities and generate a report.

allowed-tools:
  - Read
  - Write
  - Edit
  - Bash
  - Grep
  - mcp_veracode_local-sca-scan
  - mcp_veracode_local-sca-findings
license: Apache-2.0
compatibility: Requires Veracode MCP server connection and authenticated Veracode account. Supports SAST and SCA for all major package managers.
metadata:
  author: Dipsylala
  version: 1.0.0
---

# Security Scanner

Helps developers to package and scan their third party libraries for security vulnerabilities.

## Parse the request

If the user wants to scan third party libraries call local-sca-scan to perform the scan. Run local-sca-findings to get the SCA results, default to a page size of 10
