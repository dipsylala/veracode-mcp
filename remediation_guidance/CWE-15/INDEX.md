# CWE-15: External Control of System or Configuration Setting

## LLM Guidance

This vulnerability occurs when user input controls system or application configuration settings, allowing attackers to alter application behavior, security controls, or environment variables. The core fix is to never allow untrusted input to directly control configuration—all settings must be defined and enforced by trusted code.

## Key Remediation Principles

- Define all configuration through trusted deployment mechanisms, not runtime user input
- Use allowlists to constrain configuration values to known-safe options
- Separate user preferences from security-critical system configuration
- Validate and sanitize any user data that influences application behavior
- Enforce configuration integrity through code-based defaults

## Remediation Steps

- Review scan results for file path, line number, and variables where user input controls configuration
- Trace data flow from untrusted sources (HTTP params, headers, cookies) to configuration sinks
- Replace direct user input with predefined configuration options selected via allowlist
- Move security-sensitive settings to deployment configuration files or environment variables
- Implement validation layers that reject unexpected configuration values
- Conduct code review to identify additional instances of externally-controlled settings
