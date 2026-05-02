# Stateful Agent

Demonstrates the same manually-driven `agent.CompleteWithTools` loop as the stateless-agent example, but against a persisted ChatBotKit conversation where the server owns the message history.

## What it demonstrates

- Creating a conversation up front with `client.Conversation.Create`
- Passing `ConversationID` instead of a `Messages` slice
- Sending the initial user prompt on the first iteration only; omitting `Text` on subsequent iterations so the server continues from its own state
- Fetching the final persisted messages with `client.Conversation.Message.List`

## When to use this pattern

| Use case | Why |
| --- | --- |
| Persisted history | The server stores every message; no local slice to maintain |
| Remote orchestration | Resume a long-running conversation by ID from any process |
| Recovery | Restart without reconstructing a local message array |
| Inspection | Query the final conversation state via the API after the loop |

## Usage

```bash
export CHATBOTKIT_API_SECRET="your-api-key"
go run .
```

## Example session

```
Starting manually-driven stateful agent loop...
Conversation: conv_abc123
User: What is the weather in San Francisco and what time is it in Los Angeles?
---

[Iteration 1]
Calling get_weather...
Returned from get_weather
Calling get_time...
Returned from get_time
The weather in San Francisco is sunny at 72°F. The time in Los Angeles is 2:30 PM.

End reason: stop
Agent completed naturally

---
Final persisted conversation:
User: What is the weather in San Francisco and what time is it in Los Angeles?
Bot: The weather in San Francisco is sunny at 72°F...
```

## Key difference from the stateless example

The stateless example appends `MessageAgentEvent` entries to a local `[]agent.Message` slice. Here, that slice is replaced by a server-side conversation — the only state the loop tracks locally is the conversation ID and whether to send `Text` on the next call:

```go
nextText := &userPrompt // send prompt on first iteration

for iterationCount < maxIterations {
    events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
        ConversationID: conversation.ID,
        Text:           nextText,
        Tools:          tools,
    })
    // ...
    nextText = nil // omit on subsequent iterations
}
```
