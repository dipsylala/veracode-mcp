---
name: fixit
description: |
  Provides remediation guidance and fixes for a number of languages and libraries

  - User asks to fix security vulnerabilities
  - User mentions "fixit", "fix flaws", "remediate flaws"
  - User wants to fix a specific CVE,Flaw ID, or vulnerability type (XSS, SQL injection, path traversal, etc.)
  - User wants to upgrade a vulnerable dependency
  - User asks to "fix all" vulnerabilities or "fix all high/critical" issues

allowed-tools:
  - Read
  - Write
  - Edit
  - Bash
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

**Core Principle**: Choose dependencies wisely to minimize supply chain risk.

## Parse the request

### Fixing flaws
If the user wants to fix a particular flaw Id, call remediation_guidance with that ID.

### Fixing 3rd party dependencies
If the user wants to fix a CVE ID, review any prior output from the tool run-sca-scan and see if there is a suggested version to resolve that  CVE. If there is, review other flaws for the same component and see if there are higher versions of the flawed component recommended.

If a user wants to fix a particular 3rd party component, review any prior output from the tool run-sca-scan and check the highest recommended version across all CVEs.

When upgrading a vulnerable dependency:

 - If it is a transitive dependency, check if there is a non-vulnerable version of the direct dependency that includes a non-vulnerable version of the transitive dependency. If so, upgrade the direct dependency to that version.
 - Upgrade to latest or least-vulnerable version via package manager (npm, pip, maven, etc.)
 - Replace with alternative component with similar functionality
 - Remove dependency if not needed
 - run any commands necessary to clear caches, update lockfiles, or otherwise complete the upgrade process

 **Only update package manifests or lockfiles directly if the package manager is unable to or is not present**
