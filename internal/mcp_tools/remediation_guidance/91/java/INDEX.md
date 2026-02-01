# CWE-91: XML Injection - Java

## LLM Guidance

XML Injection in Java occurs when untrusted user input is used to construct XML documents without proper validation or escaping. Attackers inject special characters (`<`, `>`, `&`, `'`, `"`) to manipulate XML structure, potentially causing data corruption, authentication bypass, or information disclosure. **Primary Defense:** Use DOM API methods (`DocumentBuilder`, `Element.setAttribute()`, `Element.setTextContent()`) or sanitize input by escaping XML metacharacters.

## Key Principles

- Never concatenate untrusted input directly into XML strings
- Use DOM APIs (`createElement`, `setTextContent`, `setAttribute`) which auto-escape content
- Validate and sanitize all user input before XML processing
- Use XML libraries that enforce proper encoding (JAXB, DOM4J with safe configurations)
- Disable external entity processing to prevent XXE attacks

## Remediation Steps

- Replace string concatenation with DOM API methods for XML construction
- Apply XML escaping using `StringEscapeUtils.escapeXml()` for string-based approaches
- Validate input against whitelist patterns before XML processing
- Configure parsers to disable DTDs and external entities (`setFeature(XMLConstants.FEATURE_SECURE_PROCESSING, true)`)
- Use parameterized XPath queries instead of string concatenation

## Safe Pattern

```java
import org.w3c.dom.*;
import javax.xml.parsers.*;

// Safe: Using DOM API (auto-escapes)
DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
DocumentBuilder builder = factory.newDocumentBuilder();
Document doc = builder.newDocument();

Element root = doc.createElement("user");
Element name = doc.createElement("name");
name.setTextContent(userInput); // Auto-escaped, safe
root.appendChild(name);

// Alternative: Explicit escaping for string-based XML
String safe = StringEscapeUtils.escapeXml11(userInput);
String xml = "<user><name>" + safe + "</name></user>";
```
