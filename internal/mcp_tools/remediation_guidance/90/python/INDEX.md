# CWE-90: LDAP Injection - Python

## LLM Guidance

LDAP Injection occurs when untrusted user input is concatenated into LDAP queries without proper sanitization, allowing attackers to manipulate queries to bypass authentication, escalate privileges, or extract sensitive directory data. The primary defense is using parameterized queries with the `ldap3` library and escaping LDAP special characters (`*`, `(`, `)`, `\`, `NUL`) in user-controlled input. Never construct LDAP filter strings through concatenation.

## Remediation Strategy

- Use `ldap3` library with parameterized queries instead of string concatenation
- Escape all special LDAP characters in user input using `ldap3.utils.conv.escape_filter_chars()`
- Apply allowlist validation on user input before query construction
- Implement least-privilege access for LDAP service accounts
- Use DN (Distinguished Name) sanitization for attribute values

## Remediation Steps

- Replace string concatenation with parameterized filter construction
- Apply `escape_filter_chars()` to all user-controlled variables in LDAP filters
- Validate input against expected patterns (e.g., alphanumeric usernames)
- Review LDAP query logging to detect injection attempts
- Test filters with malicious payloads like `*)(objectClass=*)` and `admin)(&(password=*)`
- Restrict LDAP bind account permissions to minimum required scope

## Minimal Safe Pattern

```python
from ldap3 import Server, Connection, ALL
from ldap3.utils.conv import escape_filter_chars

server = Server('ldap://example.com', get_info=ALL)
conn = Connection(server, user='cn=admin,dc=example,dc=com', password='password')

# Safe: escape user input
username = escape_filter_chars(user_input)
search_filter = f"(&(objectClass=person)(uid={username}))"

conn.search('dc=example,dc=com', search_filter)
```
