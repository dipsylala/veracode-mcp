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

# Primary Purpose

Performs packaging and static scanning for an application

## Parse the request

* Calls the package_workspace MCP endpoint with the application_path pointing to the workspace root.
* Calls the pipeline_scan MCP endpoint to start the scan, with the application_path pointing to the workspace root.
* Calls the local_sca_scan MCP endpoint to start an SCA scan, with the application_path pointing to the workspace root.
* Let the user know that they can use pipeline_status to check when the Static scan has finished, but do not run it yourself
