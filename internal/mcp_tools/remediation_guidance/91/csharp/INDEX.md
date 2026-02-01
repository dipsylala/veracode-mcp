# CWE-91: XML Injection - C\#

## LLM Guidance

XML Injection occurs when untrusted user input is embedded into XML documents without proper validation or escaping, allowing attackers to manipulate XML structure using special characters (`<`, `>`, `&`, `'`, `"`). This can lead to data corruption, authentication bypass, or information disclosure. Use LINQ to XML APIs (`XElement`, `XAttribute`) instead of string concatenation, and validate/escape all user input with `SecurityElement.Escape()`.

## Key Principles

- Never concatenate user input directly into XML strings
- Use LINQ to XML constructors that automatically escape content
- Apply `SecurityElement.Escape()` when string manipulation is unavoidable
- Validate input against strict allowlists before XML construction
- Parse and reconstruct XML rather than modifying as strings

## Remediation Steps

- Replace string concatenation with `XElement`/`XAttribute` constructors
- Wrap untrusted data with `SecurityElement.Escape()` if using `XmlWriter` or string-based methods
- Implement input validation using allowlists for expected characters/patterns
- Use parameterized XML construction methods consistently
- Test with XML metacharacters (`<test>`, `&payload;`, `"value"`) to verify escaping
- Review existing code for `.ToString()` concatenation patterns

## Safe Pattern

```csharp
// Safe: LINQ to XML automatically escapes content
string userInput = GetUserInput();
XElement root = new XElement("user",
    new XElement("name", userInput),
    new XAttribute("id", userId)
);

// Safe: Manual escaping when necessary
string escapedInput = SecurityElement.Escape(userInput);
string xml = $"<user>{escapedInput}</user>";
```
