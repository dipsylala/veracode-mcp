# Agentic Remediation Loop

This document describes the design and rationale behind the autonomous Veracode
scan-triage-fix loop implemented as a pair of VS Code Copilot agent files.

---

## Overview

The loop automates the full identify-and-fix cycle for high/critical security
vulnerabilities in a workspace:

1. **Package and synchronously scan** (SAST pipeline scan + SCA scan)
2. **Triage** — retrieve findings and rank them by remediation priority
3. **Fix** — apply smallest-safe fixes for each high/critical finding
4. **Checkpoint** — if more than 10 files have been changed, rescan to confirm
   fixes and continue with any remaining issues
5. **Repeat** until all high/critical findings are resolved

---

## New Agent Files

| File | Name | Purpose |
| --- | --- | --- |
| `.github/agents/veracode-triage.agent.md` | Veracode Triage | Returns a machine-readable JSON array of prioritised findings for use by automated agents |
| `.github/agents/veracode-autofix.agent.md` | Veracode Autofix | Orchestrates the full scan → triage → fix loop autonomously |

---

## Veracode Triage agent

**Derived from**: `veracode-analyst.agent.md`

The Triage agent uses the same data-retrieval workflow as the Analyst (Steps 1–3:
check scan results, retrieve findings, selective drill-in via `finding-details`).

The difference is in the output: instead of human-readable prose, it emits a
machine-readable JSON array sorted by `remediationPriority` (1 = fix first).
SAST and SCA findings are ranked together in a single unified list.

### JSON schema per finding

```json
{
  "id": "<pipeline flaw ID or CVE/SID>",
  "type": "sast | sca",
  "severity": 0,
  "cwe": 89,
  "title": "<flaw title>",
  "file": "<workspace-relative path or null>",
  "line": 42,
  "component": "<name@version for SCA, null for SAST>",
  "cve": "<CVE/SID for SCA, null for SAST>",
  "remediationPriority": 1,
  "remediationRationale": "<one sentence>"
}
```

**Ranking rules:**

- Severity first (5 = VeryHigh → 0 = Informational)
- SCA findings with CVSS ≥ 9.0 are bumped to VeryHigh regardless of Veracode severity
- Among equal severities: exploitability context first, then ease of fix (quick wins rank higher)

**Constraints** (same as the Analyst agent): no code changes, no scanning unless
the caller explicitly requests it.

---

## Veracode Autofix agent

The Autofix agent is the main orchestrator. It is `user-invocable: true` so it
can be started directly from the VS Code Copilot agent picker.

### Loop structure

```text
┌─────────────────────────────────────────────────────────────┐
│  Phase 1 — Scan                                             │
│    package-workspace                                        │
│    pipeline-scan (synchronous: true)                        │
│    local-sca-scan                                           │
└────────────────────────┬────────────────────────────────────┘
                         │
┌────────────────────────▼────────────────────────────────────┐
│  Phase 2 — Triage                                           │
│    Spawn "Veracode Triage" → JSON findings array            │
│    No severity ≥ 4? → STOP (clean)                         │
└────────────────────────┬────────────────────────────────────┘
                         │
┌────────────────────────▼────────────────────────────────────┐
│  Phase 3 — Fix (iterate by remediationPriority)             │
│    SAST: remediation-guidance → apply smallest safe fix     │
│    SCA:  identify safe version → upgrade via package mgr    │
│    After each fix: run build → auto-fix compile errors      │
│    Track changed files                                      │
│    > 10 files changed? → Phase 4                           │
└────────────────────────┬────────────────────────────────────┘
                         │ (all attempted or no more sev ≥ 4)
┌────────────────────────▼────────────────────────────────────┐
│  Final Report                                               │
│    Findings fixed / remaining / files changed / scans run   │
└─────────────────────────────────────────────────────────────┘

Phase 4 (10-file checkpoint):
  → Re-run Phase 1 (rescan)
  → Re-run Phase 2 (re-triage)
  → Reset changed-file counter
  → No severity ≥ 4? STOP (clean). Otherwise back to Phase 3.
```

### 10-file gate

After more than 10 files have been modified in a single Phase 3 pass, a
checkpoint rescan is triggered. This:

- Confirms that fixes so far have not introduced regressions or new findings
- Provides an updated ranked list so subsequent fixes target the current state

The file-change count is tracked by the agent (via `git diff --name-only` or by
noting files edited during the pass, whichever is more accurate for the project).

### Compilation verification

After every fix, the agent runs the project's build command. If compilation
errors are found, the agent reads the error output and applies a minimal fix
before continuing. Compilation fixes count toward the 10-file gate.

### Termination

The loop ends when:

- The most recent triage returns no findings with severity ≥ 4 (clean), or
- All severity ≥ 4 findings from the most recent triage have been attempted

The agent does **not** fix Medium/Low/Informational findings unless explicitly
asked.

---

## Design Decisions

| Decision | Choice | Rationale |
| --- | --- | --- |
| Triage output format | JSON array | Machine-readable; consumed by Autofix without parsing prose |
| Scan mode | `synchronous: true` | Autofix must wait for results before triaging |
| File-change gate | > 10 files | Balances thoroughness against scan turnaround time |
| Post-gate behaviour | Continue looping | Loop runs until all high/critical are resolved |
| Compilation errors | Auto-fix | Prevents compounding errors across subsequent fixes |
| Severity threshold | ≥ 4 (High + VeryHigh) | Focus on exploitable, high-impact issues first |
| Autofix user-invocable | `true` | Accessible directly from the VS Code agent picker |
| Triage user-invocable | `false` | Intended as a machine callout, not a user-facing agent |

---

## Existing Skills Used as Reference

| Skill | Informs |
| --- | --- |
| `skills/scanit/SKILL.md` | Phase 1 scan steps (package-workspace, pipeline-scan, local-sca-scan) |
| `skills/fixit/SKILL.md` | Phase 3 fix steps (remediation-guidance flow, SCA upgrade logic) |
| `skills/reportit/SKILL.md` | Subagent spawn pattern (references Veracode Analyst) |

---

## Verification Checklist

- [ ] Invoke `Veracode Triage` against an existing pipeline scan; confirm JSON
      output matches the schema above
- [ ] Run `Veracode Autofix` on a test workspace; confirm it packages, scans
      synchronously, triages, then applies at least one fix
- [ ] Verify the 10-file gate triggers a rescan (introduce > 10 changes and
      confirm Phase 4 fires)
- [ ] Confirm compilation self-recovery: introduce a bad fix and verify the
      agent auto-corrects before continuing
