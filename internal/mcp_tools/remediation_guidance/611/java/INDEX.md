# CWE-611: XML External Entity (XXE) Injection - Java

## LLM Guidance

XXE vulnerabilities occur when XML parsers process external entity references in untrusted XML, allowing attackers to read files, perform SSRF attacks, or cause denial of service. Java's XML parsers (DocumentBuilder, SAXParser, XMLReader) are vulnerable by default and must be explicitly configured securely. The core fix is to disable DTDs and external entity processing entirely.

## Key Principles

- Disable DTDs completely using `FEATURE_SECURE_PROCESSING` and disallow DOCTYPE declarations
- Set all external entity features to `false` on parser factories
- Use secure parser configurations consistently across all XML processing code
- Consider using simpler data formats like JSON when XML features aren't required

## Remediation Steps

- Enable secure processing - `factory.setFeature(XMLConstants.FEATURE_SECURE_PROCESSING, true)`
- Disable DTDs - `factory.setFeature("http -//apache.org/xml/features/disallow-doctype-decl", true)`
- Disable external general entities - `factory.setFeature("http -//xml.org/sax/features/external-general-entities", false)`
- Disable external parameter entities - `factory.setFeature("http -//xml.org/sax/features/external-parameter-entities", false)`
- Disable external DTDs - `factory.setFeature("http -//apache.org/xml/features/nonvalidating/load-external-dtd", false)`
- Apply these settings to all DocumentBuilderFactory, SAXParserFactory, and XMLInputFactory instances

## Safe Pattern

```java
DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
factory.setFeature(XMLConstants.FEATURE_SECURE_PROCESSING, true);
factory.setFeature("http://apache.org/xml/features/disallow-doctype-decl", true);
factory.setFeature("http://xml.org/sax/features/external-general-entities", false);
factory.setFeature("http://xml.org/sax/features/external-parameter-entities", false);
factory.setFeature("http://apache.org/xml/features/nonvalidating/load-external-dtd", false);
factory.setXIncludeAware(false);
factory.setExpandEntityReferences(false);

DocumentBuilder builder = factory.newDocumentBuilder();
Document doc = builder.parse(inputStream);
```
