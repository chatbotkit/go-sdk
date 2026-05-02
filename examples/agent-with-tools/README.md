# Agent with Custom Tools

An interactive chatbot that registers custom tools the AI can call during a conversation.

## What it demonstrates

- Defining tools with JSON Schema parameters
- Registering tool handlers that execute when the model calls them
- Streaming `ToolCallStart` / `ToolCallEnd` events alongside response tokens
- Running an interactive multi-turn conversation with tool access

## Usage

```bash
# Option 1: .env file
echo 'CHATBOTKIT_API_SECRET=your-api-key' > .env
go run .

# Option 2: environment variable
export CHATBOTKIT_API_SECRET="your-api-key"
go run .
```

Type `exit` to quit.

## Example session

```
Agent with Tools ready! Type your message and press Enter. Type 'exit' to quit.
Available tools: get_current_time, calculate, search_knowledge

user: What time is it?
  [Calling get_current_time with map[]...]
  [get_current_time returned: map[datetime:2024-01-15T14:30:00Z timezone:UTC]]
bot: The current time is 2:30 PM UTC.

user: Calculate 15 * 7
  [Calling calculate with map[a:15 b:7 operation:multiply]...]
  [calculate returned: map[a:15 b:7 operation:multiply result:105]]
bot: 15 × 7 = 105

user: exit
Goodbye!
```

## Defining a tool

```go
tools := agent.Tools{
    "my_tool": {
        Description: "What the tool does",
        Parameters: agent.FunctionParameters{
            "properties": map[string]any{
                "input": map[string]any{
                    "type":        "string",
                    "description": "The input value",
                },
            },
            "required": []string{"input"},
        },
        Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
            input := args["input"].(string)
            return map[string]interface{}{"result": input}, nil
        },
    },
}
```
