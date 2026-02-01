# CWE-926: Android Component Export

## LLM Guidance

Android Component Export occurs when application components (activities, services, broadcast receivers, or content providers) are unintentionally exposed to other applications. This vulnerability arises when components lack explicit `android:exported` declarations, allowing unauthorized apps to interact with sensitive functionality. All exposed components must be intentionally configured and protected.

## Key Principles

- Never allow components to be accessible by other applications unless explicitly intended and protected
- All activities, services, receivers, and providers must declare `android:exported` explicitly in AndroidManifest.xml
- Any component exposed beyond the application boundary must enforce authorization through signature-level permissions or runtime caller validation
- Default to `android:exported="false"` for internal-only components

## Remediation Steps

- Review flaw details to identify which component lacks proper export declaration
- Audit all activities, services, broadcast receivers, and content providers in `AndroidManifest.xml`
- Identify components with intent filters that require explicit `android -exported` on Android 12+
- Determine which components should be internal-only versus accessible to other applications
- Set `android -exported="false"` for internal components and `android -exported="true"` only when necessary
- Add signature-level permissions or implement runtime caller validation for all exported components
