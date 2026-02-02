# CWE-90: LDAP Injection - Java

## LLM Guidance

LDAP Injection occurs when untrusted data is used to construct LDAP queries without proper encoding, allowing attackers to manipulate directory searches and access unauthorized data.

**Primary Defence:** Use OWASP ESAPI's `encodeForLDAP()` and `encodeForDN()` methods, or Spring LDAP's query builder which provides automatic encoding.

## Key Principles

- Always encode user input before inserting into LDAP queries or DN strings
- Use parameterized query builders (Spring LDAP) instead of string concatenation
- Validate input against allowlists for expected characters and patterns
- Apply principle of least privilege to LDAP service accounts
- Sanitize special LDAP characters: `*()|\&` for filters, `,=+<>#;` for DNs

## Remediation Steps

- Replace string concatenation with OWASP ESAPI encoding methods
- For search filters - wrap user input with `Encoder.encodeForLDAP(input)`
- For DN construction - wrap user input with `Encoder.encodeForDN(input)`
- Alternatively, migrate to Spring LDAP's `LdapQueryBuilder` for automatic protection
- Add input validation to reject unexpected characters before encoding
- Test with payloads like `*)(uid=*))(|(uid=*` to verify protection

## Safe Pattern

```java
import org.owasp.esapi.ESAPI;
import org.owasp.esapi.Encoder;

public String searchUser(String username) {
    Encoder encoder = ESAPI.encoder();
    String safeName = encoder.encodeForLDAP(username);
    String filter = "(uid=" + safeName + ")";
    
    // Use filter in LDAP search
    NamingEnumeration results = ctx.search("ou=users,dc=example,dc=com", 
                                           filter, searchControls);
    return results;
}
```
