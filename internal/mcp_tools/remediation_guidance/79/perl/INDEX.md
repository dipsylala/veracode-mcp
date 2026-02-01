# CWE-79: Cross-Site Scripting (XSS) - Perl

## LLM Guidance

Cross-Site Scripting (CWE-79) occurs when untrusted data is included in web pages without proper encoding, allowing attackers to inject malicious scripts that execute in victim browsers. This leads to session hijacking, credential theft, or malware distribution. Perl applications must use context-appropriate encoding functions like `escapeHTML()` from CGI.pm or `encode_entities()` from HTML::Entities before outputting user data to HTML contexts.

## Remediation Strategy

- Always HTML-encode user input before rendering in HTML contexts using `encode_entities()` or `escapeHTML()`
- Use parameterized templates with auto-escaping (Template Toolkit with HTML filter)
- Validate and sanitize input on server-side; apply allowlists for expected formats
- Set Content-Security-Policy headers to restrict script execution sources
- Never insert untrusted data directly into JavaScript, CSS, or URL contexts without proper encoding

## Remediation Steps

- Identify all locations where user input is rendered in HTML output
- Replace direct variable interpolation with HTML encoding functions
- Use `HTML - -Entities - -encode_entities($user_input)` for HTML body/attribute contexts
- Enable auto-escaping in template engines (Template Toolkit's HTML filter)
- Implement Content-Security-Policy headers to block inline scripts
- Test with XSS payloads like `<script>alert(1)</script>` to verify fixes

## Minimal Safe Pattern

```perl
use CGI qw(:standard);
use HTML::Entities;

my $user_name = param('name') || '';
my $user_comment = param('comment') || '';

# Safe: HTML-encode before output
my $safe_name = encode_entities($user_name);
my $safe_comment = encode_entities($user_comment);

print header(),
      start_html('User Profile'),
      h1("Welcome, $safe_name"),
      p("Your comment: $safe_comment"),
      end_html();
```
