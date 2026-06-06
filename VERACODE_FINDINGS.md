# Veracode Pipeline Scan Findings

**Scan Status:** SUCCESS — 12 findings across 1 module (`veracode-mcp` Go source). No High or Very High severity issues. Dominant issue types are log injection (CWE-117) and error message disclosure (CWE-209), plus one path traversal risk (CWE-73).

---

## Severity Breakdown

| Severity | Count | Exploitability  |
| -------- | -----: | --------------- |
| Medium   | 5     | Neutral, Likely |
| Low      | 7     | Neutral         |

---

## Issue Type Breakdown

| Issue Type | Count | Max Severity |
| ---------- | -----: | ----------- |
| Improper Output Neutralization for Logs (CWE-117) | 4 | Medium |
| External Control of File Name or Path (CWE-73) | 1 | Medium |
| Information Exposure Through an Error Message (CWE-209) | 7 | Low |

---

## Findings by File

### `internal/transport/stdio.go` — 7 findings (highest: Medium)

| ID | Severity | Issue Type | CWE | Line | Mitigations |
| -- | -------- | ---------- | --- | ---: | ----------- |
| 1007 | Medium   | Improper Output Neutralization for Logs | 117 | 94  | handleRequest() is only called after isValidMethod() passes. If the method contains any non-alphanumeric/slash character (including \r\n), the request is rejected and handleRequest() is never invoked. The log at line 94 only ever sees a validated method name |
| 1004 | Low      | Information Exposure Through an Error Message | 209 | 41  | Error detail is written to the operator log (`stderr` or log file) and never included in the JSON-RPC response returned to the client |
| 1001 | Low      | Information Exposure Through an Error Message | 209 | 54  | Error detail is written to the operator log and never included in the JSON-RPC response returned to the client |
| 1011 | Low      | Information Exposure Through an Error Message | 209 | 112 | Error detail is written to the operator log and never included in the JSON-RPC response returned to the client |
| 1010 | Low      | Information Exposure Through an Error Message | 209 | 130 | Error detail is written to the operator log and never included in the JSON-RPC response returned to the client |
| 1009 | Low      | Information Exposure Through an Error Message | 209 | 152 | Error detail is written to the operator log and never included in the JSON-RPC response returned to the client |
| 1002 | Low      | Information Exposure Through an Error Message | 209 | 159 | Error detail is written to the operator log and never included in the JSON-RPC response returned to the client |

### `internal/server/server.go` — 4 findings (highest: Medium)

| ID | Severity | Issue Type | CWE | Line | Mitigations |
| -- | -------- | ---------- | --- | ---: | ----------- |
| 1003 | Medium   | Improper Output Neutralization for Logs | 117 | 199 | `ValidateMethod` now runs first; this log (notification path) only executes after method is confirmed valid |
| 1008 | Medium   | Improper Output Neutralization for Logs | 117 | 203 | `ValidateMethod` now runs first; this log (missing-ID path) only executes after method is confirmed valid |
| 1005 | Medium   | Improper Output Neutralization for Logs | 117 | 253 | `ValidateMethod` now runs first; this log only executes after method passes validation |
| 1006 | Low      | Information Exposure Through an Error Message | 209 | 242 | Error detail is written to the operator log and never included in the JSON-RPC response returned to the client |

### `internal/cli/app.go` — 1 finding (highest: Medium)

| ID | Severity | Issue Type | CWE | Line | Exploitability |
| -- | -------- | ---------- | --- | ---: | -------------- |
| 1000 | Medium   | External Control of File Name or Path | 73 | 31 | logFilePath is a CLI flag — the user running the server intentionally controls where logs are written. A // #nosec G304 comment is already in place explicitly acknowledging this. |
