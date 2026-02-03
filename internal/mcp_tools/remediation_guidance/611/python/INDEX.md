# CWE-611: XML External Entity (XXE) Injection - Python

## LLM Guidance

Python XML parsing is risky for untrusted input because some parsers allow DTDs, entities, or entity expansion. Use `defusedxml` or configure parsers securely to prevent file disclosure, SSRF, and denial of service attacks.

**Primary Defence:** Use `defusedxml` library as a drop-in replacement for standard XML parsers, which disables all dangerous features by default.

## Key Principles

- Replace standard library parsers (`xml.etree`, `xml.dom`, `xml.sax`, `lxml`) with `defusedxml` equivalents
- Disable DTD processing, external entities, and entity expansion in all XML parsers
- Validate and sanitize XML input before parsing when possible
- Avoid parsing untrusted XML with default configurations
- Use secure parser configurations as defence-in-depth even with trusted sources

## Remediation Steps

- Install `defusedxml` - `pip install defusedxml`
- Replace imports - `from defusedxml.ElementTree import parse` instead of `from xml.etree.ElementTree import parse`
- For `lxml`, use `defusedxml.lxml` wrapper or configure parser with `resolve_entities=False`
- If `defusedxml` cannot be used, explicitly disable dangerous features - set `forbid_dtd=True`, `forbid_entities=True`, `forbid_external=True`
- Review all XML parsing code paths for unsafe configurations
- Add security tests with XXE payloads to validate protections

## Safe Pattern

```python
from defusedxml.ElementTree import parse

# Safe parsing with defusedxml
def parse_xml_safely(xml_file):
    tree = parse(xml_file)
    root = tree.getroot()
    return root

# Or configure ElementTree manually
from xml.etree.ElementTree import XMLParser, parse
parser = XMLParser()
parser.entity = {}  # Disable entities
parser.parser.SetParamEntityParsing(0)  # Disable external entities
tree = parse(xml_file, parser=parser)
```
