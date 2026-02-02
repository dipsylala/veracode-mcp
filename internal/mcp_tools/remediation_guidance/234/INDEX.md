# CWE-234: Failure to Handle Missing Parameter

## LLM Guidance

Applications fail to validate that required input parameters are present before use, causing null pointer exceptions, logic errors, security bypasses (missing authentication tokens), or crashes. The fix requires explicit validation that required parameters exist before accessing them, with proper error handling for missing values.

## Key Principles

- Validate parameter existence: Check all required parameters are present before accessing their values
- Fail securely: Reject requests with missing required parameters rather than using unsafe defaults
- Validate authentication/authorization params: Never proceed with security operations when tokens or credentials are missing
- Use framework validation: Leverage built-in parameter validation in frameworks (schema validators, required field annotations)
- Handle optionals explicitly: Distinguish between required and optional parameters with clear default behaviors

## Remediation Steps

- Identify unguarded parameter access - Find code that accesses request/function parameters without null/existence checks
- Add validation gates - Insert checks that verify required parameters exist before use (throw errors if missing)
- Review authentication flows - Ensure security-critical parameters (tokens, user IDs, permissions) are validated as present
- Replace unsafe defaults - Remove code that assigns default values to missing required parameters
- Use schema validation - Implement request schema validators that enforce required parameters at API boundaries
- Test missing parameter cases - Add test cases that send requests with missing parameters to verify proper rejection
