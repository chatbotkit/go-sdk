# ChatBotKit Go SDK Examples

This directory contains example applications demonstrating how to use the ChatBotKit Go SDK.

## Prerequisites

- Go 1.21 or later
- A ChatBotKit API key (get one at [chatbotkit.com](https://chatbotkit.com))

## Examples

### Chatbot

A simple interactive chatbot that demonstrates:

- Creating a ChatBotKit client
- Running an interactive conversation loop
- Streaming responses in real-time
- Maintaining conversation history

**Run the example:**

```bash
cd examples/chatbot

# Option 1: Use a .env file
echo 'CHATBOTKIT_API_SECRET=your-api-key' > .env
go run main.go

# Option 2: Use environment variable
export CHATBOTKIT_API_SECRET="your-api-key"
go run main.go
```

**Example session:**

```
ChatBot ready! Type your message and press Enter. Type 'exit' to quit.

user: Hello! What can you do?
bot: Hello! I'm an AI assistant powered by ChatBotKit. I can help you with a variety of tasks like answering questions, explaining concepts, helping with writing, brainstorming ideas, and much more. What would you like to know or discuss?

user: Tell me a short joke
bot: Why don't scientists trust atoms? Because they make up everything! 😄

user: exit
Goodbye!
```

### Agent with Tools

An advanced example demonstrating agent execution with custom tool registration:

- Defining tools with JSON Schema parameters
- Registering tool handlers that execute when called
- Streaming events including tool call start/end notifications
- Building autonomous agents that can use tools

**Run the example:**

```bash
cd examples/agent-with-tools

# Option 1: Use a .env file
echo 'CHATBOTKIT_API_SECRET=your-api-key' > .env
go run main.go

# Option 2: Use environment variable
export CHATBOTKIT_API_SECRET="your-api-key"
go run main.go
```

**Example session:**

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

**Defining Custom Tools:**

```go
tools := agent.Tools{
    "my_tool": {
        Description: "Description of what the tool does",
        Parameters: &agent.Parameters{
            Properties: map[string]agent.Property{
                "param1": {
                    Type:        "string",
                    Description: "First parameter",
                },
                "param2": {
                    Type:        "number",
                    Description: "Second parameter",
                },
            },
            Required: []string{"param1"},
        },
        Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
            param1 := args["param1"].(string)
            // Process and return result
            return map[string]interface{}{"result": param1}, nil
        },
    },
}
```

## Running Examples from the SDK Root

You can also run examples from the SDK root directory:

```bash
export CHATBOTKIT_API_SECRET="your-api-key"
go run ./examples/chatbot
go run ./examples/agent-with-tools
```

## Environment Variables

| Variable                | Description             | Required |
| ----------------------- | ----------------------- | -------- |
| `CHATBOTKIT_API_SECRET` | Your ChatBotKit API key | Yes      |

## Creating Your Own Chatbot

The chatbot example can be customized by modifying the `CompleteOptions`:

```go
events, errs := agent.CompleteStream(ctx, client, agent.CompleteOptions{
    Model:     "gpt-4o",           // AI model to use
    Messages:  messages,            // Conversation history
    Backstory: "You are a helpful coding assistant.", // Optional persona
})
```

## Creating Agents with Tools

For more advanced use cases, use `CompleteWithTools` to register custom functions:

```go
events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
    Model:     "gpt-4o",
    Messages:  messages,
    Backstory: "You are a helpful assistant with access to tools.",
    Tools:     tools,  // Your custom tools
})
```

For autonomous task execution with built-in planning and progress tracking:

```go
events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
    Model:         "gpt-4o",
    Messages:      []agent.Message{{Type: "user", Text: "Complete this task..."}},
    Backstory:     "You are an autonomous task executor.",
    Tools:         tools,
    MaxIterations: 20,
})
```

See the [SDK documentation](../README.md) for more options and features.
