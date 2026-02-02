# CWE-90: LDAP Injection - C# / .NET

## LLM Guidance

LDAP Injection occurs when untrusted data is used to construct LDAP queries without proper encoding, allowing attackers to manipulate LDAP searches and access unauthorized data. The core fix involves strict allowlist validation of input (e.g., alphanumeric usernames only) and escaping LDAP special characters (`*`, `(`, `)`, `\`, `/`, `NUL`). Never construct Distinguished Names (DNs) directly from user input—instead, search by attribute and use the returned DN for subsequent operations.

## Key Principles

- Validate with strict allowlists - Restrict input to expected patterns before any LDAP operations
- Escape LDAP metacharacters - Encode `*()\/` and null bytes when user input must appear in filters
- Search-then-use pattern - Query by safe attribute, retrieve the object's DN, use that DN for further operations
- Avoid DN construction - Never concatenate user input into Distinguished Names or filter strings
- Principle of least privilege - Use service accounts with minimal LDAP permissions

## Remediation Steps

- Apply regex allowlist validation to usernames/inputs - `^[a-zA-Z0-9._-]{3,64}$`
- Escape LDAP special characters if validation isn't feasible - `*`, `(`, `)`, `\`, `/`, and null bytes
- Use `DirectorySearcher.Filter` with escaped values instead of string concatenation
- For authentication, search by `sAMAccountName`, retrieve the object, then use its `Path` property
- Never build LDAP filter strings with `String.Format()` or interpolation on raw user input
- Test with injection payloads - `*)(objectClass=*)`, `admin*`, `*)(uid=*`

## Safe Pattern

```csharp
string EscapeLdap(string input) {
    return input.Replace("\\", "\\5c").Replace("*", "\\2a")
                .Replace("(", "\\28").Replace(")", "\\29")
                .Replace("\0", "\\00").Replace("/", "\\2f");
}

string username = EscapeLdap(userInput);
using (var searcher = new DirectorySearcher()) {
    searcher.Filter = $"(&(objectClass=user)(sAMAccountName={username}))";
    SearchResult result = searcher.FindOne();
    if (result != null) {
        string userDn = result.Path; // Use this DN, don't construct it
    }
}
```
