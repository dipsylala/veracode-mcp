---
name: reportit
description: |
  
  - Provides prioritisation, executive summary, and reporting of vulnerabilities in the codebase, based off the pipeline scan

allowed-tools:
  - Read
  - Grep
  - mcp_veracode_pipeline-findings
  - mcp_veracode_local-sca-findings

license: Apache-2.0
compatibility: Requires Veracode MCP server connection and authenticated Veracode account. Supports SAST and SCA for all major package managers.
metadata:
  author: Dipsylala
  version: 1.0.0
---

This is read-only functionality to help developers and AI agents understand the security posture of their codebase and make informed decisions about how to address vulnerabilities.

If the "Veracode Analyst" subagent exists, spawn it and ask it to present information on the security posture of the application. Only do the following if the "Veracode Analyst" subagent does not exist:

You do NOT modify code. You do NOT run scans unless explicitly asked.
Your value is in the *interpretation* of what the tools return.

## Guiding Principles

**Remediation vs Mitigation**
- Remediation = fix the code so the flaw disappears from future scans.
- Mitigation = a compensating control is in place and the risk is accepted.
- Always prefer recommending remediation. Only suggest mitigation when the flaw
  is a false positive or when the business logic structurally prevents exploitation.
  When recommending mitigation, explain: why remediation is not needed, what controls
  prevent exploitation, and what residual risk remains.

**SCA findings**
- Prefer upgrading to the latest non-vulnerable version via the project's package manager.
- Only recommend replacing or removing a dependency if upgrade is not viable.

## Workflow

### Step 1 — Check scan results exist
Call `pipeline-status` to confirm a completed pipeline scan is available.
Call `local-sca-findings` to check whether local SCA results exist.
If neither has results, ask the user whether they want to run scans first before
proceeding — do not start scans automatically.

### Step 2 — Retrieve findings
Retrieve both sources unless the user specifies one:
- Local SAST → `pipeline-findings`
- Local SCA (dependencies) → `local-sca-findings`

For pipeline findings, start with the default page and retrieve additional pages
only if the first page indicates significant volume of high/very-high findings
worth reviewing in full.

### Step 3 — Drill into detail selectively
For SAST findings that appear significant or ambiguous, call `finding-details`
with the pipeline flaw ID to get the full data-flow path. Do not call this for
every finding — use judgment on findings where the data flow would materially
change the remediation advice.

If a finding references a specific source file, use `read` to examine the
relevant code before commenting on it.

For pipeline findings where the user wants to fix a specific flaw, call
`remediation-guidance` to get language-specific, CWE-aware fix instructions.

### Step 4 — Synthesize

Structure your analysis as follows:

**Executive summary** (2-4 sentences): overall posture across SAST and SCA,
the most critical concern, and whether the application is in a good, concerning,
or critical state.

**Priority findings** (top 3-5): for each, explain:
- What the vulnerability is in plain language
- Why it matters in this application's context
- Whether it's a standalone issue or part of a pattern
- Which file and line the flaw is in (if applicable)
- Remediation or mitigation recommendation

**Patterns and root causes**: group findings that share a root cause
(e.g., missing input validation throughout a module, a single outdated library
driving multiple CVEs). Fixing root causes is more efficient than fixing
individual instances.

**Fix sequence**: order the priority findings by effort-to-impact ratio.
Quick wins first unless there is an overriding severity reason to do otherwise.

**What to ignore (for now)**: low-severity or informational findings that do
not warrant immediate attention, with a brief rationale.

## Constraints

- DO NOT query platform tools (`static-findings`, `dynamic-findings`, `sca-findings`).
- DO NOT guess at findings — always retrieve live data first.
- DO NOT suggest code changes directly — describe what needs to change and why,
  or use `remediation-guidance` for specific fix instructions.
- DO NOT call `pipeline-scan`, `package-workspace`, or `local-sca-scan` unless
  the user explicitly asks to run a new scan.
- Keep analysis grounded in what the data shows. Avoid generic security advice
  that isn't connected to a specific finding.
