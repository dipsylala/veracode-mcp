# CWE-91: XML Injection - Python

## LLM Guidance

XML Injection occurs when untrusted input is used to construct XML documents without proper validation or escaping, allowing attackers to manipulate XML structure using special characters (`<`, `>`, `&`, `'`, `"`). This can lead to data corruption, authentication bypass, or information disclosure. **Primary Defense:** Use `xml.etree.ElementTree` or `lxml` with proper element creation methods instead of string concatenation, and validate/escape all user input.

## Remediation Strategy

- Never construct XML documents using string concatenation or formatting with user input
- Use XML library APIs that handle escaping automatically (ElementTree, lxml)
- Validate and sanitize all user input before incorporating into XML structures
- Apply whitelist validation for element names, attributes, and structural components
- Disable external entity processing and DTD validation to prevent XXE attacks

## Remediation Steps

- Replace string concatenation with `ElementTree.Element()` and `SubElement()` methods
- Use `.text` and `.set()` properties to assign content and attributes safely
- Implement input validation using regex or allowlists for acceptable characters
- Configure parsers with `defusedxml` library to disable dangerous features
- Escape special XML characters using `xml.sax.saxutils.escape()` if manual construction is unavoidable
- Review all XML generation code paths for user input handling

## Minimal Safe Pattern

```python
from xml.etree.ElementTree import Element, SubElement, tostring

def create_user_xml(user_name, user_role):
    # Safe: Using ElementTree API
    root = Element('user')
    name_elem = SubElement(root, 'name')
    name_elem.text = user_name  # Automatically escaped
    role_elem = SubElement(root, 'role')
    role_elem.set('type', user_role)  # Automatically escaped
    return tostring(root, encoding='unicode')

# Unsafe alternative to avoid:
# xml = f"<user><name>{user_name}</name></user>"  # VULNERABLE
```
