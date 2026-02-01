# CWE-611: XML External Entity (XXE) Injection

## LLM Guidance

XML External Entity (XXE) injection occurs when XML input containing a reference to an external entity is processed by a weakly configured XML parser. The vulnerability exists because XML parsers, by default, often resolve external entities defined in Document Type Definitions (DTDs), allowing attackers to inject malicious entity definitions that can read arbitrary files, perform Server-Side Request Forgery (SSRF) attacks, cause Denial of Service (DoS), or in rare cases execute remote code.

## Key Principles

- Disable XML external entities and DTD processing by default in all parsers
- Only enable external entity resolution if explicitly required and with strict security constraints
- Server must fully control XML parsing behavior—never trust parser defaults
- Use the most restrictive parser configuration possible for your use case
- Apply defense-in-depth: input validation combined with secure parser settings

## Remediation Steps

- Identify the vulnerable XML parsing location from security findings (file, line number, parser library)
- Trace XML data flow from input source (user input, files, network requests) to the parser
- Determine which XML parser library is in use and locate its configuration
- Configure parser to disable external entities, DTD processing, and XInclude features
- Apply parser-specific secure settings (see language-specific guidance for your library)
- Test that external entity references are blocked and rejected by the parser
