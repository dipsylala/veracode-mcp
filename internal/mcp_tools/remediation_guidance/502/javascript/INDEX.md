# CWE-502: Insecure Deserialization - JavaScript / Node.js

## LLM Guidance

JavaScript deserialization vulnerabilities occur when `eval()`, `Function()`, `vm.runInNewContext()`, or vulnerable libraries (node-serialize, serialize-javascript) parse untrusted data, allowing attackers to execute arbitrary code. Node.js applications are particularly vulnerable when deserializing from cookies, external APIs, or user uploads.

**Primary Defense:** Use `JSON.parse()` exclusively for deserialization and validate input against strict schemas.

## Key Principles

- Replace `eval()`, `Function()`, and `vm` module usage with `JSON.parse()` for all data deserialization
- Validate deserialized data with schema validation libraries (Joi, Ajv, Zod) before use
- Remove dependencies on insecure libraries like node-serialize, serialize-javascript, or funcster
- Implement allowlists for expected object types and reject unexpected properties
- Use Content Security Policy and strict input validation at API boundaries

## Remediation Steps

- Audit codebase for `eval()`, `Function()`, `vm.runInNewContext()`, and unsafe deserialization libraries
- Replace all unsafe deserialization with `JSON.parse()` and add try-catch error handling
- Implement JSON schema validation immediately after deserialization
- Add integrity checks (HMAC signatures) to serialized data from untrusted sources
- Configure CSP headers to prevent inline script execution
- Test with malicious payloads to verify protections

## Safe Pattern

```javascript
const Ajv = require('ajv');
const ajv = new Ajv();

const schema = {
  type: 'object',
  properties: { name: { type: 'string' }, age: { type: 'number' } },
  required: ['name', 'age'],
  additionalProperties: false
};

function safeDeserialize(jsonString) {
  const data = JSON.parse(jsonString); // Safe parsing
  if (!ajv.validate(schema, data)) throw new Error('Invalid data');
  return data;
}
```
