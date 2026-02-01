# CWE-918: Server-Side Request Forgery (SSRF) - PHP

## LLM Guidance

Server-Side Request Forgery (SSRF) allows attackers to make the server perform HTTP requests to arbitrary destinations, potentially accessing internal services, cloud metadata endpoints, or bypassing firewalls. Always validate URLs against an allowlist of permitted domains, block private IP ranges (127.0.0.0/8, 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, 169.254.0.0/16), and restrict protocols to HTTPS only.

## Remediation Strategy

- Use strict allowlist validation for permitted domains before making requests
- Block private, reserved, and loopback IP ranges using PHP's filter functions
- Disable URL redirects or validate redirect destinations
- Restrict protocols to HTTPS only, never allow file://, gopher://, or other schemes
- Implement DNS rebinding protection by re-resolving hostnames after connection

## Remediation Steps

- Extract and validate the hostname from user-supplied URLs
- Use `filter_var()` with `FILTER_VALIDATE_IP` and flags `FILTER_FLAG_NO_PRIV_RANGE | FILTER_FLAG_NO_RES_RANGE` to block dangerous IPs
- Check hostname against an allowlist of permitted domains
- Ensure only HTTPS protocol is used via `parse_url()`
- Disable `CURLOPT_FOLLOWLOCATION` or validate all redirect targets
- Set timeouts and use `CURLOPT_PROTOCOLS` to restrict allowed protocols

## Minimal Safe Pattern

```php
function safeFetchUrl($url, array $allowedHosts) {
    $parsed = parse_url($url);
    if (!$parsed || $parsed['scheme'] !== 'https' || !in_array($parsed['host'], $allowedHosts, true)) {
        throw new Exception('Invalid URL');
    }
    
    $ip = gethostbyname($parsed['host']);
    if (!filter_var($ip, FILTER_VALIDATE_IP, FILTER_FLAG_NO_PRIV_RANGE | FILTER_FLAG_NO_RES_RANGE)) {
        throw new Exception('Private IP not allowed');
    }
    
    $ch = curl_init($url);
    curl_setopt($ch, CURLOPT_FOLLOWLOCATION, false);
    curl_setopt($ch, CURLOPT_PROTOCOLS, CURLPROTO_HTTPS);
    return curl_exec($ch);
}
```
