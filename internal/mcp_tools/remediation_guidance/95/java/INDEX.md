# CWE-95: Eval Injection - Java

## LLM Guidance

Java eval injection occurs when untrusted input flows into dynamic code execution mechanisms like scripting engines (Nashorn, Groovy), reflection APIs, or expression languages (SpEL, OGNL, MVEL). While Java lacks a native `eval()`, these components enable runtime code execution. The core fix is eliminating dynamic code execution entirely or using strict allow-lists with sandboxed execution contexts.

## Key Principles

- Eliminate dynamic evaluation: Replace scripting engines and expression languages with static business logic
- Use allow-lists over deny-lists: If dynamic features are required, restrict to predefined safe operations only
- Apply strict input validation: Validate against narrow patterns before any dynamic processing
- Isolate execution contexts: Use SecurityManager, separate classloaders, or restricted script engine bindings
- Prefer safe alternatives: Use configuration files, domain-specific languages, or rule engines with declarative syntax

## Remediation Steps

- Identify all uses of ScriptEngine, expression evaluators (SpEL, OGNL), and reflection with user input
- Remove dynamic evaluation and replace with static method calls or lookup maps
- If dynamic features are unavoidable, implement strict allow-list validation of all inputs
- Configure script engine bindings to expose only required, safe objects
- Enable Java SecurityManager with restrictive policies for scripted code
- Add automated scanning to detect new eval injection vectors in code reviews

## Safe Pattern

```java
// Replace dynamic evaluation with safe lookup map
public class SafeCalculator {
    private static final Map<String, BiFunction<Double, Double, Double>> OPS = Map.of(
        "add", (a, b) -> a + b,
        "subtract", (a, b) -> a - b,
        "multiply", (a, b) -> a * b
    );
    
    public double calculate(String operation, double a, double b) {
        BiFunction<Double, Double, Double> op = OPS.get(operation);
        if (op == null) {
            throw new IllegalArgumentException("Invalid operation");
        }
        return op.apply(a, b);
    }
}
```
