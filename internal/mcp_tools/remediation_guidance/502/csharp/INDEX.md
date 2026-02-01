# CWE-502: Insecure Deserialization - C# / .NET

## LLM Guidance

Insecure deserialization in .NET occurs when untrusted data is deserialized using unsafe formatters like BinaryFormatter, NetDataContractSerializer, or ObjectStateFormatter, enabling remote code execution through arbitrary type instantiation. The core fix is to avoid deserializing untrusted data entirely, or use safe serializers like System.Text.Json with strict type controls.

## Key Principles

- Replace BinaryFormatter, NetDataContractSerializer, and ObjectStateFormatter with System.Text.Json or DataContractSerializer
- Validate and restrict types that can be deserialized using allow-lists, never deserialize arbitrary types from untrusted sources
- Apply input validation and integrity checks (HMAC signatures) before deserialization to ensure data authenticity
- Isolate deserialization in sandboxed environments with minimal privileges if unsafe formatters cannot be removed

## Remediation Steps

- Identify all deserialization points using unsafe formatters (BinaryFormatter, NetDataContractSerializer, SoapFormatter, ObjectStateFormatter)
- Replace with System.Text.Json for JSON or DataContractSerializer for XML with known types configured
- Implement type allow-lists using SerializationBinder for legacy formatters that cannot be immediately removed
- Add HMAC-based integrity validation to verify data has not been tampered with before deserialization
- Run static analysis tools to detect remaining unsafe deserialization usage

## Safe Pattern

```csharp
using System.Text.Json;

public class UserData 
{
    public string Name { get; set; }
    public int Age { get; set; }
}

public UserData SafeDeserialize(string jsonInput) 
{
    var options = new JsonSerializerOptions { PropertyNameCaseInsensitive = true };
    return JsonSerializer.Deserialize<UserData>(jsonInput, options);
}
```
