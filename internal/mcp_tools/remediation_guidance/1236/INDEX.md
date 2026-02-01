# CWE-1236: Formula Injection

## LLM Guidance

Formula Injection (also known as CSV Injection or Excel Injection) occurs when untrusted data containing formula metacharacters (=, +, -, @, tab, carriage return) is exported to spreadsheet files (CSV, Excel, etc.) without proper sanitization. Spreadsheet applications interpret these characters as formula directives, executing embedded commands.

## Remediation Strategy

- Treat all spreadsheet exports as potential code execution vectors requiring input sanitization
- Neutralize formula metacharacters (=, +, -, @) at export time before writing to cells
- Apply defense-in-depth by combining prefix detection, sanitization, and CSV-safe encoding
- Validate all untrusted data sources (user input, databases, external files) before export
- Use established libraries that handle formula injection protection automatically

## Remediation Steps

- Identify all spreadsheet export functionality and trace untrusted data flows to export points
- Prepend single quote (') to any cell value starting with =, +, -, @ to force literal interpretation
- Remove or escape tab and carriage return characters from all cell content
- Implement CSV-safe encoding libraries designed to prevent formula injection
- Test exports by opening files in Excel/LibreOffice to verify formulas don't execute
- Add automated tests that attempt injection with malicious formulas to validate protection
