# CWE-489: Active Debug Code

## LLM Guidance

Leftover debug code in production (print statements, test backdoors, disabled authentication, verbose error messages, debug endpoints) exposes sensitive information, creates security bypasses, and provides attackers with reconnaissance data. Core principle: Remove or strictly gate all debug paths before production deployment.

## Key Principles

- Remove before shipping: Strip debug code, test accounts, and development features from production builds
- Gate debug features: If debug functionality is needed, protect it behind authentication, authorization, and environment checks
- Disable by default: Ensure DEBUG flags, verbose logging, and development modes are off in production configurations
- Sanitize error output: Replace detailed error messages with generic responses; log details server-side only
- Use build-time removal: Leverage build tools to automatically strip debug code from production artifacts

## Remediation Steps

- Search for debug patterns - Find print/console statements, debug flags, test credentials, and debug endpoints (e.g., `/debug/*`, `/test/*`)
- Review configurations - Check for `DEBUG=true`, development settings, and verbose logging enabled in production config files
- Remove test backdoors - Eliminate test accounts, authentication bypasses, and `?debug=true` parameters
- Sanitize error handlers - Replace stack traces and verbose errors exposed to users with generic messages
- Examine data flows - Review scan data_paths to locate where debug code exists in production paths
- Implement environment checks - Wrap any necessary debug features with strict environment validation and access controls
