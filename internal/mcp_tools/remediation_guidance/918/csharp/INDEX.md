# CWE-918: Server-Side Request Forgery (SSRF) - C\#

## LLM Guidance

Server-Side Request Forgery (SSRF) allows attackers to make the server perform HTTP requests to arbitrary destinations, potentially accessing internal services, cloud metadata endpoints (169.254.169.254), or bypassing firewalls. The primary defense is validating URLs against an allowlist of permitted domains/IPs, blocking private IP ranges (10.x, 172.16-31.x, 192.168.x, 127.x), and using `AllowAutoRedirect = false` to prevent redirect-based bypasses.

## Key Principles

- Validate all URLs against an allowlist of permitted domains before making requests
- Block private IP ranges (RFC 1918: 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16), loopback (127.0.0.0/8), and link-local addresses (169.254.0.0/16)
- Disable automatic redirects with `AllowAutoRedirect = false` to prevent redirect-based SSRF bypasses
- Resolve DNS and validate resulting IP addresses to prevent DNS rebinding attacks
- Enforce HTTPS-only and implement request timeouts to prevent DoS

## Remediation Steps

- Create a URL validator that checks URLs against an allowlist of allowed domains/schemes
- Implement `IsPrivateIp()` checks covering RFC 1918 ranges, loopback, and link-local addresses
- Use `Dns.GetHostAddresses()` to resolve and validate IPs after initial URL validation
- Configure HttpClient with `AllowAutoRedirect = false` and `UseProxy = false`
- Block cloud metadata endpoints (169.254.169.254, metadata.google.internal) explicitly
- Return generic error messages to prevent information disclosure during validation failures

## Safe Pattern

```csharp
private static readonly HashSet<string> AllowedHosts = new() 
    { "api.example.com", "cdn.example.com" };

private Uri ValidateUrl(string url)
{
    if (!Uri.TryCreate(url, UriKind.Absolute, out Uri? uri) || 
        uri.Scheme != "https" || !AllowedHosts.Contains(uri.Host.ToLowerInvariant()))
        throw new SecurityException("Invalid URL");
    
    var addresses = Dns.GetHostAddresses(uri.Host);
    foreach (var addr in addresses)
    {
        if (IPAddress.IsLoopback(addr)) throw new SecurityException("Private IP blocked");
        var bytes = addr.GetAddressBytes();
        if (bytes[0] == 10 || (bytes[0] == 172 && bytes[1] >= 16 && bytes[1] <= 31) ||
            (bytes[0] == 192 && bytes[1] == 168) || (bytes[0] == 169 && bytes[1] == 254))
            throw new SecurityException("Private IP blocked");
    }
    return uri;
}
```
