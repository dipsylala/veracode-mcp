# CWE-611: XML External Entity (XXE) Injection - PHP

## LLM Guidance

PHP's XML parsers can process external entities by default, leading to file disclosure, SSRF attacks, and denial of service. Proper configuration is essential to prevent XXE vulnerabilities.

**Primary Defence:** Call `libxml_disable_entity_loader(true)` before parsing XML (PHP < 8.0), or set `LIBXML_NOENT` to false and avoid `LIBXML_DTDLOAD` for all XML functions.

## Key Principles

- Disable external entity loading globally before any XML parsing operations
- Use parser options that explicitly prevent entity expansion and DTD loading
- Validate and sanitize XML input to reject documents containing entity declarations
- Prefer JSON over XML when possible to eliminate XXE risk entirely
- Keep PHP updated (8.0+ has safer defaults with entity loader disabled by default)

## Remediation Steps

- Call `libxml_disable_entity_loader(true)` at application initialization for PHP < 8.0
- Remove `LIBXML_NOENT` flag from all `simplexml_load_*`, `DOMDocument - -load*`, and `XMLReader` calls
- Explicitly pass `LIBXML_NONET` to prevent network access during parsing
- Never use `LIBXML_DTDLOAD` unless absolutely required with trusted input only
- Test with payloads containing `<!ENTITY>` declarations to verify protection
- Review all usages of `simplexml_*`, `DOMDocument`, `XMLReader`, and `xml_parse` functions

## Safe Pattern

```php
<?php
// PHP < 8.0: Disable entity loading globally
libxml_disable_entity_loader(true);

// Safe XML parsing with DOMDocument
$dom = new DOMDocument();
$dom->loadXML($xmlString, LIBXML_NONET);

// Safe XML parsing with SimpleXML
$xml = simplexml_load_string($xmlString, 'SimpleXMLElement', LIBXML_NONET);

// Safe XML parsing with XMLReader
$reader = new XMLReader();
$reader->XML($xmlString, null, LIBXML_NONET);
?>
```
