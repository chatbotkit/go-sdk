# Stateless Agent

Demonstrates manually driving `agent.CompleteWithTools` one iteration at a time, with full control over the loop.

## What it demonstrates

- Calling `agent.CompleteWithTools` directly instead of `agent.ExecuteWithTools`
- Inspecting `end.reason` after each iteration to decide whether to continue
- Appending `MessageAgentEvent` entries to the local message slice so the model sees tool results on the next call
- Patterns for observability, rate limiting, and early termination

## When to use this pattern

| Use case | Why |
| --- | --- |
| Observability | Log or audit every step before continuing |
| Custom logic | Transform responses or inject messages between iterations |
| Rate limiting | Add delays between API calls |
| Early termination | Stop on a custom condition without waiting for the model to call `exit` |

## Usage

```bash
export CHATBOTKIT_API_SECRET="your-api-key"
go run .
```

## Example session

```
Starting manually-driven agent loop...
User: What is the weather in San Francisco and what time is it in Los Angeles?
---

[Iteration 1]

🔧 Calling get_weather...
   ✓ get_weather returned

🔧 Calling get_time...
   ✓ get_time returned

The weather in San Francisco is sunny at 72°F. The current time in Los Angeles is 2:30 PM.

End reason: stop
→ Model completed naturally

---
Final conversation:
User: What is the weather in San Francisco and what time is it in Los Angeles?
Bot: The weather in San Francisco is sunny at 72°F...
```

## Key difference from `ExecuteWithTools`

`ExecuteWithTools` manages the iteration loop for you. This example replaces that loop with your own `for` statement, calling `CompleteWithTools` once per pass and checking `end.reason`:

```go
for iterationCount < maxIterations {
    events, errs := agent.CompleteWithTools(ctx, client, opts)

    // ... process events, collect endReason ...

    if endReason == "stop" {
        break // model finished naturally
    }
    // endReason == "iteration" → continue the loop
}
```
