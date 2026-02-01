# CWE-942: Overly Permissive CORS

## LLM Guidance

Overly permissive CORS occurs when web applications allow requests from any origin or untrusted origins, exposing sensitive resources to unauthorized websites. This enables data theft, account compromise, and cross-origin attacks. The core fix is restricting allowed origins to a specific allowlist of trusted domains and never using wildcards with credentials.

## Key Principles

- **Restrict allowed origins**: Only allow specific trusted origins; never use `*` in production or reflect Origin headers dynamically
- **Limit methods and headers**: Permit only required HTTP methods and headers; minimize the attack surface
- **Disable credentials by default**: Set `Access-Control-Allow-Credentials: false` unless absolutely necessary; never combine wildcards with credentials
- **Monitor and audit**: Log CORS requests, track unexpected origins, and regularly review policy effectiveness

## Actionable Steps

- **Check response headers** for `Access-Control-Allow-Origin: *` and `Access-Control-Allow-Credentials: true` combinations
- **Review CORS middleware** for origin reflection patterns that echo back request origins without validation
- **Replace wildcards** with explicit allowlists of trusted domains in configuration
- **Validate origin allowlists** against business requirements; remove unused or overly broad entries
- **Audit preflight handling** to ensure OPTIONS responses don't permit excessive methods or headers
- **Test CORS policies** by sending requests from untrusted origins and verifying they're rejected
