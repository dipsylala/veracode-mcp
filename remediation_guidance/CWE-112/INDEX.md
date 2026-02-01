# CWE-112: Missing XML Validation

## LLM Guidance

Missing XML validation occurs when applications parse XML without validating it against a defined schema (XSD, DTD, or RelaxNG), allowing malformed, malicious, or unexpected XML structures to be processed. This can lead to injection attacks, denial of service, or business logic bypasses. Never process untrusted XML without strict schema validation.

## Key Principles

- Always validate untrusted XML against a strict, application-defined schema before processing
- Reject any XML that does not conform exactly to the expected structure
- Use XSD types and constraints to enforce data validation at the schema level
- Limit element cardinality and data sizes to prevent denial of service attacks
- Employ allowlists for attribute values and element content where possible

## Remediation Steps

- Define a comprehensive XSD that specifies all allowed elements, types, and requirements
- Set strict data type constraints using XSD types (string, int, date, etc.)
- Enforce length limits with `maxLength` attributes to prevent oversized payloads
- Define cardinality constraints using `minOccurs` and `maxOccurs` to limit element repetition
- Restrict attribute values with enumerations or regex patterns to allowlist valid inputs
- Configure your XML parser to validate against the schema and reject invalid documents before processing
