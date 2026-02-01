# CWE-103: Struts: Incomplete validate() Method Definition

## LLM Guidance

Struts ActionForm's `validate()` method that returns null or is improperly implemented bypasses validation, allowing unvalidated user input to reach application logic. This enables injection attacks, data integrity issues, and business logic bypasses.

## Key Principles

- All ActionForm classes must implement complete `validate()` methods that return ActionErrors (never null)
- Every form field accepting user input requires explicit validation rules
- Security requirements and validations must be fully documented and enforced
- Use Struts validator framework (`ValidatorForm`, `DynaValidatorForm`) for comprehensive validation
- Trace data flows to ensure no unvalidated input reaches business logic

## Remediation Steps

- Review security findings to identify ActionForm classes with incomplete or null-returning `validate()` methods
- Check for empty implementations, stub code, or missing field validations
- Extend proper validation classes (`ValidatorForm` or `DynaValidatorForm`) instead of basic `ActionForm`
- Implement complete `validate()` method that validates all form fields and returns ActionErrors object
- Configure validation rules in validation.xml or use annotations for declarative validation
- Test all input paths to verify validation is enforced before data reaches business logic
