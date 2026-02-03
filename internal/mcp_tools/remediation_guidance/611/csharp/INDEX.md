# CWE-611: XML External Entity (XXE) Injection - C# / .NET

## LLM Guidance

XXE vulnerabilities in .NET occur when XML parsers process external entity references in untrusted XML, enabling file disclosure, SSRF, and DoS attacks. Modern .NET versions (.NET Core 3.0+) have safe defaults, but .NET Framework and misconfigured parsers remain vulnerable.

**Primary Defence:** Set `DtdProcessing = DtdProcessing.Prohibit` and `XmlResolver = null` in `XmlReaderSettings`.

## Key Principles

- Prohibit DTD processing entirely using `DtdProcessing.Prohibit` to reject DOCTYPE declarations
- Set `XmlResolver = null` to block external entity resolution even if DTDs bypass other controls
- Apply secure settings to all XML parsing APIs (XmlDocument, XmlReader, XmlSerializer, DataContractSerializer)
- Use `MaxCharactersFromEntities = 0` to prevent entity expansion DoS attacks
- Create reusable secure configuration helpers to ensure consistent protection across the codebase

## Remediation Steps

- Identify all XML parsing locations (XmlDocument.LoadXml, XmlReader.Create, XmlSerializer.Deserialize, etc.)
- Create `XmlReaderSettings` with `DtdProcessing = DtdProcessing.Prohibit` and `XmlResolver = null`
- Replace direct XML parsing calls with secure XmlReader-wrapped patterns
- For XmlDocument, explicitly set `doc.XmlResolver = null` even with secure reader
- Test with XXE payloads (<!DOCTYPE, SYSTEM entities) to verify rejection
- Validate legitimate XML workflows still function correctly after hardening

## Safe Pattern

```csharp
using System.Xml;

public void ParseXmlSecure(string xml)
{
    var settings = new XmlReaderSettings
    {
        DtdProcessing = DtdProcessing.Prohibit,
        XmlResolver = null,
        MaxCharactersFromEntities = 0
    };
    
    using (var reader = XmlReader.Create(new StringReader(xml), settings))
    {
        var doc = new XmlDocument { XmlResolver = null };
        doc.Load(reader);
        // Process safely
    }
}
```
