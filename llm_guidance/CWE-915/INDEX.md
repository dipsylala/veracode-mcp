# CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes

## LLM Guidance

This vulnerability occurs when user input dynamically modifies object attributes or properties, allowing attackers to alter application behavior or access unauthorized data. Mass assignment (CWE-915) controls **which** properties can be modified, while CWE-1174 validates the **values** of allowed properties. Never allow mass assignment of object attributes; allowlist permitted fields and enforce invariants server-side.

## Key Principles

- Use explicit allowlists to restrict which object attributes can be modified by user input
- Reject dynamic attribute assignment patterns that accept arbitrary property names from untrusted sources
- Enforce security-critical attributes (isAdmin, role, price, balance) server-side, never from client input
- Validate both property names and values before modification
- Prefer explicit property binding over reflection-based or dynamic assignment

## Remediation Steps

- Locate the vulnerability - Review flaw details for file, line number, and code pattern. Trace data flow from source (HTTP parameters, JSON, form data) to sink (`setattr()`, `__dict__`, `obj[user_input] =`, `Object.defineProperty()`, mass assignment)

- Identify risk - Determine if attackers can modify security-critical attributes that control authorization, pricing, or application state

- Implement allowlist - Define explicitly permitted attributes that users can modify, rejecting all others by default

- Replace dynamic assignment - Use explicit property setters or data transfer objects (DTOs) with known fields instead of dynamic reflection or dictionary-based assignment

- Enforce server-side validation - Validate property values and maintain security invariants regardless of client input

- Test protection - Verify that unauthorized attribute modification attempts are blocked and logged
