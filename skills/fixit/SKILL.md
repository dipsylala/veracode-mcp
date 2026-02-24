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
  - mcp_veracode_run-sca-scan
license: Apache-2.0
compatibility: Requires Veracode MCP server connection and authenticated Veracode account. Supports SAST and SCA for all major package managers.
metadata:
  author: Dipsylala
  version: 1.0.0
---

# Secure Remediation Advisor

Help developers fix security vulnerabilities identified by Veracode — whether first-party static flaws (SAST) or third-party dependency vulnerabilities (SCA).

**Core Principle**: Obtain authoritative remediation guidance before making code or dependency changes, then apply the smallest safe fix.

## Identify the request type

### Flaw IDs — first-party SAST findings

A flaw ID is a numeric identifier, optionally with a pipeline suffix:
- Plain numeric: `12345`
- With pipeline suffix: `12345-1`

**For numeric flaw IDs only**, call `remediation_guidance` with the flaw ID first to obtain fix guidance before making any code change. Prefer to explain the guidance first unless the user explicitly asks to fix immediately.

**Do not call `remediation_guidance` for CVE or SID identifiers** — those are third-party findings and are handled below.

### CVE / SCA IDs — third-party dependency findings

A third-party vulnerability identifier starts with `cve-` or `sid-` (case-insensitive), for example `CVE-2021-44228` or `sid-12345`.

Do not call `remediation_guidance` for these. Instead, use any prior output from `run-sca-scan` or other scan results already in context, and follow the **Fixing 3rd party dependencies** steps below.

## Fixing 3rd party dependencies

1. Review any prior output from `run-sca-scan` for the affected component.
2. Identify the suggested safe version that resolves the CVE/SID. If multiple CVEs affect the same component, select the highest recommended version across all of them.
3. Determine whether the dependency is direct or transitive:
   - **Transitive**: check whether a newer version of the *direct* dependency already bundles a safe version of the transitive one. If so, upgrade the direct dependency instead.
   - **Direct**: upgrade to the latest non-vulnerable version.
4. Apply the upgrade using the appropriate package manager (npm, pip, maven, gradle, go, etc.) and run any commands needed to clear caches and update lockfiles.
5. If no safe version exists, consider replacing the component with a maintained alternative or removing it if unused.

**Only edit package manifests or lockfiles by hand if the package manager is unavailable or unable to complete the upgrade.**
