# CWE-316: Cleartext Storage of Sensitive Information in Memory - C\#

## LLM Guidance

Storing sensitive data (passwords, keys, tokens) as cleartext strings in C# memory exposes them to memory dumps, debuggers, and disclosure attacks. Immutable strings persist until garbage collection, creating extended exposure windows. Use `char[]` for passwords with explicit `Array.Clear()` in `finally` blocks, or `SecureString` for Windows-only scenarios.

## Remediation Strategy

- Replace `string` with mutable `char[]` or `byte[]` for sensitive data to enable immediate zeroing
- Clear sensitive arrays explicitly in `finally` blocks using `Array.Clear()`
- Minimize sensitive data lifetime in memory—clear immediately after use
- Use `SecureString` for Windows-specific password handling with pinning
- Avoid serialization, logging, or concatenation of sensitive data

## Remediation Steps

- Identify all sensitive data stored as `string` (passwords, keys, tokens)
- Replace with `char[]` or `byte[]` and refactor dependent code
- Wrap usage in `try/finally` with `Array.Clear()` in the `finally` block
- For unmanaged resources, use `SafeHandle` or `Marshal.ZeroFreeBSTR()`
- Review logging, exception messages, and serialization for leakage
- Test with memory profilers to verify clearance

## Minimal Safe Pattern

```csharp
char[] password = null;
try
{
    password = GetPasswordAsCharArray();
    // Use password (e.g., authentication)
    AuthenticateUser(password);
}
finally
{
    if (password != null)
    {
        Array.Clear(password, 0, password.Length);
    }
}
```
