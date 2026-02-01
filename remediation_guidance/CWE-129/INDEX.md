# CWE-129: Improper Validation of Array Index

## LLM Guidance

Improper validation of array index occurs when user-controlled or untrusted data is used as an array index without proper bounds checking, allowing attackers to read or write arbitrary memory locations, leak sensitive data, or cause crashes. The core fix is to validate array indices before use, including checks for negative and overflowed values.

## Key Principles

- Validate all array indices before use, checking both negative and overflowed values
- Always verify both lower bound (>= 0) and upper bound (< array.length)
- Use unsigned types for indices to prevent negative value exploits
- Reject invalid indices immediately rather than attempting correction
- Trace data flow from untrusted sources to array access points

## Remediation Steps

- Review security findings to identify where untrusted data is used as an array index
- Locate the array access and identify the index source (user input, external file, database, network request)
- Trace the data flow from source to array access to understand the complete path
- Implement bounds checking - ensure `index >= 0 && index < array.length` before every access
- Use unsigned types (`size_t` in C/C++, `unsigned int`) for indices to prevent negative values
- Reject invalid indices with appropriate error handling rather than clamping or wrapping
