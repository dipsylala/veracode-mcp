---
name: thirdit
description: |
  Packages and starts a scan of the codebase to identify vulnerabilities and generate a report.

allowed-tools:
  - Read
  - Grep
  - mcp_veracode_local-sca-scan
  - mcp_veracode_local-sca-findings
  - mcp_veracode_local-iac-findings
license: Apache-2.0
compatibility: Requires Veracode MCP server connection and authenticated Veracode account. Supports SCA and IaC analysis for all major package managers.
metadata:
  author: Dipsylala
  version: 1.0.0
---

# Security Scanner

Helps developers to package and scan their third party libraries for security vulnerabilities, and check their IaC files for misconfigurations.

## Parse the request

* Calls the local-sca-scan MCP endpoint with the application_path pointing to the workspace root.
* Display the results using local-sca-findings and local-iac-findings, defaulting to a page size of 10
