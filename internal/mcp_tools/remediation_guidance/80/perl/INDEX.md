# CWE-80: Cross-Site Scripting (XSS) - Perl

## LLM Guidance

Cross-Site Scripting (CWE-80) occurs when untrusted data is included in web pages without proper encoding. Attackers inject malicious scripts that execute in victim browsers, stealing sessions, credentials, or performing actions on behalf of users. Perl applications must encode all user-controlled output using context-appropriate functions.

**Primary Defence:** Use CGI.pm's `escapeHTML()` for HTML contexts, URL encoding for URLs, and JavaScript encoding for script contexts.

## Key Principles

- Encode all user input before outputting to HTML using `CGI::escapeHTML()` or HTML::Entities
- Use Content Security Policy headers to restrict script execution sources
- Apply context-specific encoding (HTML, URL, JavaScript) based on output location
- Validate input against allowlists where possible, then encode before output
- Never insert untrusted data directly into JavaScript blocks or event handlers

## Remediation Steps

- Replace all direct output of user data with `escapeHTML()` wrapper calls
- Implement CSP headers with `script-src 'self'` directive to block inline scripts
- Review all CGI parameter usage and apply appropriate encoding functions
- Use templating systems with auto-escaping (Template Toolkit with HTML filter)
- Test with XSS payloads like `<script>alert(1)</script>` to verify protection
- Audit code for direct `print` statements containing CGI parameters

## Safe Pattern

```perl
#!/usr/bin/perl
use strict;
use warnings;
use CGI qw(:standard);

my $cgi = CGI->new;
my $user_input = $cgi->param('name') || '';

# Encode user input before output
my $safe_output = escapeHTML($user_input);

print $cgi->header('text/html');
print "<html><body>";
print "<h1>Welcome, $safe_output</h1>";
print "</body></html>";
```
