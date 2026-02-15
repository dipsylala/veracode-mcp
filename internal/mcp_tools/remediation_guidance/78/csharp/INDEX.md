# CWE-78: OS Command Injection - C#

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands on the host. In C#, eliminate Process.Start() and ProcessStartInfo calls entirely by using .NET Framework classes (System.IO, System.Net, System.IO.Compression) for file operations, HTTP requests, and archive handling.

## Key Principles

- Replace all Process.Start() and ProcessStartInfo calls with .NET Framework class alternatives
- Use System.IO.File and System.IO.Directory for file operations instead of system commands
- Use System.Net.Http.HttpClient for HTTP requests instead of curl/wget
- Use System.Net.NetworkInformation.Ping for network checks instead of ping command
- Use System.IO.Compression for archive operations instead of zip commands
- Never concatenate user input into command strings
- Only use ProcessStartInfo as a last resort with ArgumentList and UseShellExecute = false

## Remediation Steps

- Locate command execution - Identify all Process.Start() and ProcessStartInfo instances
- Determine the operation's purpose - Understand what the command is trying to accomplish
- Find the .NET class alternative - Use System.IO for file ops, HttpClient for HTTP, Ping for network
- Replace process execution - Delete Process.Start() code and use the appropriate .NET class
- For unavoidable commands - Use ProcessStartInfo with ArgumentList and UseShellExecute = false, validate all inputs
- Test thoroughly - Verify the .NET class replacement provides the same functionality

## Safe Pattern

```csharp
// UNSAFE: Executing ping command
var startInfo = new ProcessStartInfo 
{
    FileName = "cmd.exe",
    Arguments = $"/c ping {host}"
};
Process.Start(startInfo);

// UNSAFE: Even with ArgumentList
var psi = new ProcessStartInfo("ping") { UseShellExecute = false };
psi.ArgumentList.Add(host);
Process.Start(psi);

// SAFE: Use Ping class for reachability check
using System.Net.NetworkInformation;

public async Task<bool> IsHostReachable(string host, int timeout = 5000)
{
    try 
    {
        using var ping = new Ping();
        var reply = await ping.SendPingAsync(host, timeout);
        return reply.Status == IPStatus.Success;
    } 
    catch (PingException ex) 
    {
        Console.WriteLine($"Ping failed: {ex.Message}");
        return false;
    }
}

// SAFE: File copy with File class
using System.IO;
File.Copy(source, dest, overwrite: true);

// SAFE: HTTP request with HttpClient
using System.Net.Http;
using var httpClient = new HttpClient();
var response = await httpClient.GetAsync(url);
var content = await response.Content.ReadAsStringAsync();
```
