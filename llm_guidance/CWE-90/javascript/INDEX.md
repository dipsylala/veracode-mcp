# CWE-90: LDAP Injection - JavaScript

## LLM Guidance

LDAP Injection occurs when untrusted user input is concatenated into LDAP queries without proper sanitization, allowing attackers to modify query logic. This can lead to authentication bypass, unauthorized data access, or privilege escalation in applications using LDAP for directory services. The core fix is to use parameterized queries or strict input validation with allowlisting.

## Remediation Strategy

- Use libraries that support parameterized LDAP queries or prepared filters
- Validate and sanitize all user inputs with strict allowlists before using in LDAP queries
- Escape LDAP special characters: `*`, `(`, `)`, `\`, `/`, `NUL` using proper encoding
- Implement least-privilege access controls on LDAP directory operations
- Use framework-provided LDAP query builders instead of string concatenation

## Remediation Steps

- Identify all LDAP query construction points that use user input
- Replace string concatenation with parameterized queries or safe query builders
- Apply LDAP escaping functions to all user-supplied values
- Implement input validation with allowlists for username/search patterns
- Add logging and monitoring for suspicious LDAP query patterns
- Conduct security testing with LDAP injection payloads

## Minimal Safe Pattern

```javascript
const ldap = require('ldapjs');

// Safe: Escape LDAP special characters
function escapeLDAP(str) {
  return str.replace(/[*()\\\/\x00]/g, (char) => '\\' + char.charCodeAt(0).toString(16).padStart(2, '0'));
}

const username = escapeLDAP(userInput);
const filter = `(&(objectClass=user)(uid=${username}))`;

client.search('ou=users,dc=example,dc=com', { filter, scope: 'sub' }, (err, res) => {
  // Handle results
});
```
