---
name: scanit
description: |
  Packages and starts a scan of the codebase to identify vulnerabilities and generate a report.

allowed-tools:
  - Read
  - Write
  - Edit
  - Bash
  - Grep
  - mcp_veracode_package-workspace
  - mcp_veracode_pipeline-scan
  - mcp_veracode_pipeline-status
  - mcp_veracode_pipeline-findings
license: Apache-2.0
compatibility: Requires Veracode MCP server connection and authenticated Veracode account. Supports SAST and SCA for all major package managers.
metadata:
  author: Dipsylala
  version: 1.0.0
---

# Security Scanner

Helps developers to package and scan their first party code for security vulnerabilities.

## Parse the request

if the user wants to scan an application call package_workspace to prepare the code for scanning, then call pipeline_scan to start the scan. Default to a page size of 10. Let the user know that they can use pipeline_status to check when the scan has finished.
