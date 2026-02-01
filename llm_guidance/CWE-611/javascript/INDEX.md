# CWE-611: XML External Entity (XXE) Injection - JavaScript/Node.js

## LLM Guidance

XXE injection occurs when XML parsers process external entity references, allowing attackers to read local files, perform SSRF attacks, or cause denial of service. Node.js applications parsing XML from SOAP APIs, SVG uploads, RSS feeds, or configuration files are vulnerable when using libraries like `libxmljs`, `xml2js`, `fast-xml-parser` without disabling external entities. The core fix is to disable external entity resolution and DTD processing in parser configurations.

## Remediation Strategy

- Disable external entity processing and DTD resolution in all XML parser configurations
- Use secure XML parsing libraries with XXE protection enabled by default
- Validate and sanitize XML input before parsing, rejecting suspicious patterns
- Prefer JSON over XML for data exchange when architecturally feasible

## Remediation Steps

- Configure `libxmljs` with `noent - false`, `nonet - true`, `dtdload - false` options
- For `xml2js` v0.5.0+, verify external entities are disabled (default behavior)
- Set `fast-xml-parser` with `processEntities - false` option
- Review all XML parsing code and apply secure configurations consistently
- Add input validation to reject XML containing DOCTYPE declarations or entity references

## Minimal Safe Pattern

```javascript
// libxmljs with XXE protection
const libxmljs = require('libxmljs');

const parseXMLSafely = (xmlString) => {
  return libxmljs.parseXml(xmlString, {
    noent: false,    // Disable entity substitution
    nonet: true,     // Disable network access
    dtdload: false,  // Disable DTD loading
    dtdvalid: false  // Disable DTD validation
  });
};
```
